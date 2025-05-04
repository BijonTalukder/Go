package main

import "fmt"

func processOperation(a int, b int, fn func(int, int)) {
    fn(a, b)
}

func add(a int, b int) {
    z := a + b
    fmt.Println(z)
}

func main() {
    processOperation(10, 20, add)
}