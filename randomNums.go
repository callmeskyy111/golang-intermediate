package main

import (
	"fmt"
	"math/rand"
	"time"
)

// random numbers

func main() {
fmt.Println(rand.Intn(101))
fmt.Println(rand.Intn(6)+5)

// fixed output - seed 🌱
val:= rand.New(rand.NewSource(29))
val1:= rand.New(rand.NewSource(time.Now().Unix()))
fmt.Println(val.Intn(101)+2)
fmt.Println(val1.Intn(101)+2)

// random float nums
fmt.Println(rand.Float64()) // between 0.0 and 1.0 (excl.)

// Dice-Game 🎮🎲

for{
	// show the menu 📃
	fmt.Println("Welcome to the Dice-Game 🎮")
	fmt.Println("1. Roll the dice 🎲")
	fmt.Println("2. Exit.. 🚪")
	fmt.Println("Enter your choice (1️⃣  or 2️⃣): ")

	var choice int

	_,err:= fmt.Scan(&choice)
	if err!= nil || (choice!=1 && choice !=2){
		fmt.Println("🔴Invalid Choice. Please choose 1 or 2")
		continue
	}
	if choice == 2{
		fmt.Println("Thanks for playing.. G'bye! 💖")
		break
	}

	die1:= rand.Intn(6)+1
	die2:= rand.Intn(6)+1

	// show the result
	fmt.Printf("You rolled a %d and a %d.🟢\n",die1,die2)
	fmt.Println("Total:",die1+die2)

	// ask if the user wants to roll again
	fmt.Print("Do you wanna roll again? (y/n)")
	var rollAgain string
	_,err=fmt.Scan(&rollAgain)
	if err!=nil || (rollAgain != "y" && rollAgain != "n"){
	fmt.Println("🔴 Invalid input.. Assuming 'no'")
	break
	}
	if rollAgain == "n"{
		fmt.Println("Thanks for playing.. G'bye! 💖")
	}
}


}