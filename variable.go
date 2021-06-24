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

    var f1 
    f1 = (i1)

    fmt.Print(f1, "\n")

}
