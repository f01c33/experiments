package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

var site = flag.String("w", "https://reddit.com/", "site to grep")
var regx = flag.String("r", "r/", "regex to use at the urls")
var dbg = flag.Bool("g", false, "debug file while running")

// var output = flag.String("o", "crypt_out"+fmt.Sprint(time.Now().Local().Unix())+".go", "Select name of the output file")

func init() {
	flag.Parse()
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains(*site),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		data := ""
		// if os.Args[2] == "-a" {
		data = e.Attr("href")
		// } else {
		// 	fmt.Print("not implemented: ", os.Args[2])
		// }
		// Print link
		if match, err := regexp.Match(strings.Join(os.Args[3:], " "), []byte(data)); match && err != nil {
			fmt.Println(data)
		}
		// fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit(os.Args[1])
}
