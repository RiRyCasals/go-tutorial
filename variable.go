package main
import "fmt"

func main(){
    var i1 int
    i1 = 1

    var i2 int = 2
    var i3 = 3

    var (
        i4 int = 4
        i5 int = 5
    )

    i6 := 6

    name, age := "Jhon", 20

    fmt.Print(i1, i2, i3, i4, i5, i6, "\n")
    fmt.Print(name, age, "\n")


    const J1 = "constant"

    const (
        J2 = 123
        J3 = 456
    )

    fmt.Print(J1, J2, J3, "\n")
}
