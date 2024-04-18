package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

	// jsonにエンコードする
	jsonData, err := json.Marshal(pokemonData)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func apiPokemonColorHandler(w http.ResponseWriter, r *http.Request){
	// クエリを取得
	color := r.URL.Query().Get("color")
	colorNum,err := strconv.Atoi(color)

	if err != nil || colorNum < 0 || colorNum > 2{
		APIError(w, "Invalid Value", http.StatusBadRequest)
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

