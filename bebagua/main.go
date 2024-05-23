package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	_ "modernc.org/sqlite"

	"github.com/roylee0704/gron"
	naturaldate "github.com/tj/go-naturaldate"
	"github.com/yanzay/tbot/v2"
)

var db *sql.DB

func init() {
	var err error
	first := false
	if _, err := os.Stat("db.sqlite"); err != nil {
		first = true
	}
	db, err = sql.Open("sqlite", "file:db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	if first {
		createTable(db)
	}
}

func main() {
	defer db.Close()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGHUP)

	defer signal.Stop(signalChan)

	go func() {
		for {
			select {
			case s := <-signalChan:
				switch s {
				case syscall.SIGHUP:
				case os.Interrupt:
					os.Exit(1)
				}
			}
		}
	}()

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() (err error) {
	bot := tbot.New("409634068:AAHeW7kcnPaE6e71VBRe4kAcCAL9OTxTstc")
	c := bot.Client()
	bot.HandleMessage("/new.*", func(m *tbot.Message) {
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		insertQuote(db, m.From.Username, strings.Split(m.Text, "/new")[1][1:])
		c.SendMessage(m.Chat.ID, strings.Split(m.Text, "/new")[1][1:])
	})
	bot.HandleMessage(".*/help.*", func(m *tbot.Message) {
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		c.SendMessage(m.Chat.ID, `create quote: /new [quote]
get quote: /get`)
	})
	bot.HandleMessage("/get.*", func(m *tbot.Message) {
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		creators, quotes := getQuotes(db)
		i := rand.Intn(len(quotes))
		c.SendMessage(m.Chat.ID, fmt.Sprintf("%s - %s", quotes[i], creators[i]))
	})
	bot.HandleMessage("/program.*", func(m *tbot.Message) {
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		txt := strings.Split(m.Text, "/program")
		tmr, err := naturaldate.Parse(txt[1], time.Now(), naturaldate.WithDirection(naturaldate.Future))
		if err != nil {
			c.SendMessage(m.Chat.ID, err.Error())
			return
		}
		startGron(tmr.Sub(time.Now()), func() {
			creators, quotes := getQuotes(db)
			i := rand.Intn(len(quotes))
			c.SendMessage(m.Chat.ID, fmt.Sprintf("%s - %s", quotes[i], creators[i]))
		})
	})
	err = bot.Start()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func createTable(db *sql.DB) {
	createQuoteTableSQL := `CREATE TABLE quote (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"creator" TEXT,
		"quote" TEXT
	  );`

	log.Println("Create Quote table...")
	statement, err := db.Prepare(createQuoteTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Quote table created")
}

func insertQuote(db *sql.DB, creator string, quote string) {
	log.Println("Inserting Quote record ...")
	insertQuoteSQL := `INSERT INTO quote (creator, quote) VALUES (?, ?)`
	statement, err := db.Prepare(insertQuoteSQL)

	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(creator, quote)
	if err != nil {
		log.Fatalln(err)
	}
}

func getQuotes(db *sql.DB) (creators []string, quotes []string) {
	row, err := db.Query("SELECT * FROM quote")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	creators = []string{}
	quotes = []string{}
	for row.Next() {
		var id int
		var creator string
		var quote string
		row.Scan(&id, &creator, &quote)
		creators = append(creators, creator)
		quotes = append(quotes, quote)
		log.Println("Quote: ", creator, " ", quote)
	}
	return
}

var c = gron.New()

func startGron(t time.Duration, cb func()) {
	c.AddFunc(gron.Every(t), cb)
	go c.Start()
}
