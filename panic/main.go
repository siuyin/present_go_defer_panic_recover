package main

import "fmt"

func onlyPositive(i int) {
	if i < 0 {
		panic(i) // Drop everything. Scream for help,yelling out "i"! // HL
	}
	fmt.Printf("Thanks, I got %v.\n", i)
}
func main() {
	defer fmt.Println("We seem to have to have come to an end.")
	onlyPositive(1) // Try changing to -1 // HL
	fmt.Println("After func call.")
}
