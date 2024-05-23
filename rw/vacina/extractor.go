package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	// _ "github.com/kevinburke/go-bindata"
)

func decryptWithKey(key, cyphertext string) string {
	data, err := Asset(key)
	if err != nil {
		panic(err)
	}
	priv := BytesToPrivateKey(data)
	ct, err := base64.StdEncoding.DecodeString(cyphertext)
	if err != nil {
		panic(err)
	}
	out := DecryptWithPrivateKey(ct, priv)
	return string(out)
}

// BytesToPrivateKey bytes to private key
func BytesToPrivateKey(priv []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(priv)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		log.Fatal(err)
	}
	return key
}

// DecryptWithPrivateKey decrypts data with private key
func DecryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) []byte {
	hash := sha512.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)
	if err != nil {
		log.Fatal(err)
	}
	return plaintext
}

func main() {
	dec := decryptWithKey("keys/private.key", os.Args[1])
	fmt.Println(dec)
}
