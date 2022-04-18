package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// f, err := os.Create("log.txt")

	// if err != nil {
	// 	log.Println(err)
	// }

	f, _ := os.Open("log.txt")
	log.SetOutput(f)
	defer f.Close()

	file, err := os.Open("sanket.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	msg, err := io.ReadAll(file)
	fmt.Printf("%s", msg)

}
