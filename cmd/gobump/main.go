package main

import (
	"flag"
	"fmt"
	"log"

	. "github.com/pierdipi/gobump"
	skip "github.com/pierdipi/gobump"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	target := flag.String("target", "", "Go target version")
	flag.Parse()

	if target == nil || *target == "" {
		return fmt.Errorf("invalid target")
	}

	return ApplyTarget(*target, ".", ReadWriter, skip.NonGoMod, skip.Vendor)
}
