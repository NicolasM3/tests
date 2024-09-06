package main

import (
	"time"
)

func main() {

	for i := 0; i < 3; i++ {

		// 1 - shadow i
		// i scapes to the heap
		// i := i
		// go func() {
		// 	fmt.Println(i)
		// }()

		// 2 - pass the value of i to the goroutine
		// n scapes to the heap
		// go func(n int) {
		// 	fmt.Println(n)
		// }(i)

		// 3 - create a goroutine to fmt.Println(i)
		// i escapes to the heap
		go println(i)
	}

	time.Sleep(time.Second)
}
