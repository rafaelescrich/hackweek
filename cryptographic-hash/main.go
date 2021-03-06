package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

// Hashs

func main() {

	// message to be hashed
	msg := "Hackweek é o melhor evento da PANDEMIA!"

	// MD5 - https://en.wikipedia.org/wiki/MD5
	md5 := md5.Sum([]byte(msg))

	// SHA1 - https://en.wikipedia.org/wiki/SHA-1
	sha1 := sha1.Sum([]byte(msg))

	// SHA2 - https://en.wikipedia.org/wiki/SHA-2 (most used nowadays)
	sha2 := sha256.Sum256([]byte(msg))

	// HMAC - https://pt.wikipedia.org/wiki/HMAC
	key := []byte("HACKWEEK")
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(msg))

	// Print all hashes
	fmt.Printf("MD5: %x\n", md5)
	fmt.Printf("SHA1: %x\n", sha1)
	fmt.Printf("SHA2-256: %x\n", sha2)
	fmt.Printf("HMAC: %x\n", mac.Sum(nil))

}
