package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

const VERSION = "0.2.0pre"

func main() {
	var opts struct {
		Name    string `short:"n" long:"name" description:"Name to greet" default:"world"`
		Version bool   `long:"version" description:"Display version and exit"`
	}

	parser := flags.NewParser(&opts, flags.Default)
	parser.Usage = "[options]"
	if _, err := parser.Parse(); err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing command line")
		os.Exit(1)
	}

	if opts.Version {
		fmt.Println("go-hello " + VERSION)
		os.Exit(0)
	}

	fmt.Printf("Hello, %s!\n", opts.Name)
}
