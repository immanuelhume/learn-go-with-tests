package main

import "fmt"

func main() {
	fmt.Println(Hello("Doge", ""))
}

func Hello(name string, lang string) string {
	if name == "" {
		name = "Cruel World"
	}
	return greetingPrefix(lang) + name
}

const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const englishHelloPrefix = "Hello, "

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
