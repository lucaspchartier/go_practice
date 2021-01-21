package main

import "fmt"

var thing string

func main() {
	thing = "x"
	defer deferTest(makeThingY(&thing))

	fmt.Println("Good")
	fmt.Println("Bye")
}

func makeThingY(thing *string) string {
	*thing = "y"
	return "thing called"
}

func deferTest(thingReturn string) {
	fmt.Println(thingReturn)
	fmt.Println("Hi")
}
