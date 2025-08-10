package main

import "fmt"

func main() {

	// Printing F(x)
	fmt.Print("Hello")
	fmt.Print("World")
	fmt.Print(12,29)

	fmt.Println("Hello")
	fmt.Println("World")
	fmt.Println(12,29)

	name:="Skyy"
	age:=29

	fmt.Printf("Name: %s, Age: %d\n",name,age)
	fmt.Printf("Binary: %b, Hex: %X\n",age,age)

	// Formatting F(x)
	s:= fmt.Sprint("Hi", "World",123,456)
	s1:= fmt.Sprintln("Hi", "World",123,456)
	fmt.Println(s)
	fmt.Println(s1)

	sf:=fmt.Sprintf("Name: %s, Age:%d",name,age)
	fmt.Println(sf)
	fmt.Println(sf)

	// Scanning F(x)
	var batName string
	var batAge int

	fmt.Print("🦇Enter Batman's name and age:")
	//fmt.Scan(&batName, &batAge) // Waits for all the values
	//fmt.Scanln(&batName, &batAge) // Takes empty values
	//fmt.Scanf("%s %d",&batName,&batAge) // Expects Formatting
	fmt.Printf("Name: %s, Age:%d\n", batName, batAge)

	// Error F(x)
	err := CheckAge(11)
	if err!=nil{
		fmt.Println("🔴 Error:",err)
	}


	
}

func CheckAge(age int)error{
	if age<18{
		return fmt.Errorf("Age %d is too young to drive!",age)
	}
	return nil
}