package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"

	. "github.com/pierdipi/gobump"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	skips := []SkipFile{
		NonGoMod,
		Vendor,
	}

	target := flag.String("target", "", "Go target version")
	exclude := flag.String("exclude-regex", "", "Exclude regex")
	flag.Parse()

	if target == nil || *target == "" {
		return fmt.Errorf("invalid target")
	}
	if exclude != nil && *exclude != "" {
		excludeRegex, err := regexp.Compile(*exclude)
		if err != nil {
			return fmt.Errorf("invalid exclude regex: %w", err)
		}

		skips = append(skips, Exclude(excludeRegex))
	}

	return ApplyTarget(*target, ".", ReadWriter, skips...)
}
