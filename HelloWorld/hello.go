package main

import "fmt"

// const englishHelloPrefix = "Hello, "

// func Hello(name string) string {
// 	if name == "" {
// 		name = "World"
// 	}
// 	return englishHelloPrefix + name
// }

// func Hello(name string, language string) string {
// 	if name == "" {
// 		name = "World"
// 	}
// 	return englishHelloPrefix + name
// }

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("world", "french"))
}