package main

import (
	"flag"
	"fmt"

	"github.com/sina-devel/gwc"
)

var (
	lc    bool
	wc    bool
	files []string
)

func init() {
	app := flag.NewFlagSet("gwc", flag.ExitOnError)
	app.BoolVar(&lc, "l", false, "print the newline counts")
	app.BoolVar(&wc, "w", false, "print the word counts")
	app.Parse(flag.Args())
	if !lc && !wc {
		lc, wc = true, true
	}
	files = app.Args()
	if len(files) == 0 {
		app.Usage()
	}
}

func main() {
	gwc := gwc.New(gwc.Config{
		Filenames: files,
		WC:        wc,
		LC:        lc,
	})
	gwc.Compute()
	fmt.Println(gwc)
}
