package main

import "fmt"

const englishHelloPrefix string = "Hello"
const world string = "World"

func Hello(name string) string {
	if name == "" {
		name = world
	}
	return fmt.Sprintf("%s %s!", englishHelloPrefix, name)
}

func main() {
	fmt.Println(Hello("Puchi"))
	fmt.Println(Hello(""))
}
