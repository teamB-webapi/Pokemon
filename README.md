# Pokemon API
## 概要
ランダムな10体のポケモンの名前、技、画像を表示するアプリです。

## DEMO
[Uploading 無題の動画 ‐ Clipchampで作成.mp4…](https://github.com/teamB-webapi/Pokemon/assets/116293482/2dd630cb-d3d2-46c9-931c-db4b689a2b61)


## 使用技術
- フロントエンド：HTML、CSS、JavaScript
- バックエンド：Go
- その他：Ubuntu

## 使用方法
### クローン
このプロジェクトをあなたのPCで実行するために、クローンします。

下記手順でクローンしてください。

1. リポジトリをクローンする
```
git clone https://github.com/teamB-webapi/Pokemon
```

1. クローンしたリポジトリへ移動する
```
cd POKEMON
```
### ローカルでサーバーを起動する
```
go run main.go
```


## エンドポイント
- ポケモンの名前の取得`/api/v2/pokemon/?offset=0&limit=30`
- ポケモンの身長、体重、能力、画像の取得`/api/v2/pokemon/{id}`

## 改善点
### ポケモンの表示数を増やす
- 決まった30体の中からしか取得できないので取得数を増やしていきたいと考えている。
### フィルタリング機能の追加
- 色以外のものでもフィルタリングできるようしたいと考えている
### 検索機能の追加
- 自分の好きなポケモンについての情報を取得できるようにしたいと考えている。
### 本番環境でも使えるようにする
- 今はローカル環境で環境構築しないと動かせないが、本番環境でも使えるようにしたいと考えている。

## チームメンバー
- [REI](https://github.com/ReiNagahashi)
- [kolomame](https://github.com/kolomame)
- [nyantarou](https://github.com/nyantarou2001002)
