package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	i := 0
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println(i, "\n")

	var counter int = 0
	for counter < 5 {
		fmt.Println(counter)
		counter++
	}
	fmt.Println(counter, "\n")

	counter = 0
	for {
		if counter > 5 {
			break
		}
		counter++
	}
	fmt.Println(counter, "\n")

}
