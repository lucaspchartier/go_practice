package main

import "fmt"

func main() {
	sum(2, 2)
	sum(1, 2, 3)
}

func sum(...nums int) {
	total := 0
	for _, num := range nums {
		total+=num
	}

	fmt.Println(total)
}
