package main

import "fmt"

// this program has a data race condition
func main() {
	var data int
	go func() { data++ }() // shared memory access (not atomic)
	if data == 0 {         // reading the value is also not atomic
		fmt.Println("the value is 0.")
	} else {
		fmt.Printf("the value is %v.\n", data) // printing the value is also not atomic
	}

}
