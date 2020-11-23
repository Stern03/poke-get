package cmd

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestFetchTrainer(t *testing.T) {
	input := "satoshi"
	result := fetchTrainer(input)
	expected := Trainer {
		Name:       "satoshi",
		PokemonID:   []string{"1", "6", "7", "25"},
	}
	assert.Equal(t, result, expected)
}