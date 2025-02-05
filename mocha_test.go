package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitorsalgado/mocha/v3"
	"github.com/vitorsalgado/mocha/v3/expect"
	"github.com/vitorsalgado/mocha/v3/reply"
)

func Test_Mocha(t *testing.T) {
	j, err := json.Marshal(RespBody{Results: []Pokemon{{Name: "Charizard"}}})
	assert.Nil(t, err)

	m := mocha.New(t)
	m.Start()

	scoped := m.AddMocks(mocha.Get(expect.URLPath("/api/v2/pokemon")).
		Reply(reply.OK().BodyString(string(j))))

	p, err := FetchPokemon(m.URL())
	fmt.Println(m.URL())
	assert.Nil(t, err)
	assert.True(t, scoped.Called())
	assert.Equal(t, p[0].Name, "Charizard")
}
