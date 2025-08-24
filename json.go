package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	Age       int    `json:"age,omitempty"`
	EmailAddress string `json:"email"`
	Address  Address `json:"address"`
}

type Address struct{
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

type Pokemon struct{
	Name string `json:"name"`
	Number int `json:"no."`
	Type string `json:"type,omitempty"`
	Trainer Trainer `json:"trainer"`
}

type Trainer struct{
	TrainerName string `json:"trainer_name,omitempty"`
	HomeTown string `json:"hometown,omitempty"`
}

func main() {
	john := Person{FirstName: "John", Age: 36, EmailAddress: "john@golang.com"}

		fmt.Println("\n----------- * MARSHALLING * -----------")
	// Marshalling
	jsonData,err:=json.Marshal(john)
	if err!=nil{
		fmt.Println("🔴 ERROR marshalling to JSON:",err)
		return
	}
	fmt.Println(string(jsonData))

	skyy := Person{FirstName:"Skyy", Age:29, EmailAddress: "skyy@golang.org", Address: Address{City: "Munich", State: "Bavaria"}}

	jsonData1,err:=json.Marshal(skyy)
	if err!=nil{
		fmt.Println("🔴 ERROR marshalling to JSON:",err)
		return
	}
	fmt.Println(string(jsonData1))

	// Now, unmarshalling
	fmt.Println("\n-------------- * UNMARSHALLING * --------------")
	treeckoJson:= `{"name":"Treecko","no.":252,"type":"🌿", "trainer":{"trainer_name":"Ash Ketchum", "hometown":"Pallet Town"}}`

	var treeckoDecoded Pokemon
	err= json.Unmarshal([]byte(treeckoJson),&treeckoDecoded)
	if err!=nil{
		fmt.Println("🔴 ERROR unmarshalling from JSON:",err)
		return
	}
	fmt.Println("\nUnmarshalled json-data:",treeckoDecoded)

	fmt.Println("Treecko's No:",treeckoDecoded.Number)
	fmt.Println("Treecko's Trainer:",treeckoDecoded.Trainer.TrainerName)
	fmt.Println("Treecko's Type:",treeckoDecoded.Type)

	fmt.Println("\n-------------- * LIST OF DATA * --------------")

	listOFCityState := []Address{
		{City: "Berlin_St.", State: "Berlin "},
		{City: "Sarasota", State: "Florida"},
		{City: "Hyderabad", State: "Telengana"},
		{City: "Denver", State: "Colorado"},
		{City: "Westland", State: "Michigan"},
	}

	fmt.Println("\nlistOFCityState:",listOFCityState)

	encodedList,err :=json.Marshal(listOFCityState)

	if err!=nil{
		fmt.Println("🔴 ERROR marshalling to JSON:",err)
		return
	}
	fmt.Println("\nJSON list: ",string(encodedList))

	// Handling unknown JSON-structures
	fmt.Println("\n-------- * UNKNOWN JSON-STRCUTURES * --------")

	unknownData:= `{"name":"Umbreon", "type":"🖤","no.":197, "trainer":{"trainer_name":"Gary", "hometown":"Pallet Town"}}`

	var data map[string]any

	err= json.Unmarshal([]byte(unknownData), &data)
	if err!=nil{
		log.Fatal("🔴 ERROR unmarshalling from JSON:",err)
	}
	fmt.Println("\nDecoded Unmarshalled-JSON:",data)
	fmt.Println("\nDecoded Unmarshalled-JSON:",data["trainer"])
	fmt.Println("\nDecoded Unmarshalled-JSON:",data["type"])

	




}