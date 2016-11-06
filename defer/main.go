package main

import "fmt"

func main() {
	defer fmt.Println("We seem to have to have come to an end.")
	defer fmt.Println("So Long and Thanks for All the Fish!")

	fmt.Println("main running ...")
}
