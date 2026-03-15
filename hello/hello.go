package hello

const (
	spanish = "Spanish"
        french = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Salut, "
)

// The function takes two arguments: name and language and returns a greeting in the desired language.
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

