package main

import "fmt"

func myFunc() {
	defer fmt.Println("So Long and Thanks for All the Fish!")
	fmt.Println("myFunc running ...")
}
func main() {
	defer fmt.Println("We seem to have to have come to an end.")
	fmt.Println("main running ...")
	myFunc()
}
