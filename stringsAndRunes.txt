package main

import (
	"fmt"
	"unicode/utf8"
)

// string - a sequence of bytes, bytes - unsigned uint8 values, strings are immutable
// rune - an integer value that represents a CHAR
// runes - fixed amount in memory, int32, occupy 4 bytes
// strings are made up of RUNEs.. They are kinda like arrays, strs have len()

// runes in GoLang and chars. in C are kinda similar, with minor differences
// handling runes, UNICODE etc., easier in GoLang, compared to C
// GOLANG natively supports UNICODE
// GoLang has better and efficient libraries. ex- unicode/utf8

func main() {
	message:= "Hello, GoLang!"
	message2:= "Hello, \tGoLang TAB!"
	message3:= "Hello, \rGoLang r!"
	rawMessage:= `Hello\n Go!`

	fmt.Println(message)
	fmt.Println(message2)
	fmt.Println(message3)
	fmt.Println(rawMessage)

	fmt.Println(len(message)) // 14

	// Extracting CHARs from strings, kinda like arrays
	fmt.Println("1st character in 'message2':",message2[0]) // 72 (ASCII)

	greeting:= "Guten Tag"
	name:= "Dominik"
	fmt.Println(greeting+" "+name+"!")

	// Str comparision (LEXICOGRAPHICAL, kinda like word-orders in dictionaries)
	str1:="Apple 🍎"
	str2:="Banana 🍌"
	str3:="App 📱"
	str4:= "ab 🔤"

	fmt.Println(str1<str2) // true
	fmt.Println(str3<str1) // true
	fmt.Println(str4<str1) // false : a(ASCII 97), A(ASCII 65)

	// Str iteration

	for i,char:= range message2{
		fmt.Printf("CHAR. at idx %d: %c, HEX: %x, ASCII:%v\n", i, char,char,char)
	}

	fmt.Println("Rune count:", utf8.RuneCount([]byte(greeting))) // 9

	//! Now, runes
	var ch rune='s'
	bch:= 'অ' // Bengali character
	fmt.Println(ch) // 115
	fmt.Println(bch) // 2437

	// Bengali characters
	bnRunes := []rune{'অ', 'আ', 'ই', 'উ', 'ঋ', 'এ', 'ঐ', 'ও', 'ঔ'}

	for _, r := range bnRunes {
		fmt.Printf("Character: %c | Rune: %U | Decimal: %d\n", r, r, r)
	}

	// Now.. Runes -> str conversion
	cstr:= string(ch)
	fmt.Println("Converted str:",cstr)

	// Confirming with TypeCheck
	fmt.Printf("Type of cstr is %T\n",cstr) // string

	// UNICODE support
	const RajanidGondha = "রজনীগন্ধা" // Bengali for TUBEROSE
	fmt.Println(RajanidGondha) // রজনীগন্ধা

	shantiNiketon:= "শান্তিনিকেতন" // Google it, if not known

	for _, runeVal:= range shantiNiketon{
		fmt.Printf("runeVal: %c, unicodeVal: %v\n",runeVal,runeVal)
	}

	// Emoji support too
	h:= '🎃'
	fmt.Printf("char: %c, unicodeVal: %v\n",h,h)
}

//! ALL OUTPUT:

// Hello, GoLang!
// Hello, 	GoLang TAB!
// Hello, 
// GoLang r!
// Hello\n Go!
// 14
// 1st character in 'message2': 72
// Guten Tag Dominik!
// true
// true
// false
// CHAR. at idx 0: H, HEX: 48, ASCII:72
// CHAR. at idx 1: e, HEX: 65, ASCII:101
// CHAR. at idx 2: l, HEX: 6c, ASCII:108
// CHAR. at idx 3: l, HEX: 6c, ASCII:108
// CHAR. at idx 4: o, HEX: 6f, ASCII:111
// CHAR. at idx 5: ,, HEX: 2c, ASCII:44
// CHAR. at idx 6:  , HEX: 20, ASCII:32
// CHAR. at idx 7: 	, HEX: 9, ASCII:9
// CHAR. at idx 8: G, HEX: 47, ASCII:71
// CHAR. at idx 9: o, HEX: 6f, ASCII:111
// CHAR. at idx 10: L, HEX: 4c, ASCII:76
// CHAR. at idx 11: a, HEX: 61, ASCII:97
// CHAR. at idx 12: n, HEX: 6e, ASCII:110
// CHAR. at idx 13: g, HEX: 67, ASCII:103
// CHAR. at idx 14:  , HEX: 20, ASCII:32
// CHAR. at idx 15: T, HEX: 54, ASCII:84
// CHAR. at idx 16: A, HEX: 41, ASCII:65
// CHAR. at idx 17: B, HEX: 42, ASCII:66
// CHAR. at idx 18: !, HEX: 21, ASCII:33
// Rune count: 9
// 115
// 2437
// Character: অ | Rune: U+0985 | Decimal: 2437
// Character: আ | Rune: U+0986 | Decimal: 2438
// Character: ই | Rune: U+0987 | Decimal: 2439
// Character: উ | Rune: U+0989 | Decimal: 2441
// Character: ঋ | Rune: U+098B | Decimal: 2443
// Character: এ | Rune: U+098F | Decimal: 2447
// Character: ঐ | Rune: U+0990 | Decimal: 2448
// Character: ও | Rune: U+0993 | Decimal: 2451
// Character: ঔ | Rune: U+0994 | Decimal: 2452
// Converted str: s
// Type of cstr is string
// রজনীগন্ধা
// runeVal: শ, unicodeVal: 2486
// runeVal: া, unicodeVal: 2494
// runeVal: ন, unicodeVal: 2472
// runeVal: ্, unicodeVal: 2509
// runeVal: ত, unicodeVal: 2468
// runeVal: ি, unicodeVal: 2495
// runeVal: ন, unicodeVal: 2472
// runeVal: ি, unicodeVal: 2495
// runeVal: ক, unicodeVal: 2453
// runeVal: ে, unicodeVal: 2503
// runeVal: ত, unicodeVal: 2468
// runeVal: ন, unicodeVal: 2472
// char: 🎃, unicodeVal: 127875