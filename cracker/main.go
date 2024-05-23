package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	// "fmt"
	"hash"
	"log"

	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

var (
	hasher *string
	h *string
	wl *string
)

func init(){
	hasher = flag.String("alg", "sha1", "Choose hash kind")
	h = flag.String("hash","74686567616d65da39a3ee5e6b4b0d3255bfef95601890afd80709", "Hash to crack")
	wl = flag.String("word","wordlist", "wordlist")
	flag.Parse()
}

func main(){
	var hsr hash.Hash
	_ = hsr
	wrds, err := ioutil.ReadFile(*wl)
	if err != nil {
		panic(err)
	}
	wrd := make(chan string,200)
	out := make(chan string,1)
	for i := 0; i < 20; i++{
		go func(wrd chan string,out chan string){
			var hs hash.Hash
			switch(*hasher){
			case "sha1":
				hs = sha1.New()
			case "md5":
				hs = md5.New()
			case "sha512":
				hs = sha512.New()
			case "ripemd160":
				hs = ripemd160.New()
			case "sha224":
				hs = sha3.New224()
			case "sha256":
				hs = sha3.New256()
			case "sha384":
				hs = sha3.New384()
			case "sha512_224":
				hs = sha512.New512_224()
			case "sha512_256":
				hs = sha512.New512_256()
			default:
				log.Fatal("You need to specify a valid hash")
			}
			for{
				select{
				case w,ok := <- wrd:
					hs.Reset()
					// fmt.Printf("%x\n",h.Sum([]byte(w)))
					if fmt.Sprintf("%x",hs.Sum([]byte(w))) == *h{
						// fmt.Println(w)
						out <- w
						return
					}
					if !ok {
						return
					}
				default:
					if len(out) != 0 {
						return
					}
				}
			}
		}(wrd,out)
	}
	for _,w := range strings.Split(string(wrds),"\n"){
		wrd <- w
	}
	// close(wrd)
	for{
		if len(wrd) != 0{
			<-time.After(1*time.Second)
		}else{
			break
		}
	}
	if len(out) == 0 {
		fmt.Println("Found nothing")
	}else{
		fmt.Println(<-out)
	}
}