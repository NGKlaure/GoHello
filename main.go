package main

import "fmt"

func main() {
	var locale, greeting string
	var languages = [4]string{"en", "es", "de", "fr"}
	locale = languages[1]
/*
	if locale == "en" {
		greeting = "Hello"
	} else if locale == "es" {
		greeting = "Hola"
	} else if locale == "de" {
		greeting = "Guten Tag"
	} else {
		greeting = "Yo"
	}
*/

switch locale {
case "en":
	greeting = "Hello"
case "es":
	greeting = "Hola"
case "de":
	greeting = "Guten Tag"
case "fr":
	greeting = "Bonjour"
default:
	greeting = "Yo"
}

  fmt.Println(greeting + ",Go!/n" )
  fmt.Printf(greeting + ", Go!\n")
}