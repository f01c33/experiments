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
	code := `code="THwghhruGsbvS3jGZWLKK-jp0Nh_8kMWfJPFbpox1rTDkdHdy6xzuDC8iEx5PfPiuOj9s0vEuMhg5aRplSURPP6CRJgWTX6rjs0JpJ--mmzYlZFODF8LDC96ihZAvtu7jC1L5ou-RFeiC23_KC5BDpLtTrixA0Inr0r4-mEzDnCXESWEIsOBUy8F_LQf5ebJDYvi235tLCX8JsO4uFKToA1J5Fot2mgi1YcB7mBv3SPUvA-bUxJGKdKHFHXFOqMK0deIHY3eq_QNYCXB1yo4tI32tRpSM7U4puo-sBChr67rVBGQpvL9ErmmeOLzhmH2-__SPGyVC2qb97eMo6ffh6IjhZzV2fEiEXKcZ8tOSuprlXjJpRCNdp3nzpKA4csVnrcdwA7lewRw0D6mjwf9P42_a0I2Q1D3VgaznxkAiz0YN9KR_KMPIp70GX9s5acLGUd95RgjL8S28GfFk04VmahkzzgB_ZQjN5jJtFs8_tgwnhXyRD-USlNrXbY2KiPE6SEKR-allwG0R7Y1Z2k2Ble3u6TXGsu9-zdPPy8pTdX-45fwv2RAnBF3-1bzcGv_qs8KPggqFnjTTj9RkgNQ1YQiDIsEtVDt3j_8iVUlgW5w0YlLrDWFZ0A6VpLLh2pa0E_njnD3v4uadu1ouEpnyV1PhLugZj5wtQGJxRtKDW3XuagdBqZUhuIWiTyUdXWIZCYko0yFVPpTB_NLNMjzK5ZwwfSBRl4ahU9RMIzqwSFYuNlSbSmMX7bah5HPeeKBxy5OppMsuXO0gf-ktJ5gMH1zP_P5_wSvsrpqzJLaiJqHyDDH8gZ6wUfjezeZGimB7kxWzqkZoHpG1UTNiTfUizSUKO6Cyz2iLlO6tJpt1cIA7WRmvjoTr_Xo3fnbCd0t-7GwDZrqdIdvFyedeKp-Zf2L2_oO33KkLCyUWoxIhgDicr5DKHHec2iEcO3gKmO30yiGF5LMzSukE1JKIeG7uuKwhu1ZHWnga16Px4Ir8pJQtEpRjYRrBXuMFKagdedl0JLKTB6MOcwh3nlfhQKewq6fMq3zmUzeGoEagvYx8BsKD6HEUeRlAxk3oklfnvciRZDD2Hl-ivII7SgrhG_35nJPkWPaBBASWVK9Yd4Hx7q5eyvfcnci9_Qvpu7yebOW5TN1_zymtDcdySOiAa995cqhUzxWNmhrjdRvhFZbqO2gWZ9V4PyWGz5N4fbmiVwpD6Y2NyvE4Ju7eq5nO0bcmMQXF3Y00bZqiIFNJE17lXTljvsrWfvRItFjvC33kaDsSjBCKIABAAEu0H4PKFKoqLbNqImLjDqP74wkx57hBYvF9zFS6zp6dFIoHoKbm5vh8qv-gf5yRIwZHaVadG4k3O4cS4Vx5UL1vgOhdVXTM1n1-ulqQX6goCtvlCrr72PHkZPLdVAuGXsh7F4hOvT39fdDTcFcO7-Ul8t6GLVclI-DSi52uTaVPlvZ70adB0M2lrSFrdoxj1AbUyaKEe5Gpi_iBoWBH-GU1vMpOL6C-M8nXXm_iuXH3mkSEtrhLn-ZKqRMvs5gGi6S69awk-jwipdWqDy_tohA5DCAdnRR9Qyo5ZyXw8Johjr_vSjEPM5k2uTn6wqVtY4wXogHeYVcq7QUVXOmxJX2rHa162Pso_owyo9lZKGR_3LEvSyJfvY5389Bd4eg2Q6SCiTLl0WbK3cBPx2wSFg4_avhLoQODWQ7xsMcQC9-mUSL1lTjYgb3mzPdk6AjQCEhQCMD3KnZarOERBAIWGaaIWaaT-S4LrkzNawXMmCpFFREiA"
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
		"key":      skeleton,
		"sepFinal": "#@!!@#",
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
