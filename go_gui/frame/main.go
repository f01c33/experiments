package main

import (
	"fmt"
	"io/ioutil"

	"github.com/fsnotify/fsnotify"
	"github.com/night-codes/frame"
)

func main() {
	app := frame.MakeApp("suckless-browser") // please, use this row as first in main func
	app.SetIconFromFile("./moon.png")

	window := app.NewWindow("web-repl", 450, 300).
		LoadHTML(`<body style="color:#000;">
					<h1>Hello world</h1>
					<p>Test test test...</p>
				</body>`, "about:blank").
		SetBackgroundColor(255, 255, 255, 1)
	go func() {
		window.Show()
	}()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer watcher.Close()

	// done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("modified file:", event.Name)
				}
				data, err := ioutil.ReadFile(event.Name)
				if err != nil {
					panic(err)
				}
				window.LoadHTML(string(data), "about:blank")
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("index.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	// go func() {
	// 	fileAndHash := func(fname string) (data string, hash string) {
	// 		dat, _ := ioutil.ReadFile(fname)
	// 		data = string(dat)
	// 		hsh := sha256.New224()
	// 		hash = string(hsh.Sum(dat))
	// 		return
	// 	}
	// 	file, hash := fileAndHash("index.html")
	// 	for tick := range time.NewTicker(100 * time.Millisecond).C {
	// 		fmt.Println(tick)
	// 		currFile, currHash := fileAndHash("index.html")
	// 		if currHash != hash {
	// 			file = currFile
	// 			hash = currHash
	// 			window.LoadHTML(file, "about:blank")
	// 		}
	// 	}
	// }()
	// window.SetCenter().
	// 	KeepAbove(true).
	// 	SetSize(900, 600).
	// 	Load("https://html5test.com/")

	// go func() {
	// 	window.Show() // show window asynchronously from another go routine

	// 	// You don't have to worry about high resolution screens,
	// 	// the app will look equally good on all screens.
	// 	fmt.Print("Window size: ")
	// 	fmt.Println(window.GetSize()) // Used DPI-related pixels as in browser
	// 	fmt.Print("Window inner size: ")
	// 	fmt.Println(window.GetInnerSize())
	// }()

	// go func() { // Yes! You can change everything in another threads!
	// 	// 	time.Sleep(time.Second * 5)

	// 	// 	fmt.Print("Screen size: ")
	// 	// 	fmt.Println(window.GetScreenSize()) // Used DPI-related pixels as in browser

	// 	// time.Sleep(time.Second * 15)
	// 	// window.Hide() // will close window and finish application after 15 second
	// }()

	app.WaitAllWindowClose() // lock main to avoid app termination (you can also use your own way)
}
