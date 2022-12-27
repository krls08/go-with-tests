package main

import "fmt"

// This is not possible to test, we need to move the domain logic from main function to be able to test it

//	func main() {
//		fmt.Println("Hello world")
//	}
const (
	SPANISH         = "Spanish"
	FRENCH          = "French"
	EN_HELLO_PREFIX = "Hello"
	SP_HELLO_PREFIX = "Hola"
	FR_HELLO_PREFIX = "Bonjour"
)

func Hello() string {
	return "Hello, world!"
}

func HelloC(s string) string {
	if s == "" {
		s = "World"
	}
	return fmt.Sprintf("Hello, %s!", s)
}

func HelloLang(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(lang) + ", " + name + "!"
}

func getPrefix(lang string) string {
	var helloPrefix string
	switch lang {
	case SPANISH:
		helloPrefix = SP_HELLO_PREFIX
	case FRENCH:
		helloPrefix = FR_HELLO_PREFIX
	default:
		helloPrefix = EN_HELLO_PREFIX

	}
	return helloPrefix
}

func main() {
	fmt.Println(Hello())
	fmt.Println(HelloC("Dominik"))
}
