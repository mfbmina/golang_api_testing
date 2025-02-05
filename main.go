package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RespBody struct {
	Results []Pokemon `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
}

const URL = "https://pokeapi.co"

func main() {
	pkmns, err := FetchPokemon(URL)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, pkmn := range pkmns {
		fmt.Println(pkmn.Name)
	}
}

func FetchPokemon(u string) ([]Pokemon, error) {
	r, err := http.Get(fmt.Sprintf("%s/api/v2/pokemon", u))
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()
	resp := RespBody{}
	err = json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return resp.Results, nil
}
