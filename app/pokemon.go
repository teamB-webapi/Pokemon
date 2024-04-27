package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"github.com/mtslzr/pokeapi-go"
	"strconv"

)
//何個取得してシャッフルするか
const limit int = 30
//javascriptに送る数
const sent int = 10


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





func getPokemon() []Pokemon{

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
		Height int `json:"height"`
		Weight int `json:"weight"`


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
	pokemon.Sprite = response.Sprites.Frontdefault

	return pokemon
}

func shuffle(arr []Pokemon) []Pokemon{
	for i := 0; i < len(arr); i++{
		r := rand.Intn(len(arr)-i) + i


		temp := arr[r]
		arr[r] = arr[i]
		arr[i] = temp
	}
	return arr
}




// 機能２ ブラウザからクライエントが入力した整数を使ってポケモンデータをフィルタリング。身長を降順にソートした上で配列を返す
// →機能２は機能１をそのまま使いたい→機能１で作ったポケモン構造体を持つ配列を使ってフィルタリング・ソーティングをするといいと思う

type PokemonOrigin struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	BaseExperience int    `json:"base_experience,omitempty"`
	Height         int    `json:"height,omitempty"`
	IsDefault      bool   `json:"is_default,omitempty"`
	Order          int    `json:"order,omitempty"`
	Weight         int    `json:"weight,omitempty"`
	Abilities      []struct {
		IsHidden bool `json:"is_hidden,omitempty"`
		Slot     int  `json:"slot,omitempty"`
		Ability  struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"ability,omitempty"`
	} `json:"abilities,omitempty"`
	Forms []struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"forms,omitempty"`
	GameIndices []struct {
		GameIndex int `json:"game_index,omitempty"`
		Version   struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"version,omitempty"`
	} `json:"game_indices,omitempty"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"item,omitempty"`
		VersionDetails []struct {
			Rarity  int `json:"rarity,omitempty"`
			Version struct {
				Name string `json:"name,omitempty"`
				URL  string `json:"url,omitempty"`
			} `json:"version,omitempty"`
		} `json:"version_details,omitempty"`
	} `json:"held_items,omitempty"`
	LocationAreaEncounters string `json:"location_area_encounters,omitempty"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"move,omitempty"`
		VersionGroupDetails []struct {
			LevelLearnedAt int `json:"level_learned_at,omitempty"`
			VersionGroup   struct {
				Name string `json:"name,omitempty"`
				URL  string `json:"url,omitempty"`
			} `json:"version_group,omitempty"`
			MoveLearnMethod struct {
				Name string `json:"name,omitempty"`
				URL  string `json:"url,omitempty"`
			} `json:"move_learn_method,omitempty"`
		} `json:"version_group_details,omitempty"`
	} `json:"moves,omitempty"`
	Species struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"species,omitempty"`
	Sprites struct {
		BackDefault      string `json:"back_default,omitempty"`
		BackFemale       any    `json:"back_female,omitempty"`
		BackShiny        string `json:"back_shiny,omitempty"`
		BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
		FrontDefault     string `json:"front_default,omitempty"`
		FrontFemale      any    `json:"front_female,omitempty"`
		FrontShiny       string `json:"front_shiny,omitempty"`
		FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
		Other            struct {
			DreamWorld struct {
				FrontDefault string `json:"front_default,omitempty"`
				FrontFemale  any    `json:"front_female,omitempty"`
			} `json:"dream_world,omitempty"`
			Home struct {
				FrontDefault     string `json:"front_default,omitempty"`
				FrontFemale      any    `json:"front_female,omitempty"`
				FrontShiny       string `json:"front_shiny,omitempty"`
				FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
			} `json:"home,omitempty"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default,omitempty"`
				FrontShiny   string `json:"front_shiny,omitempty"`
			} `json:"official-artwork,omitempty"`
			Showdown struct {
				BackDefault      string `json:"back_default,omitempty"`
				BackFemale       any    `json:"back_female,omitempty"`
				BackShiny        string `json:"back_shiny,omitempty"`
				BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
				FrontDefault     string `json:"front_default,omitempty"`
				FrontFemale      any    `json:"front_female,omitempty"`
				FrontShiny       string `json:"front_shiny,omitempty"`
				FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
			} `json:"showdown,omitempty"`
		} `json:"other,omitempty"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault  string `json:"back_default,omitempty"`
					BackGray     string `json:"back_gray,omitempty"`
					FrontDefault string `json:"front_default,omitempty"`
					FrontGray    string `json:"front_gray,omitempty"`
				} `json:"red-blue,omitempty"`
				Yellow struct {
					BackDefault  string `json:"back_default,omitempty"`
					BackGray     string `json:"back_gray,omitempty"`
					FrontDefault string `json:"front_default,omitempty"`
					FrontGray    string `json:"front_gray,omitempty"`
				} `json:"yellow,omitempty"`
			} `json:"generation-i,omitempty"`
			GenerationIi struct {
				Crystal struct {
					BackDefault  string `json:"back_default,omitempty"`
					BackShiny    string `json:"back_shiny,omitempty"`
					FrontDefault string `json:"front_default,omitempty"`
					FrontShiny   string `json:"front_shiny,omitempty"`
				} `json:"crystal,omitempty"`
				Gold struct {
					BackDefault  string `json:"back_default,omitempty"`
					BackShiny    string `json:"back_shiny,omitempty"`
					FrontDefault string `json:"front_default,omitempty"`
					FrontShiny   string `json:"front_shiny,omitempty"`
				} `json:"gold,omitempty"`
				Silver struct {
					BackDefault  string `json:"back_default,omitempty"`
					BackShiny    string `json:"back_shiny,omitempty"`
					FrontDefault string `json:"front_default,omitempty"`
					FrontShiny   string `json:"front_shiny,omitempty"`
				} `json:"silver,omitempty"`
			} `json:"generation-ii,omitempty"`
			GenerationIii struct {
				Emerald struct {
					FrontDefault string `json:"front_default,omitempty"`
					FrontShiny   string `json:"front_shiny,omitempty"`
				} `json:"emerald,omitempty"`
				FireredLeafgreen struct {
					BackDefault  string `json:"back_default,omitempty"`
					BackShiny    string `json:"back_shiny,omitempty"`
					FrontDefault string `json:"front_default,omitempty"`
					FrontShiny   string `json:"front_shiny,omitempty"`
				} `json:"firered-leafgreen,omitempty"`
				RubySapphire struct {
					BackDefault  string `json:"back_default,omitempty"`
					BackShiny    string `json:"back_shiny,omitempty"`
					FrontDefault string `json:"front_default,omitempty"`
					FrontShiny   string `json:"front_shiny,omitempty"`
				} `json:"ruby-sapphire,omitempty"`
			} `json:"generation-iii,omitempty"`
			GenerationIv struct {
				DiamondPearl struct {
					BackDefault      string `json:"back_default,omitempty"`
					BackFemale       any    `json:"back_female,omitempty"`
					BackShiny        string `json:"back_shiny,omitempty"`
					BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					FrontFemale      any    `json:"front_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
				} `json:"diamond-pearl,omitempty"`
				HeartgoldSoulsilver struct {
					BackDefault      string `json:"back_default,omitempty"`
					BackFemale       any    `json:"back_female,omitempty"`
					BackShiny        string `json:"back_shiny,omitempty"`
					BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					FrontFemale      any    `json:"front_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
				} `json:"heartgold-soulsilver,omitempty"`
				Platinum struct {
					BackDefault      string `json:"back_default,omitempty"`
					BackFemale       any    `json:"back_female,omitempty"`
					BackShiny        string `json:"back_shiny,omitempty"`
					BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					FrontFemale      any    `json:"front_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
				} `json:"platinum,omitempty"`
			} `json:"generation-iv,omitempty"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackDefault      string `json:"back_default,omitempty"`
						BackFemale       any    `json:"back_female,omitempty"`
						BackShiny        string `json:"back_shiny,omitempty"`
						BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
						FrontDefault     string `json:"front_default,omitempty"`
						FrontFemale      any    `json:"front_female,omitempty"`
						FrontShiny       string `json:"front_shiny,omitempty"`
						FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
					} `json:"animated,omitempty"`
					BackDefault      string `json:"back_default,omitempty"`
					BackFemale       any    `json:"back_female,omitempty"`
					BackShiny        string `json:"back_shiny,omitempty"`
					BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					FrontFemale      any    `json:"front_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
				} `json:"black-white,omitempty"`
			} `json:"generation-v,omitempty"`
			GenerationVi struct {
				OmegarubyAlphasapphire struct {
					FrontDefault     string `json:"front_default,omitempty"`
					FrontFemale      any    `json:"front_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
				} `json:"omegaruby-alphasapphire,omitempty"`
				XY struct {
					FrontDefault     string `json:"front_default,omitempty"`
					FrontFemale      any    `json:"front_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
				} `json:"x-y,omitempty"`
			} `json:"generation-vi,omitempty"`
			GenerationVii struct {
				Icons struct {
					FrontDefault string `json:"front_default,omitempty"`
					FrontFemale  any    `json:"front_female,omitempty"`
				} `json:"icons,omitempty"`
				UltraSunUltraMoon struct {
					FrontDefault     string `json:"front_default,omitempty"`
					FrontFemale      any    `json:"front_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
				} `json:"ultra-sun-ultra-moon,omitempty"`
			} `json:"generation-vii,omitempty"`
			GenerationViii struct {
				Icons struct {
					FrontDefault string `json:"front_default,omitempty"`
					FrontFemale  any    `json:"front_female,omitempty"`
				} `json:"icons,omitempty"`
			} `json:"generation-viii,omitempty"`
		} `json:"versions,omitempty"`
	} `json:"sprites,omitempty"`
	Cries struct {
		Latest string `json:"latest,omitempty"`
		Legacy string `json:"legacy,omitempty"`
	} `json:"cries,omitempty"`
	Stats []struct {
		BaseStat int `json:"base_stat,omitempty"`
		Effort   int `json:"effort,omitempty"`
		Stat     struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"stat,omitempty"`
	} `json:"stats,omitempty"`
	Types []struct {
		Slot int `json:"slot,omitempty"`
		Type struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"type,omitempty"`
	} `json:"types,omitempty"`
	PastTypes []struct {
		Generation struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"generation,omitempty"`
		Types []struct {
			Slot int `json:"slot,omitempty"`
			Type struct {
				Name string `json:"name,omitempty"`
				URL  string `json:"url,omitempty"`
			} `json:"type,omitempty"`
		} `json:"types,omitempty"`
	} `json:"past_types,omitempty"`
}

type Pokemon struct{
	Name           string `json:"name,omitempty"`
	Height         int    `json:"height,omitempty"`
	Weight         int    `json:"weight,omitempty"`
	Ability 	   string `json:"ability,omitempty"`
	Sprite 		   string `json:"sprite,omitempty"`
}


// Resource関数ではポケモンの個数を選択して構造体を持ったスライスを取得する
// そのスライスをループする❶
// ❶デコード(全ての情報を持ったデータを構造体(PokemonOrigin)にはめ込む)
// ❷❶の構造体のデータから必要なフィールドだけを使って、構造体(Pokemon)を作る
// ❸Pokemonタイプのスライスに❷を格納
// ❸をヘッドにwrite

// height以下の値を持った構造体のみを持った配列を返す
func getPokemonsByHeight() []Pokemon{
	height := 30
	// lには構造体Resultのデータが入ったスライスがくる→ResultはName, URLの2つのフィールドを持つ
	// l,err := pokeapi.Resource("pokemon")
	l,err := pokeapi.Resource("pokemon-color")
	if err != nil{
		log.Fatalln(err)
	}

	// ポケモン構造体の配列を作る
	var pokemons = []Pokemon{}

	fmt.Println(len(l.Results))

	for i := 0; i < len(l.Results); i++{
		resp, err := http.Get(l.Results[i].URL)
		if err != nil{
			log.Fatal(err)
		}
		
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil{
			log.Fatal(err)
		}
	
		var pokemonOriginal = PokemonOrigin{}
	
		if err := json.Unmarshal([]byte(body), &pokemonOriginal); err != nil{
			log.Fatal(err)
		}

		fmt.Println(pokemonOriginal)

		if(pokemonOriginal.Height > height){
			continue
		}

		var pokemon = Pokemon{}

		pokemon.Name = pokemonOriginal.Name
		pokemon.Height = pokemonOriginal.Height
		pokemon.Weight = pokemonOriginal.Weight
		pokemon.Ability = pokemonOriginal.Abilities[0].Ability.Name
		pokemon.Sprite = pokemonOriginal.Sprites.FrontDefault

		pokemons = append(pokemons, pokemon)
	}

	// pokemonsをheightの値をもとに昇順にソート
	sort.Slice(pokemons, func(i,j int) bool { return pokemons[i].Height < pokemons[j].Height })

	return pokemons
}


type PokemonColor struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonSpecies []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon_species"`
}

// カラーズのエンドポイントからpokemon_species構造体の中にurlがあり、それが個体情報をとるエンドポイント位になっている→そのエンドポイントを使ってオリジナル構造体→ポケモン構造体って感じでやっていけばOK
func getPokemonsByColor(colorCode int) []Pokemon{
	color := "black"
	if colorCode == 0{
		color = "red"
	}else if colorCode == 1{
		color = "green"
	}else{
		color = "blue"
	}

	pokemonsByColor,err := pokeapi.PokemonColor(color,)
	if err != nil{
		log.Fatalln(err)
	}

	var names = []string{}

	// ここで取ってくるデータ数を制限する
	for i := 0; i < min(5, len(pokemonsByColor.PokemonSpecies)); i++{
		names = append(names, pokemonsByColor.PokemonSpecies[i].Name)
	}

	// ポケモン構造体の配列を作る
	var pokemons = []Pokemon{}

	baseURL := "https://pokeapi.co/api/v2/pokemon/"

	for i := 0; i < len(names); i++{
		// URLはポケモンの個体情報を取得するエンドポイント
		resp, err := http.Get(baseURL + names[i])
		if err != nil{
			log.Fatal(err)
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil{
			log.Fatal(err)
		}
	
		var pokemonOriginal = PokemonOrigin{}
	
		if err := json.Unmarshal([]byte(body), &pokemonOriginal); err != nil{
			log.Fatal(err)
		}

		var pokemon = Pokemon{}

		pokemon.Name = pokemonOriginal.Name
		pokemon.Height = pokemonOriginal.Height
		pokemon.Weight = pokemonOriginal.Weight
		pokemon.Ability = pokemonOriginal.Abilities[0].Ability.Name
		pokemon.Sprite = pokemonOriginal.Sprites.FrontDefault

		pokemons = append(pokemons, pokemon)
	}

	return pokemons
}
