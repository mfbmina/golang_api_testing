package main

import (
	"encoding/json"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func Test_Gock(t *testing.T) {
	defer gock.Off()

	j, err := json.Marshal(RespBody{Results: []Pokemon{{Name: "Charizard"}}})
	assert.Nil(t, err)

	gock.New("https://pokeapi.co").
		Get("/api/v2/pokemon").
		Reply(200).
		JSON(j)

	p, err := FetchPokemon(URL)
	assert.Nil(t, err)
	assert.Equal(t, p[0].Name, "Charizard")
}
