// ブルーベリー収集ゲームのターミナル版です。
//
//	go run .
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const saveFile = "blueberry_save.json"

func main() {
	g := NewGame(rand.New(rand.NewSource(time.Now().UnixNano())))
	if err := g.Load(saveFile); err != nil {
		fmt.Println("セーブデータを読み込めませんでした:", err)
	}

	fmt.Println("🫐 ブルーベリーコレクター (Go版)")
	printHelp()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if earned := g.Tick(time.Now()); earned > 0 {
			fmt.Printf("👨‍🌾 農夫たちが %d コイン稼ぎました\n", earned)
		}
		fmt.Printf("\n[コイン %d | 図鑑 %d/%d | 農夫 %d] > ",
			g.State.Money, len(g.State.Discovered), len(berryTypes), g.State.Farmers)

		if !scanner.Scan() {
			break
		}
		switch strings.TrimSpace(scanner.Text()) {
		case "g", "grow":
			if b, err := g.Grow(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("🌱 実りました!  %s\n", describe(g, b))
			}
		case "h", "harvest":
			b, isNew, err := g.Harvest()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Printf("🧺 %s を収穫! +%d コイン\n", b.Name, b.Reward)
			if isNew {
				fmt.Printf("✨ 新品種発見! [%s] %s\n", b.Rarity.Label(), b.Family)
			}
		case "f", "farmer":
			if err := g.HireFarmer(); err != nil {
				fmt.Printf("%s (%d コイン必要)\n", err, farmerCost)
			} else {
				fmt.Printf("農夫を雇いました。毎秒 %d コイン入ります\n", g.State.Farmers)
			}
		case "d", "dex":
			printDex(g)
		case "s", "save":
			if err := g.Save(saveFile); err != nil {
				fmt.Println("保存に失敗しました:", err)
			} else {
				fmt.Println("💾 保存しました")
			}
		case "?", "help":
			printHelp()
		case "q", "quit":
			if err := g.Save(saveFile); err != nil {
				fmt.Println("保存に失敗しました:", err)
			}
			fmt.Println("またね!")
			return
		case "":
		default:
			fmt.Println("知らないコマンドです。? でヘルプ")
		}
	}
}

// describe は未発見の品種を ??? で隠して表示します。
func describe(g *Game, b Berry) string {
	if g.State.CurrentBerry != nil && !g.IsDiscovered(*g.State.CurrentBerry) {
		return "??? (収穫すると判明)"
	}
	return fmt.Sprintf("%s [%s]", b.Name, b.Rarity.Label())
}

func printDex(g *Game) {
	fmt.Println("📖 図鑑")
	for i, b := range berryTypes {
		if g.IsDiscovered(i) {
			fmt.Printf("  %2d. %-14s %-6s %s\n", i+1, b.Name, b.Rarity.Label(), b.Family)
		} else {
			fmt.Printf("  %2d. ???\n", i+1)
		}
	}
}

func printHelp() {
	fmt.Println(`コマンド:
  g / grow     育てる
  h / harvest  収穫する
  f / farmer   農夫を雇う (20コイン)
  d / dex      図鑑を見る
  s / save     保存
  q / quit     保存して終了`)
}
