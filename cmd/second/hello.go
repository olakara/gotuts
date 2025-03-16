package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
)

type language = string

var name = flag.String("n", "World", "A name to say hello to.")
var userLanguage language

var validLanguages = []string{
	"en",
	"sp",
	"fr",
	"de",
}

const (
	Spanish language = "sp"
	English language = "en"
	French  language = "fr"
	German  language = "de"
)

func init() {

	flag.StringVar(&userLanguage, "l", "en", "Language to use are en, sp, fr, de")
	flag.StringVar(&userLanguage, "lang", "en", "Language to use are en, sp, fr, de")

	flag.Parse()
	if !slices.Contains(validLanguages, userLanguage) {
		fmt.Println("Invalid language")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	switch userLanguage {
	case Spanish:
		fmt.Printf("Hola %s!\n", *name)
	case English:
		fmt.Printf("Hello %s!\n", *name)
	case French:
		fmt.Printf("Bonjour %s!\n", *name)
	case German:
		fmt.Printf("Hallo %s!\n", *name)
	}
}
