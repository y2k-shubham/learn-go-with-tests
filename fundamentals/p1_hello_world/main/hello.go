package main

import "fmt"

const englishHelloPrefix string = "Hello"

func Hello(name string) string {
	return fmt.Sprintf("%s %s!", englishHelloPrefix, name)
}

func main() {
	fmt.Println(Hello("Puchi"))
}
