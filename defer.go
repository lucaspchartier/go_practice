package main

import (
	"fmt"
	"time"
)

func main() {
	defer deferTest(time.Now)

	fmt.Println("Good")
	fmt.Println("Bye")
}

func makeThingY(thing *string) string {
	*thing = "y"
	return "thing called"
}

func deferTest() {
	fmt.Println("Hi")
}
