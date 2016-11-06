package main

import "fmt"

func onlyPositive(i int) {
	if i < 0 {
		panic(i) // Drop everything. Scream for help,yelling out "i"! // HL
	}
	fmt.Printf("Thanks, I got %v.", i)
}
func main() {
	onlyPositive(1) // Try changing to -1 // HL
}
