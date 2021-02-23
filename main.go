package main

import "fmt"

func main() {
	buy(2, 2)
	buy(1, 2, 3)
}

func buy(...nums int) {
	total := 0
	for _, num := range nums {
		total+=num
	}

	fmt.Println(total)
}
