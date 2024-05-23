package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	cr "crypto/rand"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"time"
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

// Encrypt encrypts data using 256-bit AES-GCM.  This both hides the content of
// the data and provides a check that it hasn't been altered. Output takes the
// form nonce|ciphertext|tag where '|' indicates concatenation.
func Encrypt(plaintext []byte, key []byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(cr.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// Decrypt decrypts data using 256-bit AES-GCM.  This both hides the content of
// the data and provides a check that it hasn't been altered. Expects input
// form nonce|ciphertext|tag where '|' indicates concatenation.
func Decrypt(ciphertext []byte, key []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
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

var fl = flag.String("f", "engine.py", "Run this python-starlink file")
var dbg = flag.Bool("g", false, "debug file while running")
var output = flag.String("o", "crypt_out"+fmt.Sprint(time.Now().Local().Unix())+".go", "Select name of the output file")

func init() {
	flag.Parse()
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

func fromb64(v string) []byte {
	tmp, err := base64.RawURLEncoding.Strict().DecodeString(v)
	logandForget(err)
	return tmp
}
func dec(text []byte, key []byte) ([]byte, []byte) {
	d, err := Decrypt(text, key)
	if err != nil {
		if len(d) == 0 {
			return []byte("finish() # "), NewEncryptionKey()
		} else {
			// fmt.Println(err)
		}
	}
	return d, []byte(strings.Split(string(d), " # ")[1])
}
func finalJoin(enc [][]byte, key []byte, sep string) []byte {
	// fmt.Println("x0r4")
	enc = append(enc, key)
	// fmt.Println(string(bytes.Join(enc, []byte(sep))))
	return bytes.Join(enc, []byte(sep))
}
func main() {
	inp := `for i in range(10):
	print(i)
`
	if *fl != "" {
		i, err := ioutil.ReadFile(*fl)
		if err != nil {
			panic(err)
		}
		inp = string(i)
	}
	sepFinal := "#@!!@#"
	encoded := [][]byte{}
	key := NewEncryptionKey()
	lines := strings.Split(inp, "\r\n")
	k := key
	for _, l := range lines {
		if l == "" {
			continue
		}
		nk := NewEncryptionKey()
		// if *dbg {
		// 	// fmt.Println("line: ", l, l+" # "+string(tob64(nk)))
		// }
		e := enc([]byte(l), nk, k, " # ")
		// fmt.Println(string(e))
		encoded = append(encoded, e)
		g, nk1 := dec(e, k)
		if *dbg {
			fmt.Println("E > ", string(g))
			if string(nk1) != string(tob64(nk)) {
				fmt.Println(string(g))
				fmt.Println("nk1: \t", string(nk))
				fmt.Println("nk: \t", string(tob64(nk)))
			}
		}
		k = nk
	}
	final := tob64(finalJoin(encoded, key, sepFinal))
	if output != nil && *output != "" {
		b, err := ioutil.ReadFile("runner.go")
		if err != nil {
			panic(err)
		}
		ioutil.WriteFile(*output, []byte(fmt.Sprintf(string(b), final)), 0777)
	} else {
		fmt.Printf(final)
	}
}
