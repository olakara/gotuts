package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd *cobra.Command

func init() {
	helloCmd = &cobra.Command{
		Use:   "hello",
		Short: "Prints hello world",
		Run:   sayHello,
	}
	helloCmd.Flags().StringP("name", "n", "World", "Name to say hello to")
	helloCmd.MarkFlagRequired("name")
	helloCmd.Flags().StringP("language", "l", "en", "Which langauge to say hello in")
}

func main() {
	helloCmd.Execute()
}

func sayHello(cmd *cobra.Command, args []string) {

	name, _ := cmd.Flags().GetString("name")
	language, _ := cmd.Flags().GetString("language")
	switch language {
	case "en":
		sayHelloInEnglish(name)
	case "es":
		sayHelloInSpanish(name)
	case "fr":
		sayHelloInFrench(name)
	case "de":
		sayHelloInGerman(name)
	default:
		sayHelloInEnglish(name)
	}
}

func sayHelloInEnglish(name string) {
	fmt.Printf("Hello %s\n", name)
}

func sayHelloInSpanish(name string) {
	fmt.Printf("Hola %s\n", name)
}

func sayHelloInFrench(name string) {
	fmt.Printf("Bonjour %s\n", name)
}

func sayHelloInGerman(name string) {
	fmt.Printf("Hallo %s\n", name)
}
