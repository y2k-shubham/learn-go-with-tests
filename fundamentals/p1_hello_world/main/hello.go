package main

import "fmt"

const (
	English = 1
	Spanish = 2
	French  = 3
)

var langGreetingMap = map[int]string{
	English: "Hello",
	Spanish: "Hola",
	French:  "Bonjour",
}
var langWorldMap = map[int]string{
	English: "World",
	Spanish: "Mundo",
	French:  "Monde",
}

func Hello(name string, lang int) string {
	if _, ok := langGreetingMap[lang]; !ok {
		lang = English
	}
	greeting := langGreetingMap[lang]
	if name == "" {
		name = langWorldMap[lang]
	}

	return fmt.Sprintf("%s %s!", greeting, name)
}

func main() {
	fmt.Println(Hello("Puchi", 0))
	fmt.Println(Hello("Monchi", Spanish))
	fmt.Println(Hello("", Spanish))
	fmt.Println(Hello("", 0))
}
