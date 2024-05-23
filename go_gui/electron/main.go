package main

import (
	"log"
	"os"

	"github.com/asticode/go-astilectron"
)

func main() {
	// fmt.Println("hello world")
	// Initialize astilectron
	var a, _ = astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
		AppName: "testname",
		// AppIconDefaultPath: "<your .png icon>",  // If path is relative, it must be relative to the data directory
		// AppIconDarwinPath:  "<your .icns icon>", // Same here
		// BaseDirectoryPath:  "<where you want the provisioner to install the dependencies>",
		// VersionAstilectron: "<version of Astilectron to utilize such as `0.33.0`>",
		// VersionElectron:    "<version of Electron to utilize such as `4.0.1` | `6.1.2`>",
	})
	defer a.Close()

	// Start astilectron
	a.Start()

	// Blocking pattern
	a.Wait()

}
