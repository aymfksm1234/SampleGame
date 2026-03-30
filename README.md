# SampleGame

ブルーベリーを育てて収穫し、図鑑を埋めていく放置系ミニゲームです。  
単一ファイルの `index.html` だけで動くので、そのまま GitHub Pages に公開できます。

## 遊び方

1. `育てる` を押してブルーベリーを実らせる
2. `収穫する` でコインを獲得する
3. 新しい品種を見つけて図鑑を埋める
4. `農夫を雇う` で毎秒コインが入るようにする

## できること

- 実在するブルーベリー30品種を収集
- 未発見の品種は `???` で表示
- 農夫による自動収入
- `localStorage` を使った進行保存
- スマホでも遊びやすいレスポンシブ表示

## ファイル構成

- [index.html](/Users/ayumif/Documents/GameSample/index.html): ゲーム本体

## GitHub Pages

GitHub Pages を有効化すると、以下のURLで公開される想定です。

[https://aymfksm1234.github.io/SampleGame/](https://aymfksm1234.github.io/SampleGame/)

公開されない場合は、GitHub の `Settings > Pages` で以下を確認してください。

- Source: `Deploy from a branch`
- Branch: `main`
- Folder: `/ (root)`

## 今後の拡張案

- レア品種の専用演出
- 成長アニメーションの強化
- アップグレードショップ追加
- 効率の違う農夫や施設追加
