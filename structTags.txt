package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Pokemon struct {
	Name      string `json:"name,omitempty" db:"pokemon_name" xml:"poke_name"` 
	//chaining multiple struct-tags
	Type      string `json:"type,omitempty"`
	PokeDexNo int    `json:"no.,omitempty"`
	OmitValue string `json:"-"` // always omit in final result
}

func main() {

	pikachu := Pokemon{Name: "Pikachu", Type: "⚡", PokeDexNo: 25}
	pikachuJson, err := json.Marshal(pikachu)
	if err!=nil{
		log.Fatal("🔴 ERROR marshalling struct:",err)
	}
	fmt.Println("\nPikachu JSON-Data:",string(pikachuJson))

	// omitempty
	charmander := Pokemon{Name: "Charmander", Type: "🔥"}
	charmanderJson,err:=json.Marshal(charmander)
	if err!=nil{
		log.Fatal("🔴 ERROR marshalling struct:",err)
	}
	fmt.Println("\nCharmander JSON-Data:",string(charmanderJson))

}