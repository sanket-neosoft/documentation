package main

import (
	"fmt"

	"./dog"
)

func main() {
	years := dog.Years(10)
	fmt.Printf("10 years of human is equals to %v of dog", years)
}
