package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HTTPtest(t *testing.T) {
	j, err := json.Marshal(RespBody{Results: []Pokemon{{Name: "Charizard"}}})
	assert.Nil(t, err)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v2/pokemon" {
			t.Errorf("Expected to request '/api/v2/pokemon', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(j))
	}))
	defer server.Close()

	p, err := FetchPokemon(server.URL)
	assert.Nil(t, err)
	assert.Equal(t, p[0].Name, "Charizard")
}
