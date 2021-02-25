package main

import "fmt"

func main() {
	buy("milk", 3)
	buy("cookies", 4, 1)
	buy("chips", 4, 2)
}

func buy(item string, amount int, discount ...int) {
	if len(discount) > 0 {
		amount -= discount[0]
	}

	fmt.Printf("You bought %v for %d\n", item, amount)
}
