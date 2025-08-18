package main

import (
	"encoding/base64"
	"fmt"
)

// Encoding - Translating data from it's original format to another format, that can be used by other systems, applictaions, protocols.

// ⚠️ base64 encoding != ENCRYPTION

func main() {
	data:= []byte("Hello, Base64 Encoding")

	// Encode Base64

	encoded:= base64.StdEncoding.EncodeToString(data)
	fmt.Println("Encoded string:",encoded) // base64 Encoded Str.

	// Decode from Base64
	decoded,err:=base64.StdEncoding.DecodeString(encoded) // Encoded string: SGVsbG8sIEJhc2U2NCBFbmNvZGluZw==
	if err!=nil{
		fmt.Println("🔴ERROR decoding:",err)
		return
	}
	fmt.Println("Decoded string:",string(decoded)) // Decoded string: Hello, Base64 Encoding

	// URL-Safe encoding, avoid '/' and '+'

	urlSafeEncoded:=base64.URLEncoding.EncodeToString(data) 
	fmt.Println("URL-Safe encoded:",urlSafeEncoded) // removes + and /




}