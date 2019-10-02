package main

import (
	"fmt"
	"os"
)

func main() {
	var locale string
	var translations = make(map[string]string)
	translations["en"] = "hello"
	translations["de"] = "gtten tag"
	translations["fr"] = "bonjour"
	//var languages = [4]string{"en", "es", "de", "fr"}
	//locale = languages[1]

	if len(os.Args) == 1 {
		fmt.Println("enter language code")
		fmt.Scanf("%s", &locale)
	} else {
		locale = os.Args[1]
	}
	output := translations[locale]
	if output == "" {
		output = "Yo"
	}

	fmt.Printf(output, "Go")
}
