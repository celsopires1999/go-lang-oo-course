package main

import "fmt"

type Names []interface{}

func (n *Names) load() {
	*n = Names{
		"Celso",
		"Inês",
		"André",
		"Suzane",
		1,
	}
}

func (n *Names) print() {
	fmt.Println(*n)
}

func main() {

	var names Names

	names.load()
	names.print()
}
