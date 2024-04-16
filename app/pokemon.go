package app

import (
	"math/rand"
	"encoding/json"
	"net/http"
	"fmt"

	"strconv"
)
//何個取得してシャッフルするか
const limit int = 100
//javascriptに送る数
const sent int = 50


// 機能１ポケモンデータを取ってきて必要なフィールだけを持った構造体に格納した上で配列にプッシュ。配列を返す

type Pokemon struct{
	Name string 
	Height float64
	Weight float64
	Ability string
	Imgurl string

}

type Abilitiesinfo struct{
	Ability Item `json:"ability"`
	//もし'is_hidden'や'slot'がほしければここに追加
}


//正面の画像
type Sprit struct{
	Frontdefault string `json:"front_default"`
}




type Item struct{
	Name string `json:"name"`
	URL string `json:"url"`

}








func sum() []Pokemon{
	var namelist []Item
	var pokemonlist []Pokemon
	// query := r.URL.Query()
	// offset := query.Get("offset")
	// if offset == ""{
	// 	offset = "0"
	// }
	offset := "0"
	//offsetをランダムにしたらもっと幅広く取得できるかも。その代わりcountを取得するか定数(1302)を用いなければいけない。
	//offset = rand.Intn(count-limit)
	url := "https://pokeapi.co/api/v2/pokemon/?offset=" + offset + "&limit=" + strconv.Itoa(limit)

		

	resp, err := http.Get(url)

	if err != nil{
		//javascriptにエラーを送る
		fmt.Println("error")

	}
	defer resp.Body.Close()

	var response struct {
		Results []Item `json:"results"`
		//すべてを取得するときには必要だけどなくてもいいかも
		Next string `json:"next"`

	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to decode JSON:", err)

	}


	// レスポンスの配列に含まれる各要素を表示
	for _, item := range response.Results {
		pokemon := statusHandler(item.URL, item.Name)


		pokemonlist = append(pokemonlist, pokemon)

	}
	// url = response.Next


	

	pokemonlist = shuffle(pokemonlist)

	//returnする
    // Content-Typeヘッダーをapplication/jsonに設定
    return pokemonlist[:sent]

}

//urlを取得してその中から必要なデータを持ってくる。


func statusHandler(url, name string) Pokemon {
	resp, err := http.Get(url)
	if err != nil{
		fmt.Println("error")

	}
	defer resp.Body.Close()
	//urlから名前以外の情報を取得

	var response struct{
		Abilities []Abilitiesinfo `json:"abilities"`
		Height float64 `json:"height"`
		Weight float64 `json:"weight"`

		//画像の取得
		Sprites Sprit `json:"sprites"`

	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to decode JSON:", err)
		fmt.Println("response")

	}

	var skill string
	//技が一つもないことも考えないといけないかも(if文で)
	if len(response.Abilities) == 0{
		fmt.Println("null")
		skill = "null"
	}else{
		ability := response.Abilities[0].Ability
		skill = ability.Name
	}
	
	var pokemon Pokemon
	pokemon.Name = name
	pokemon.Ability = skill
	pokemon.Height = response.Height
	pokemon.Weight = response.Weight
	pokemon.Imgurl = response.Sprites.Frontdefault

	return pokemon
}

func shuffle(arr []Pokemon) []Pokemon{
	for i := 0; i < len(arr); i++{
		r := rand.Intn(len(arr)-1-i) + i


		temp := arr[r]
		arr[r] = arr[i]
		arr[i] = temp
	}
	return arr
}






// 機能２ ブラウザからクライエントが入力した整数を使ってポケモンデータをフィルタリング。身長を降順にソートした上で配列を返す
// →機能２は機能１をそのまま使いたい→機能１で作ったポケモン構造体を持つ配列を使ってフィルタリング・ソーティングをするといいと思う
