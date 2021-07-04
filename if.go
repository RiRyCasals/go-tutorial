package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			fmt.Println("0")
		} else if i%3 == 1 {
			fmt.Println("1")
		} else {
			fmt.Println("2")
		}
	}
}
