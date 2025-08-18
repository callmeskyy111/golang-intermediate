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

// ⚠️ Use crypto/rand rather than math/rand, for security purposes 🔴 - SEEDING not require with crypto, as it uses high-level of entropy

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
	fmt.Println("Original Salt:",string(salt))
	fmt.Printf("Original Salt in HEX: %x\n",salt)
	if err!=nil{
			fmt.Println("🔴ERROR generating salt:",err)
			return
		}
	
	// hash the password with salt
	signupHash:=HashPassword(password,salt)
	// signupHash:=HashPassword("password124",salt) // 🔴 ERROR
	
	// store the salt and password in DB (mocking that here for now..)
	saltStr:=base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt:",saltStr)
	fmt.Println("Hash:",signupHash)

	hashOriginalPassword:= sha256.Sum256([]byte(password))
	fmt.Println("Hash of just the password-string without salt:",base64.StdEncoding.EncodeToString(hashOriginalPassword[:]))


	// verify after logging in
	// retrieve the saltStr and decode it
	decodedSalt,err:= base64.StdEncoding.DecodeString(saltStr) // simulating a DB storage
	if err!=nil{
			fmt.Println("🔴ERROR decoding:",err)
			return
		}
	loginHash:= HashPassword(password, decodedSalt)

	// Compare the stored signupHash with the loginHash	
	if signupHash==loginHash{
		fmt.Println("Password is correct. You are logged in ✅")
		// then send the user to the appropriate end-point (dashboard, signIn etc.)
	}else{
		fmt.Println("🔴 Login failed.. Please check user credentials! ⚠️")
		// real-world generic message, not letting the user know whether pasword, email or username is incorrect
	}



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
