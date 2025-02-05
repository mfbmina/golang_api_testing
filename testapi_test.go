package main

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
)

func Test_TestAPI(t *testing.T) {
	j, err := json.Marshal(RespBody{Results: []Pokemon{{Name: "Charizard"}}})
	assert.Nil(t, err)

	defer apitest.NewMock().
		Get("https://pokeapi.co/api/v2/pokemon").
		RespondWith().
		Body(string(j)).
		Status(http.StatusOK).
		EndStandalone()()

	p, err := FetchPokemon(URL)
	assert.Nil(t, err)
	assert.Equal(t, p[0].Name, "Charizard")
}
