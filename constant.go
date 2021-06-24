package main

import "fmt"

func main(){
    const J1 = "constant"
    const (
        J2 = 123
        J3 = 456
    )
    fmt.Print(J1, J2, J3, "\n")
}
