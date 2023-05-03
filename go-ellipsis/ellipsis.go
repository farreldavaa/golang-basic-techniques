package main

import "fmt"

func main() {
	sayHello()
	sayHello("Devin")
	sayHello("Farrel", "Athallah", "Alfian", "Rafli")
}

// using Ellipsis
func sayHello(names ...string) {
	for _, n := range names {
		fmt.Printf("Hello %s\n", n)
	}
}
