// type assertion
// A story about a Gopher, named Gophy who joint the Gophers 
// but lost his identity in the process.
package main

import "fmt"

type I interface{}

var gophers map[uint]I = make(map[uint]I)

type gopher struct {
	Name string
}

func main() {

	g := AddToGophers("Goghy")

	fmt.Printf("Hello, %s\n", g.Name)
	fmt.Printf("Now %s is a %T, %s\n", g.Name, gophers[1], gophers[1])

	gg := GetGopher(1)
	fmt.Printf("Bye, %s\n", gg.Name)
}

func GetGopher(i uint) *gopher {
	g := gophers[i]
	// I wont my gopher identity back
	return g.(*gopher)
}

func AddToGophers(n string) gopher {
	g := new(gopher)
	g.Name = n
	gophers[1] = g
	return *g
}
