package main

import (
	"errors"
	"math/rand"
	"path/filepath"
	"testing"
	"time"
)

func newTestGame() *Game {
	return NewGame(rand.New(rand.NewSource(1)))
}

func TestGrowAndHarvest(t *testing.T) {
	g := newTestGame()

	if _, _, err := g.Harvest(); !errors.Is(err, ErrNothingToHarvest) {
		t.Fatalf("育てる前の収穫は失敗するはず: %v", err)
	}

	b, err := g.Grow()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := g.Grow(); !errors.Is(err, ErrAlreadyGrown) {
		t.Fatalf("二重に育てられないはず: %v", err)
	}

	got, isNew, err := g.Harvest()
	if err != nil {
		t.Fatal(err)
	}
	if got != b || !isNew || g.State.Money != b.Reward {
		t.Fatalf("収穫結果がおかしい: %+v isNew=%v money=%d", got, isNew, g.State.Money)
	}
}

func TestHireFarmerAndTick(t *testing.T) {
	g := newTestGame()
	if err := g.HireFarmer(); !errors.Is(err, ErrNotEnoughMoney) {
		t.Fatalf("お金が無ければ雇えないはず: %v", err)
	}

	g.State.Money = 25
	if err := g.HireFarmer(); err != nil {
		t.Fatal(err)
	}
	if g.State.Money != 5 || g.State.Farmers != 1 {
		t.Fatalf("money=%d farmers=%d", g.State.Money, g.State.Farmers)
	}

	start := g.State.LastTick
	if earned := g.Tick(start.Add(3500 * time.Millisecond)); earned != 3 {
		t.Fatalf("3秒で3コインのはず: %d", earned)
	}
	// 端数の0.5秒は次回に持ち越される
	if earned := g.Tick(start.Add(4 * time.Second)); earned != 1 {
		t.Fatalf("持ち越し分で1コインのはず: %d", earned)
	}
}

func TestPickBerryDistribution(t *testing.T) {
	g := newTestGame()
	counts := map[Rarity]int{}
	const n = 100000
	for i := 0; i < n; i++ {
		counts[berryTypes[g.pickBerry()].Rarity]++
	}
	for r, w := range rarityWeights {
		got := float64(counts[r]) / n * 100
		if diff := got - float64(w); diff > 1 || diff < -1 {
			t.Errorf("%s: %.1f%% (期待 %d%%)", r, got, w)
		}
	}
}

func TestSaveLoad(t *testing.T) {
	path := filepath.Join(t.TempDir(), "save.json")
	g := newTestGame()
	g.State.Money = 42
	g.State.Farmers = 2
	g.State.Discovered = []int{0, 27}
	if err := g.Save(path); err != nil {
		t.Fatal(err)
	}

	loaded := newTestGame()
	if err := loaded.Load(path); err != nil {
		t.Fatal(err)
	}
	if loaded.State.Money != 42 || loaded.State.Farmers != 2 || !loaded.IsDiscovered(27) {
		t.Fatalf("読み込み結果がおかしい: %+v", loaded.State)
	}

	if err := newTestGame().Load(filepath.Join(t.TempDir(), "none.json")); err != nil {
		t.Fatalf("ファイルが無くてもエラーにしない: %v", err)
	}
}
