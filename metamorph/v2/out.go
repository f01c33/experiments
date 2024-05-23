package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	cr "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/cosmos72/gomacro/fast"
)

func logandForget(e error) {
	fmt.Println("error:\t", e)
}

func NewEncryptionKey() []byte {
	key := [32]byte{}
	_, err := io.ReadFull(cr.Reader, key[:])
	if err != nil {
		panic(err)
	}
	return key[:]
}
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Encrypt encrypts data using 256-bit AES-GCM.  This both hides the content of
// the data and provides a check that it hasn't been altered. Output takes the
// form nonce|ciphertext|tag where '|' indicates concatenation.
func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}
func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

var fl = flag.String("f", "", "Run this go file")
var dbg = flag.Bool("g", false, "debug file while running")
var output = flag.String("o", "crypt_out"+fmt.Sprint(time.Now().Local().Unix())+".go", "Select name of the output file")

func init() {
	flag.Parse()
}

// func tob64(v []byte) string {
// 	return base64.RawURLEncoding.EncodeToString(v)
// }
// func enc(text, newkey, key []byte, sep string) []byte {
// 	text = append(text, []byte(sep)...)
// 	e, err := Encrypt(append(text, []byte(tob64(newkey))...), key)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return e
// }

func finalSplit(d string, sep string) []string {
	return strings.Split(d, sep)
}

func fromb64(v string) []byte {
	tmp, err1 := base64.RawURLEncoding.Strict().DecodeString(v)
	logandForget(err1)
	return tmp
}

// func fromb64(v string) []byte {
// 	tmp, err := base64.RawURLEncoding.Strict().DecodeString(v)
// 	logandForget(err)
// 	return tmp
// }
// func dec(text []byte, key []byte) ([]byte, []byte) {
// 	d, err := Decrypt(text, key)
// 	if err != nil {
// 		if len(d) == 0 {
// 			return []byte("finish(); // owned"), NewEncryptionKey()
// 		} else {
// 			// fmt.Println(err)
// 		}
// 	}
// 	return d, []byte(strings.Split(string(d), " // ")[1])
// }
func finalJoin(enc []string, key []byte, sep string) string {
	// fmt.Println("x0r4")
	enc = append(enc, tob64(key))
	// fmt.Println(string(bytes.Join(enc, []byte(sep))))
	return strings.Join(enc, sep)
}
func tob64(v []byte) string {
	return base64.RawURLEncoding.EncodeToString(v)
}
func enc(text string) string {
	// text = append(text, []byte(sep)...)
	// text = append(text, []byte(tob64(newkey))...)
	// e := encrypt(text, string(key))
	// e, err := Encrypt(append(text, []byte(tob64(newkey))...), key)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// return tob64(e)
	return tob64([]byte(text))
}

// func fromb64(v string) []byte {
// 	tmp, err2 := base64.RawURLEncoding.Strict().DecodeString(v)
// 	logandForget(err2)
// 	return tmp
// }
func dec(text string) string {
	return string(fromb64(text))
	// d := fromb64(text)
	// d2 := string(d)
	// // d, err := Decrypt(fromb64(text), key)
	// // if err != nil {
	// // 	if len(d) == 0 {
	// // 		return "", ""
	// // 	} else {
	// // 		fmt.Println(err)
	// // 	}
	// // }
	// // fmt.Println(strings.Split(string(d), " // ")[1])
	// if len(strings.Split(d2, " // ")) > 1 {
	// 	return d2, strings.Split(d2, " // ")[1][:len(strings.Split(d2, " // ")[1])-1]
	// }
	// return d2, ""
}
func dexec(code string) {
	encoded2 := strings.Split(code, "|")
	// k2 := fromb64(encoded2[len(encoded2)-1])
	// encoded2 = encoded2
	fmt.Println("exec")
	interp := fast.New()
	for _, v := range encoded2 {
		d := dec(v)
		// fmt.Println(d)
		fmt.Println(interp.Eval(d))
		// k2 = fromb64(k)
	}
}

// const sepFinal = `#@!!@#`

func main() {
	// key := NewEncryptionKey()

	dexec(`Cg|aW1wb3J0ICJmbXQiCg|Cg|ZnVuYyBzdW0gKGEsYiBpbnQpIGludCB7CglyZXR1cm4gYStiOwp9Cg|Zm10LlByaW50bG4oc3VtKDEsMikpCg|Zm10LlByaW50bG4oc3VtKDEsMikpCg|Zm10LlByaW50bG4oc3VtKDEsMikpCg|Zm10LlByaW50bG4oc3VtKDEsMikpCg|Zm10LlByaW50bG4oc3VtKDEsMikpCg|Cg|ZnVuYyBzMigpIHN0cmluZyB7CglyZXR1cm4gIm93bmVkIjsKfQo|Zm10LlByaW50bG4oczIoKSkK`)
}
