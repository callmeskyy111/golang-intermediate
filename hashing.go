package main

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

// Hashing f(x) -> irreversible algos. (1 way) // Cryprographic Hash f(x)
// Produces fixed sized output irrespective of input size
// Good Hash f(x)s produce less collisions
// Small changes cause massive effects in the HASH F(x)s
// Passwords need to/must be HASHED first before storing them in DBs 🛡️

// Ex. of hashing algorithms: SHA-256, SHA-512 etc.

// salting: Adding some value with the hashed-password for extra security. (protection against dictionary attacks, rainbow-table attacks)

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

	fmt.Println("--------------- SALTING🛡️ ---------------")
	
	salt,err:=GenSalt()
	if err!=nil{
			fmt.Println("🔴ERROR generating salt:",err)
			return
		}
	
	// hash the password with salt
	hash:=HashPassword(password,salt)
	
	// store the salt and password in DB (mocking that here for now)
	saltStr:=base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt:",saltStr)
	fmt.Println("Hash:",hash)



}

// generate salt
	func GenSalt()([]byte, error)  {
		salt:=make([]byte,16) // 16 bytes
		_,err:=io.ReadFull(rand.Reader, salt)
		// We already know the no. of bytes
		if err!=nil{
			fmt.Println("🔴ERROR generating salt:",err)
			return nil,err
		}
		return salt,nil
	}

// function to hash the password (using the salt)
// Storing the salt in an ENCODED way: Added security
	func HashPassword(password string, salt []byte)string{
		saltedPassword:= append(salt, []byte(password)...)
		hash:=sha256.Sum256(saltedPassword)
		return base64.StdEncoding.EncodeToString(hash[:])
	}