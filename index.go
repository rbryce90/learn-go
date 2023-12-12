package main
import (
    "fmt"
)

func typesPlay () {
    // number
    var num int = 12
    fmt.Println(num)

    // string
    var myString = "myString"
    fmt.Println(myString)

    // string methods

}

func forLoop (){
    for i := 1; i <= 10; i++ {
        fmt.Println("i = ", i)
    }
}

func main () {
    fmt.Println("hello world")
    typesPlay()
    forLoop()

}
