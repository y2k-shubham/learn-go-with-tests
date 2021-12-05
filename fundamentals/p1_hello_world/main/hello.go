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

func Greeting(lang int) string {
	if _, ok := langGreetingMap[lang]; !ok {
		lang = English
	}
	greeting := langGreetingMap[lang]
	return greeting
}

func World(lang int) string {
	if _, ok := langGreetingMap[lang]; !ok {
		lang = English
	}
	world := langWorldMap[lang]
	return world
}

func Hello(name string, lang int) string {
	greeting := Greeting(lang)
	if name == "" {
		name = World(lang)
	}

	return fmt.Sprintf("%s %s!", greeting, name)
}

func main() {
	fmt.Println(Hello("Puchi", 0))
	fmt.Println(Hello("Monchi", Spanish))
	fmt.Println(Hello("", French))
	fmt.Println(Hello("", 0))
}
