package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	gfe "github.com/kiltum/go-file-encrypt"
	"github.com/skratchdot/open-golang/open"
)

func decryptFiles(files []string, key string) {
	wg := sync.WaitGroup{}
	for _, f := range files {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := gfe.DecryptFile(f, key)
			if err != nil {
				fmt.Printf("Error %v", err)
				return
			}
		}()
	}
	wg.Wait()
}

func main() {
	root := []string{"../teste"}
	files := []string{}
	dirs := []string{}
	for _, r := range root {
		err := filepath.Walk(r, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() && path[len(path)-len(".encrypt"):] == ".encrypt" {
				files = append(files, path)
			} else {
				dirs = append(dirs, path)
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("wowie")
		if r.Method == "POST" {
			// fmt.Println("wowie2")
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			fmt.Println(r.FormValue("keyname"))
			decryptFiles(files, r.FormValue("keyname"))
			http.ServeFile(w, r, "index.html")
		} else if r.Method == "GET" {
			http.ServeFile(w, r, "index.html")
		}
	})
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("wowie")
		if r.Method == "POST" {
			// fmt.Println("wowie2")
			wg := sync.WaitGroup{}
			for _, f := range files {
				wg.Add(1)
				if strings.Index(f, ".encrypt") != -1 {
					go func() {
						defer wg.Done()
						os.Remove(f)
						// if err != nil {
						// 	fmt.Printf("Erro %v", err)
						// 	return
						// }
					}()
				}
			}
			wg.Wait()
			http.ServeFile(w, r, "index.html")
		} else if r.Method == "GET" {
			http.ServeFile(w, r, "index.html")
		}
	})
	open.Run("http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
