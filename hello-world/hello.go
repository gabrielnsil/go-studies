package main

import "fmt"

const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Salut, "

func Hello(name string, language string) string {
    if name == "" {
	    name = "World"
    }

    prefix := greetingPrefix(language)

    return prefix + name
}

func greetingPrefix(language string) (prefix string) {

    switch language {
    default:
	    prefix = englishHelloPrefix
    case spanish:
	    prefix = spanishHelloPrefix
    case french:
	    prefix = frenchHelloPrefix
    }

    return prefix
}

func main() {
	fmt.Println(Hello("Gabriel", ""))
}
