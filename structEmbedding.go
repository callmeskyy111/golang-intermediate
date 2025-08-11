package main

// field promotion

import "fmt"

type Pokemon struct {
	name string
	no   int
}

type Pikachu struct {
	//PokemonInfo Pokemon // named field, not anonymous
	Pokemon // embedded struct
	trainer string
	weight  float64
}

func (p Pokemon) attack(){
	fmt.Printf("Hi, the Pokemon is %s, PokeDex_No: %d\n",p.name,p.no)
}

func (p Pikachu) attack(){
	fmt.Printf("Hi, the Pokemon is %s, PokeDex_No: %d, Weighing %.2f kgs, belonging to %s",p.name,p.no,p.weight,p.trainer)
}

func main() {
	pikachu := Pikachu{
		Pokemon: Pokemon{name: "Sparky ⚡", no: 25},
		trainer: "Ritchie 🧢",
		weight:  6.00,
	}

	fmt.Println("NAME:",pikachu.name) 	// Accessing embedded struct{ } directly
	fmt.Println("PokeDex_No.",pikachu.no) 	// Accessing embedded struct{ } directly
	fmt.Println("Trainer:",pikachu.trainer)
	fmt.Println("Weight(kgs):",pikachu.weight)

	pikachu.attack()

	// Method overriding
}