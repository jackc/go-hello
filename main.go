package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	var opts struct {
		Name string `short:"n" long:"name" description:"Name to greet" default:"world"`
	}

	parser := flags.NewParser(&opts, flags.Default)
	parser.Usage = "[options]"
	if _, err := parser.Parse(); err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing command line")
		os.Exit(1)
	}

	fmt.Printf("Hello, %s!\n", opts.Name)
}
