package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Create(fmt.Sprint(time.Now().Unix()) + ".db")
	if err != nil {
		panic(err)
	}
	n, err := io.Copy(file, r.Body)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(fmt.Sprintf("%d bytes are recieved.\n", n)))
	// ioutil.WriteFile(,)
}

func main() {
	// tip: use bine to torify this
	http.HandleFunc("/", uploadHandler)
	http.ListenAndServe(":5050", nil)
}
