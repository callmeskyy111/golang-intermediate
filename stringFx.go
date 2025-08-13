package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

// strings - sequence of bytes
// Strings are immutable

func main() {
	str:= "Hello there"
	fmt.Println(len(str))

	fmt.Println(str[2]) // 108 -> ASCII Value
	fmt.Println(str[4]) // 111 -> ASCII Value

	// Sub-str
	fmt.Println(str[0:5]) // "Hello"

	// Str-conversion
	num:=18
	str1:=strconv.Itoa(num)
	fmt.Println(len(str1)) // 2

	// Str-splitting
	fruits:= "apple🍎, orange🍊, bananas🍌"
	spetsnaz:= "GRU🧠-Vympel⚔️-Alpha🛡️-SSO🥷🏻-Frogmen⚓"
	jets:=[]string{"Su-27","Mig-29","Su-57","Su-35"}

	parts:= strings.Split(fruits, ",")
	forces:= strings.Split(spetsnaz, "-")
	fmt.Println(parts)
	fmt.Println(forces)

	joinedJets:= strings.Join(jets," 🔥 ")
	fmt.Println(joinedJets)  // slice to str

	fmt.Println(strings.Contains(jets[1], "Mig-29"))
	fmt.Println(strings.Contains(fruits, "zzz"))

	replaced:= strings.Replace(fruits,"orange🍊","cherry🍒",1)
	fmt.Println(replaced)

	// Now, trimming leading and trailing white-spaces
	strwspace:= "     Hallo Leute  "
	fmt.Println(strwspace)
	fmt.Println(strings.TrimSpace(strwspace))

	// Changing the case
	fmt.Println(strings.ToLower(strwspace))
	fmt.Println(strings.ToUpper(strwspace))

	// Repeat
	fmt.Println(strings.Repeat("Scramble!⚠️ ",3))

	// Counting
	fmt.Println(strings.Count(fruits,"a")) // 5

	// Prefix & Suffix
	fmt.Println(strings.HasPrefix("Batman🦇","Ba")) // true
	fmt.Println(strings.HasSuffix("Spiderman🕷️","an🕷️")) // true

	// Intro to the Regexp Package - Brief Overview
	mySlogan:= "I'm Batman 123 And 55"
	herosSlogan:= "তোমরা আমাকে রক্ত দাও, আমি তোমাদের স্বাধীনতা দেবো✊🏻"

	rg:= regexp.MustCompile(`\d+`)
	allMatches:= rg.FindAllString(mySlogan,-1) // -1 -> all str
	fmt.Println(allMatches)

	// UTF8
	// Golang by default supports other languages & scripts
	fmt.Println(utf8.RuneCountInString(herosSlogan)) // 50

	//STR BUILDER - Efficent than concat/+ ops. or Sprintf()

	var builder strings.Builder
	var runeBuilder strings.Builder

	// Write some strs
	builder.WriteString("Guten")
	builder.WriteString(" ")
	builder.WriteString("Tag!")
	builder.WriteString("🌞")

	// Another method
	runeBuilder.WriteRune(' ')
	runeBuilder.WriteString("Wie geht's Brudi?👋🏻")

	runeRes:= runeBuilder.String()
	fmt.Println("Built-up str from RUNES(chars):",runeRes)


	// Convert builder -> str
	builderRes:=builder.String()
	fmt.Println("Built-up str:",builderRes)

	// Reset the builder
	builder.Reset()
	builder.WriteString("  New str. Yaayyy")
	builder.WriteString(" RESET done✔️")
	resetRes:= builder.String()
	fmt.Println("New built-up str, after RESET:", resetRes)




}