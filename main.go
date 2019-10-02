package main

import (
	"fmt"
	"os"
)

var translations map[string]string

func main() {
	translations = make(map[string]string)
	translations["en"] = "Hello"
	translations["de"] = "Guten tag"
	translations["fr"] = "Bonjour"

	locale := argswitch(os.Args)
	output := translate(locale)
	if output == "" {
		output = "Yo"
	}

	fmt.Println(output, "Go!")
}

func translate(input string) string {
	return translations[input]
}

func argswitch(input []string) string {
	var locale string
	if len(input) == 1 {
		help()
	} else {
		locale = input[1]
	}
	return locale
}

func help() {
	fmt.Println("usage: hellgo args1")
	fmt.Println("args: ")

	for index := range translations {
		fmt.Println(index)
	}
	os.Exit(0)
}
