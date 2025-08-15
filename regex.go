package main

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Println("He said, \"I am great\"") // He said, "I am great"
	fmt.Println(`He said, "I am great"`) // He said, "I am great"

	// Compile a regex pattern to match email address
	re:=regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// Testing.. 🧪
	email1:="user@email.com"
	email2:="invalid_email"

	// Match
	fmt.Println("Email:",re.MatchString(email1)) // Email: true
	fmt.Println("Email:",re.MatchString(email2)) // Email: false

	// Capturing groups ()
	// Compile a regex pattern to capture date components

	re1:= regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	// Testing.. 🧪
	date:= "2024-09-26"

	// Find all submatches
	submatches:= re1.FindStringSubmatch(date)

	// All the subStrs.
	fmt.Println(submatches) // [2024-09-26 2024 09 26]
	fmt.Println(submatches[0]) // 2024-09-26
	fmt.Println(submatches[1]) // 2024
	fmt.Println(submatches[2]) // 09
	fmt.Println(submatches[3]) // 26

	// Replace characters in a TARGET/SOURCE str.
	srcStr:= "Guten Tag Leute!"
	re2:= regexp.MustCompile(`[aeiou]`) // Replacing the vowels
	result:= re2.ReplaceAllString(srcStr,"*")
	fmt.Println(result) // G*t*n T*g L**t*!

	// Flags And Options
	// i - case insensitive
	// m - multi line model
	// s - dot matches all
	re3:= regexp.MustCompile(`(?i)go`)
	re4:= regexp.MustCompile(`go`)

	// Testing.. 🧪
	testTxt:= "GoLang is awesome!"

	// Match
	fmt.Println("Match:",re3.MatchString(testTxt)) // Match: true
	fmt.Println("Match:",re4.MatchString(testTxt)) // Match: false

}