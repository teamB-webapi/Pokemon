package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

// 頻繁にレンダリングされるhtmlファイルはここに登録しておくことで、何度もパースする必要が無くなる
var templates = template.Must(template.ParseFiles("app/home.html"))

// home.htmlをレンダリングする
func viewHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}



type JSONError struct{
	Error string `json:"error"`
	Code int `json:"code"`
}

// レスポンスが不適切な時とかに、JSON型のエラーを返したい
func APIError(w http.ResponseWriter, errMessage string, code int){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	jsonError, err := json.Marshal(JSONError{Error:errMessage, Code: code})
	if err != nil{
		log.Fatal(err)
	}

	w.Write(jsonError)
}


func apiPokemonHandler(w http.ResponseWriter, _ *http.Request) {
	// TODO: 変数pokemonDataには機能１である関数を実行してポケモンの構造体配列を返す
	pokemonData := getPokemon()
	fmt.Println(pokemonData)
	// jsonにエンコードする
	jsonData, err := json.Marshal(pokemonData)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func apiPokemonColorHandler(w http.ResponseWriter, r *http.Request){
    // URLから色を抽出（パスパラメータの利用）なるほど😄😄😄apiのURIでは、パラメータは?color=1よりも、pokemons_color/1 とやってフロントエンドで呼び出した方が設計上断然わかりやすいね。
    pathParts := strings.Split(r.URL.Path, "/")
    if len(pathParts) < 4 {
        http.Error(w, "Invalid URL or color not provided", http.StatusBadRequest)
        return
    }
    color := pathParts[3] // ex: "/api/pokemons_color/2/" から '2' を取得
    colorNum, err := strconv.Atoi(strings.TrimSpace(color))
    if err != nil {
        http.Error(w, "Invalid color format", http.StatusBadRequest)
        return
    }
    if colorNum < 0 || colorNum > 2 {
        http.Error(w, "Color number out of allowed range", http.StatusBadRequest)
        return
    }

	pokemonData := getPokemonsByColor(colorNum)

	// jsonにエンコードする
	jsonData, err := json.Marshal(pokemonData)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}


func StartWebServer() error {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/api/pokemons/", apiPokemonHandler)
	http.HandleFunc("/api/pokemons_color/", apiPokemonColorHandler)

	return http.ListenAndServe(":8080", nil)
}

