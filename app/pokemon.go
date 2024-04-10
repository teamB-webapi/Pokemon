package app

import (
	"math/rand"
	"encoding/json"
	"net/http"
	"fmt"

	"strconv"
)
const limit int = 50

// テスト用のために作成
func sum() int {
	return rand.Intn(100)
}

// 機能１ポケモンデータを取ってきて必要なフィールだけを持った構造体に格納した上で配列にプッシュ。配列を返す

type Pokemon struct{
	Name string
	Height float64
	Weight float64
	Skill string
	Imgurl string
}

type Item struct{
	Name int `json:"name"`
	URL string `json:"url"`
	Next string `json:"next"`
}



var pokemonlist []Pokemon


func allgetHandler(w http.ResponseWriter, r*http.Request){
	
	query := r.URL.Query()
	offset := query.Get("offset")
	if offset == ""{
		offset = "0"
	}

	url := "https://pokeapi.co/api/v2/pokemon/?offset=" + offset + "&limit=" + strconv.Itoa(limit)

	resp, err := http.Get(url)
	if err != nil{
		//javascriptにエラーを送る
		fmt.Println("error")
		return
	}
	defer resp.Body.Close()

	var response struct {
		Results []Item `json:"results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to decode JSON:", err)
		return
	}

	// レスポンスの配列に含まれる各要素を表示
	for _, item := range response.Results {
		fmt.Println("Name:", item.Name)
		fmt.Println("URL:", item.URL)
	}
	fmt.Println("URL:")

	//必要なデータを持ってくる関数を呼び出す


	//持ってきたデータをpokemonlistにいれる


	//returnする


	
}

//urlを取得してその中から必要なデータを持ってくる。

// func statusHandler(url, name string){
// 	resp, err := http.Get(url)
// 	if err != nill{
// 		//falseを返すもしくは空を返す
// 		return
// 	}
// 	defer resp.Body.Closse()
// 	//urlから名前以外の情報を取得
// 	pokemon := Pokemon(name, height, weight, skill, imgurl)
// 	return pokemon
// }

// func main() {
//     fmt.Println("Starting the server!")
    
//     // ルートとハンドラ関数を定義
//     http.HandleFunc("/api/v2/pokemon/", allgetHandler)

//     // 8000番ポートでサーバを開始
//     http.ListenAndServe(":8000", nil)
// }




// 機能２ ブラウザからクライエントが入力した整数を使ってポケモンデータをフィルタリング。身長を降順にソートした上で配列を返す
// →機能２は機能１をそのまま使いたい→機能１で作ったポケモン構造体を持つ配列を使ってフィルタリング・ソーティングをするといいと思う
