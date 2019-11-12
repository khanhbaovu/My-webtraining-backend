package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
  //  "log"
    "net/http"
//    "os"
)

type Response struct {
  Area string `json:"name"`
  Pokemon []Pokemons `json:"pokemon_entries"`
}

type Pokemons struct {
  Index int `json:"entry_number"`
  PokemonType PokemonName  `json:"pokemon_species"`
}

type PokemonName struct {
  Name string `json:"name"`
}

func main() {
    response,err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

    if err != nil {
      panic(err)
    }

    responseData,err1 := ioutil.ReadAll(response.Body)

    if err1 != nil {
      panic(err)
    }

    var responseObject Response

    json.Unmarshal(responseData, &responseObject)

    fmt.Println(responseObject)


}
