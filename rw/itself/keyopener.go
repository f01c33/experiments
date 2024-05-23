package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"log"
	// _ "github.com/kevinburke/go-bindata"
)

// BytesToPublicKey bytes to public key
func BytesToPublicKey(pub []byte) *rsa.PublicKey {
	block, _ := pem.Decode(pub)
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
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		log.Fatal(err)
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		log.Fatal("not ok")
	}
	return key
}

// EncryptWithPublicKey encrypts data with public key
func EncryptWithPublicKey(msg []byte, pub *rsa.PublicKey) []byte {
	hash := sha512.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
	if err != nil {
		log.Fatal(err)
	}
	return ciphertext
}

func EncryptWithKey(key, text string) string {
	data, err := Asset(key)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(data))
	pub := BytesToPublicKey(data)
	return base64.StdEncoding.EncodeToString(EncryptWithPublicKey([]byte(text), pub))
}

// func main() {
// 	enc := EncryptWithKey("keys/public.key", os.Args[1])
// 	fmt.Print(enc)
// 	// statikFS, err := fs.New()
// }
