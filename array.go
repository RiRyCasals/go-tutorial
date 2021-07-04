package main

import "fmt"

func main() {
	array := [3]string{}
	array[0] = "red"
	array[1] = "green"
	array[2] = "blue"
	fmt.Println(array[0], array[1], array[2])

	array2 := [3]string{"red", "green", "blue"}
	fmt.Println(array2[0], array2[1], array2[2])

	array3 := [...]string{"red", "green", "blue"}
	fmt.Println(array3[0], array3[1], array3[2])
}
