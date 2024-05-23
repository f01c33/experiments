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
	"io/ioutil" //
	"math"
	"os"
	"strings"

	"github.com/starlight-go/starlight"
)

func logandForget(e error) {
	// fmt.Println("error:\t", e)
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

var fl = flag.String("f", "", "Run this python-starlink file")
var dbg = flag.Bool("g", false, "debug file while running")

// var output = flag.String("o", "crypt_out"+fmt.Sprint(time.Now().Local().Unix())+".go", "Select name of the output file")

func init() {
	flag.Parse()
}

func main() {
	var g map[string]interface{}
	var err error
	// var c []byte
	code := `code="%v"
def dexec(code):
	# print(fromb64(toStr(code)))
	print("DEXEC")
	encoded = finalSplit(fromb64(code), sepFinal)
	k = encoded[-1]
	for i in range(len(encoded)):
		d,nk = dec(encoded[i],k)
		#print(toStr(d))
		eval(d)
		k = fromb64(nk)
dexec(code)
	`
	if *fl != "" {
		c, err := ioutil.ReadFile(*fl)
		if err != nil {
			panic(err)
		}
		code = string(c)
	}
	skeleton := NewEncryptionKey()
	stack := []map[string]interface{}{}
	globals := map[string]interface{}{
		"":      skeleton,
		"sepFinal": ,
		"stack":    stack,
		// 		"inp": `print("wow")
		// print("mom")
		// print("wom")
		// print("mow")`,
		// "wait": func(t string) {
		// 	// fmt.Println("Backtogo")
		// 	tmp, err := time.ParseDuration(t)
		// 	logandForget(err)
		// 	time.Sleep(tmp)
		// },
		"evalStar": func(s string) (map[string]interface{}, error) {
			// fmt.Println("Backtogo")
			return starlight.Eval(interface{}([]byte(s)), g, nil)
		},
		"newkey": func() []byte {
			return NewEncryptionKey()
		},
		"toByteArr": func(s string) []byte {
			return []byte(s)
		},
		// "toByte": func(ints []string) []byte {
		// 	out := []byte{}
		// 	for _, i := range ints {
		// 		val, err := strconv.Atoi(i)
		// 		if err != nil {
		// 			panic(err)
		// 		}
		// 		out = append(out, byte(val))
		// 	}
		// 	return out
		// },
		"toStr": func(b []byte) string {
			return string(b)
		},
		// "enc": func(text []byte, key []byte) []byte {
		// 	e, err := Encrypt(text, key)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	}
		// 	return e
		// },
		"dec": func(text []byte, key []byte) ([]byte, []byte) {
			d, err := Decrypt(text, key)
			if err != nil {
				if len(d) == 0 {
					return []byte("finish() # "), NewEncryptionKey()
				} else {
					fmt.Println(d, err)
				}
			}
			// rg := regexp.MustCompile(`[0-9a-zA-Z_]+`)
			// all := rg.FindAllString(string(d), -1)
			return d, []byte(strings.Split(string(d), " # ")[1])
		},
		// "tob64": func(v []byte) string {
		// 	return base64.RawURLEncoding.EncodeToString(v)
		// },
		"fromb64": func(v string) []byte {
			tmp, err := base64.RawURLEncoding.Strict().DecodeString(v)
			logandForget(err)
			return tmp
		},
		"finish": func() {
			fmt.Println("x0r4")
			return
		},
		// "finalJoin": func(enc [][]byte, key []byte, sep string) []byte {
		// 	// fmt.Println("x0r4")
		// 	enc = append(enc, key)
		// 	fmt.Println(string(bytes.Join(enc, []byte(sep))))
		// 	return bytes.Join(enc, []byte(sep))
		// },
		"finalSplit": func(d []byte, sep string) [][]byte {
			return bytes.Split(d, []byte(sep))
		},
		// "push":
		"exit": func(code int) {
			os.Exit(code)
		},
		"sqrt": math.Sqrt,
		// "byteArrArr": func() [][]byte { return [][]byte{} },
	}
	// glob = func() map[string]interface{} { return globals }
	evalG := func(s []byte) (map[string]interface{}, error) {
		// d :=
		// err := json.Unmarshall(s,d)
		// logandForget(err)
		globals, err := starlight.Eval(s, globals, nil)
		logandForget(err)
		stack = append(stack, globals)
		return globals, err
		// js, err := json.Marshal(m)
		// logandForget(err)
		// return string(js)
	}
	// push := func(v string) {
	// globals["stack"] = append(globals["stack"].([]string), v)
	// stack = append(stack, v)
	// }
	globals["eval"] = evalG
	globals["push"] = func(s interface{}) {
		fmt.Println(s)
	}
	globals, err = starlight.Eval([]byte(code), globals, nil)
	fmt.Println("END")
	fmt.Println(g)
	fmt.Println(err)
	fmt.Println(stack)
}
