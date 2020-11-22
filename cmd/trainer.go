package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

type Trainer struct {
	Name      string   `json:"name"`
	PokemonID []string `json:"pokemon_id"`
}

var trainerCmd = &cobra.Command{
	Use:   "trainer [NAME]",
	Args:  cobra.MinimumNArgs(1),
	Short: "trainer name",
	Run: func(cmd *cobra.Command, args []string) {
		fetchTrainer(args[0])
	},
}

func init() {
	rootCmd.AddCommand(trainerCmd)
}

func fetchTrainer(name string) {
	bytes, err := ioutil.ReadFile("./trainer.json")
	if err != nil {
		log.Fatal(err)
	}
	var trainer []Trainer
	if err := json.Unmarshal(bytes, &trainer); err != nil {
		log.Fatal(err)
	}
	for _, p := range trainer {
		if name == p.Name {
			for i := 0; i < 4; i++ {
				pokemon := fetchPokemon(p.PokemonID[i])
				outputPokemon(pokemon)
			}
		}
	}
}
