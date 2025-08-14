package main

// template strings on Golang
// Usage -  Populate/Generate HTML pages, email template, generating code, creating structured docs (invoices)

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {

	tmpl, err := template.New("example").Parse("Welcome , {{.name}}! How are you doing?\n")
	if err!=nil{
		panic(err.Error())
	}

	// Define data for the welcome message template
	data := map[string]any{
		"name":"John",
	}
	err=tmpl.Execute(os.Stdout, data)
	if err!=nil{
		panic(err.Error())
	}

	// Another example

	tmpl1:= template.Must(template.New("example").Parse("Welcome , {{.name}}! How are you doing?\n"))

	data1 := map[string]any{
		"name":"Skyy",
	}
	err=tmpl1.Execute(os.Stdout, data1)
	if err!=nil{
		panic(err.Error())
	}

	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	name,err:=reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Define named templates for diff. types, without manually writing them
	templates:= map[string]string{
		"welcome":"Herzlich Willkommen {{.name}}! Freut mich sehr 💖",
		"notification":"{{.nm}} , you have a new notification: {{.ntf}}",
		"error":"Oops! An ERROR occurred: {{.errorMsg}} ⚠️",
	}

	// Parse And Store templates
	parsedTemps:= make(map[string]*template.Template)

	for name,tmp:=range templates{
		parsedTemps[name] = template.Must(template.New(name).Parse(tmp))
	}

	// For-loop as a WHILE Loop
	for{
		// Show menu
		fmt.Println("\n📃Menu:")
		fmt.Println("1. Join 👤")
		fmt.Println("2. Get Notification 🔔")
		fmt.Println("3. Get Error 🔴")
		fmt.Println("4. Exit 🚪")
		fmt.Println("Choose an option: ")

		choice, _:= reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]any
		var tmplt *template.Template

		switch choice {
		case "1":
			tmplt = parsedTemps["welcome"]
			data = map[string]any{"name":name}
		case "2":
			fmt.Println("Enter your notification msg:")
			notification,_:= reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			
			tmplt= parsedTemps["notification"]
			data = map[string]any{"nm":name,"ntf":notification}
		case "3":
			fmt.Println("Enter your errorMessage:")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)

			tmplt= parsedTemps["error"]
			data = map[string]any{"nm":name,"errorMsg":errorMessage}
		case "4":
			fmt.Println("Exiting... ✔️")	
			return
		default:
			fmt.Println("Invalid Choice.. Please select a valid option!")	
			continue
		}

		// render and print the template to the console
		err:= tmplt.Execute(os.Stdout, data)
		if err!=nil{
			fmt.Println("Error executing template:",err)
		}



	}
}