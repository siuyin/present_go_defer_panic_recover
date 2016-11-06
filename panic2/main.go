package main

import "fmt"

//010 OMIT
func onlyPositive(i int) {
	defer func() { // HL
		if r := recover(); r != nil {
			fmt.Printf("Hmm, I heard someone yelling out '%v'!\n", r)
			fmt.Println("That newbie doesn't know negative numbers!\nLet me fix it.")
		}
		fmt.Println("We have come to an end.")
	}()

	if i < 0 {
		panic(i) // Drop everything. Scream for help,yelling out "i"!
	}
	fmt.Printf("Thanks, I got %v.\n", i)
}

//020 OMIT
func main() {
	onlyPositive(1) // Try changing to -1
	fmt.Println("After func call.")
}

//030 OMIT
