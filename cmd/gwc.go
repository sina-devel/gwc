package main

import (
	"fmt"

	"github.com/sina-devel/gwc"
)

func main() {
	gwc := gwc.New(gwc.Config{
		Filenames: []string{"./", "gwc.go"},
		WC:        true,
		LC:        true,
	})
	fmt.Println(gwc)
}
