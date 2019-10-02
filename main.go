package main

import (
	"fmt"
	"os"
)

func main() {
	locale := argswitch(os.Args)
	output := translate(locale)
	if output == "" {
		output = "Yo"
	}
	fmt.Println(output, "Go!")
}

func translate(input string) string {
	var translations = make(map[string]string)
	translations["en"] = "Hello"
	translations["de"] = "Guten tag"
	translations["fr"] = "Bonjour"
	return translations[input]
}

func argswitch(input []string) string {
	var locale string
	if len(input) == 1 {
		fmt.Println("Enter language code:")
		fmt.Scanf("%s", &locale)
	} else {
		locale = input[1]
	}
	return locale
}
