package main

import (
	"fmt"
	"time"
)

func main() {
	defer deferTest(time.Now())
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Good")
	fmt.Println("Bye")
}

func deferTest(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
