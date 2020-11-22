package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Pokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	Forms          []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	Height                 int    `json:"height"`
	ID                     int    `json:"id"`
	Name                   string `json:"name"`
	Order                  int    `json:"order"`
	Species                struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

var command = &cobra.Command{
	Use:   "pokedex [ID]",
	Args:  cobra.MinimumNArgs(1),
	Short: "pokedex ID",
	Run: func(cmd *cobra.Command, args []string) {
		pokemon := fetchPokemon(args[0])
		outputPokemon(pokemon)
	},
}

func init() {
	rootCmd.AddCommand(command)
}

func fetchPokemon(id string) (pokemon *Pokemon) {
	req, _ := http.NewRequest("GET", "https://pokeapi.co/api/v2/pokemon/"+id, nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &pokemon)

	return pokemon
}

func outputPokemon(pokemon *Pokemon) {
	fmt.Println("##############ポケモン図鑑################")
	fmt.Println("ポケモン図鑑No：" + strconv.Itoa(pokemon.ID))
	fmt.Println("名前：" + pokemon.Forms[0].Name)
	fmt.Println("----------------------------------------")
	for i := 0; i < len(pokemon.Types); i++ {
		fmt.Println("タイプ" + strconv.Itoa(i+1) + "：" + pokemon.Types[i].Type.Name)
	}
	fmt.Println("----------------------------------------")
	for i := 0; i < len(pokemon.Types); i++ {
		fmt.Println("能力" + strconv.Itoa(i+1) + "：" + pokemon.Abilities[i].Ability.Name)
	}
	fmt.Println("----------------------------------------")
	weight := float64(pokemon.Weight) / 10.0
	height := float64(pokemon.Height) / 10.0
	fmt.Println("重さ：" + strconv.FormatFloat(weight, 'f', 2, 64) + "kg")
	fmt.Println("高さ：" + strconv.FormatFloat(height, 'f', 2, 64) + "m")
}
