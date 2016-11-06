package main

import "fmt"

func onlyPositive(i int) {
	//030 OMIT
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("defer func: What's your problem?")
			fmt.Printf("panicker: %#v\n", r)
		}
		fmt.Println("We have come to an end.")
	}()
	//040 OMIT
	//010 OMIT
	if i < 0 {
		panic(struct {
			Op  string
			Val int
		}{"Incoming integer", i})
	}
	fmt.Printf("Thanks, I got %v.\n", i)
	//020 OMIT
}
func main() {
	//050 OMIT
	onlyPositive(1) // Try changing to -1
	//060 OMIT
}
