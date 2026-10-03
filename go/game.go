package main

import (
	"encoding/json"
	"errors"
	"math/rand"
	"os"
	"time"
)

// Rarity はブルーベリーのレア度です。
type Rarity string

const (
	Common Rarity = "common"
	Rare   Rarity = "rare"
	Epic   Rarity = "epic"
)

// Label は画面表示用の日本語ラベルを返します。
func (r Rarity) Label() string {
	switch r {
	case Epic:
		return "超レア"
	case Rare:
		return "レア"
	default:
		return "通常"
	}
}

// Berry は品種1つ分のデータです。
type Berry struct {
	Name   string
	Family string
	Rarity Rarity
	Reward int
}

var berryTypes = []Berry{
	{"Bluecrop", "北部ハイブッシュ", Common, 3},
	{"Duke", "北部ハイブッシュ", Common, 4},
	{"Spartan", "北部ハイブッシュ", Common, 4},
	{"Patriot", "北部ハイブッシュ", Common, 4},
	{"Elliott", "北部ハイブッシュ", Common, 5},
	{"Legacy", "北部ハイブッシュ", Common, 5},
	{"Chandler", "北部ハイブッシュ", Rare, 7},
	{"Liberty", "北部ハイブッシュ", Common, 5},
	{"Aurora", "北部ハイブッシュ", Rare, 7},
	{"Jersey", "北部ハイブッシュ", Common, 4},
	{"Berkeley", "北部ハイブッシュ", Common, 4},
	{"Toro", "北部ハイブッシュ", Common, 5},
	{"Northland", "ハーフハイブッシュ", Common, 4},
	{"Northblue", "ハーフハイブッシュ", Rare, 6},
	{"Northcountry", "ハーフハイブッシュ", Common, 4},
	{"Sunshine Blue", "サザンハイブッシュ", Rare, 7},
	{"Misty", "サザンハイブッシュ", Common, 5},
	{"O'Neal", "サザンハイブッシュ", Rare, 7},
	{"Sharpblue", "サザンハイブッシュ", Common, 5},
	{"Biloxi", "サザンハイブッシュ", Rare, 7},
	{"Powderblue", "ラビットアイ", Common, 5},
	{"Tifblue", "ラビットアイ", Common, 5},
	{"Brightwell", "ラビットアイ", Common, 5},
	{"Climax", "ラビットアイ", Common, 5},
	{"Premier", "ラビットアイ", Common, 5},
	{"Gardenblue", "ラビットアイ", Rare, 7},
	{"Columbus", "ラビットアイ", Common, 6},
	{"Pink Lemonade", "ラビットアイ", Epic, 12},
	{"Ochlockonee", "ラビットアイ", Rare, 8},
	{"Baldwin", "ラビットアイ", Rare, 8},
}

var rarityWeights = map[Rarity]int{
	Common: 72,
	Rare:   24,
	Epic:   4,
}

const farmerCost = 20

// State はセーブされるゲームの進行状況です。
type State struct {
	Money        int       `json:"money"`
	Discovered   []int     `json:"discovered"`
	CurrentBerry *int      `json:"currentBerryIndex"`
	Farmers      int       `json:"farmers"`
	LastTick     time.Time `json:"lastTick"`
}

// Game はゲーム本体です。乱数は差し替えられるようにしてテストしやすくしています。
type Game struct {
	State State
	rng   *rand.Rand
}

func NewGame(rng *rand.Rand) *Game {
	return &Game{State: State{LastTick: time.Now()}, rng: rng}
}

// pickBerry はレア度の重みで抽選し、その中から品種をランダムに選びます。
func (g *Game) pickBerry() int {
	total := 0
	for _, w := range rarityWeights {
		total += w
	}
	roll := g.rng.Intn(total)

	var rarity Rarity
	for _, r := range []Rarity{Common, Rare, Epic} {
		if roll < rarityWeights[r] {
			rarity = r
			break
		}
		roll -= rarityWeights[r]
	}

	var candidates []int
	for i, b := range berryTypes {
		if b.Rarity == rarity {
			candidates = append(candidates, i)
		}
	}
	return candidates[g.rng.Intn(len(candidates))]
}

var (
	ErrAlreadyGrown     = errors.New("もう実っています。先に収穫しましょう")
	ErrNothingToHarvest = errors.New("まだ実っていません。先に育てましょう")
	ErrNotEnoughMoney   = errors.New("コインが足りません")
)

func (g *Game) Grow() (Berry, error) {
	if g.State.CurrentBerry != nil {
		return Berry{}, ErrAlreadyGrown
	}
	idx := g.pickBerry()
	g.State.CurrentBerry = &idx
	return berryTypes[idx], nil
}

// Harvest は収穫してコインを得ます。初発見なら isNew が true になります。
func (g *Game) Harvest() (berry Berry, isNew bool, err error) {
	if g.State.CurrentBerry == nil {
		return Berry{}, false, ErrNothingToHarvest
	}
	idx := *g.State.CurrentBerry
	berry = berryTypes[idx]
	g.State.Money += berry.Reward
	g.State.CurrentBerry = nil

	if !g.IsDiscovered(idx) {
		g.State.Discovered = append(g.State.Discovered, idx)
		isNew = true
	}
	return berry, isNew, nil
}

func (g *Game) HireFarmer() error {
	if g.State.Money < farmerCost {
		return ErrNotEnoughMoney
	}
	g.State.Money -= farmerCost
	g.State.Farmers++
	return nil
}

// Tick は前回から経過した秒数ぶん、農夫の自動収入を加算します。
func (g *Game) Tick(now time.Time) int {
	secs := int(now.Sub(g.State.LastTick) / time.Second)
	if secs <= 0 {
		return 0
	}
	g.State.LastTick = g.State.LastTick.Add(time.Duration(secs) * time.Second)
	earned := secs * g.State.Farmers
	g.State.Money += earned
	return earned
}

func (g *Game) IsDiscovered(idx int) bool {
	for _, d := range g.State.Discovered {
		if d == idx {
			return true
		}
	}
	return false
}

func (g *Game) Save(path string) error {
	data, err := json.MarshalIndent(g.State, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// Load はセーブファイルを読み込みます。ファイルが無ければ何もしません。
func (g *Game) Load(path string) error {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &g.State)
}
