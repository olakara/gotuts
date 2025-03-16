package main

import (
	"flag"
	"fmt"
)

type language = string

var name = flag.String("n", "World", "A name to say hello to.")
var userLanguage language

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
