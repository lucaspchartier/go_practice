package main

import (
	"fmt"
	"time"
)

func main() {
	defer deferTest(time.Now())

	fmt.Println("Good")
	fmt.Println("Bye")
}

func deferTest(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
