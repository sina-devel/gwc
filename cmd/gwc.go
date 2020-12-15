package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sina-devel/gwc"
)

var config = gwc.Config{}

func init() {
	app := flag.NewFlagSet("gwc", flag.ExitOnError)
	app.BoolVar(&(config.LC), "l", false, "print the newline counts")
	app.BoolVar(&(config.WC), "w", false, "print the word counts")
	app.Parse(os.Args[1:])
	if !config.LC && !config.WC {
		config.LC, config.WC = true, true
	}
	config.Filenames = app.Args()
	if len(config.Filenames) == 0 {
		app.Usage()
	}
}

func main() {
	gwc := gwc.New(&config)
	gwc.Compute()
	fmt.Println(gwc)
}
