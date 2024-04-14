package app

import (
	"encoding/json"
	"log"
	"net/http"
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


func apiPokemonHandler(w http.ResponseWriter, _ *http.Request) {
	// TODO: 変数pokemonDataには機能１である関数を実行してポケモンの構造体配列を返す
	pokemonData := sum()

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

	return http.ListenAndServe(":8080", nil)
}

