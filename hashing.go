package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

// Hashing f(x) -> irreversible algos. (1 way) // Cryprographic Hash f(x)
// Produces fixed sized output irrespective of input size
// Good Hash f(x)s produce less collisions
// Small changes cause massive effects in the HASH F(x)s
// Passwords need to/must be HASHED first before storing them in DBs 🛡️

// Ex. of hashing algorithms: SHA-256, SHA-512 etc.

func main() {
	password:="password123"
	hashed:= sha256.Sum256([]byte(password)) // accepts bytes
	hashed512:= sha512.Sum512([]byte(password)) // accepts bytes

	//byte values
	fmt.Println(hashed)
	fmt.Println(hashed512)

	fmt.Println("password:",password)
	fmt.Println("hashed-password:",hashed)
	fmt.Println("hashed512-password:",hashed512)

	// converting to HEX
	fmt.Printf("SHA-256 Hash hex val: %x\n",hashed) // 256 bit str.
	fmt.Printf("SHA-512 Hash hex val: %x\n",hashed512) // 512 bit str.

}