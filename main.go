package main
import "fmt"
func add(a int  , b int ) int {
	var  result int 
	result = a+b
	return result
}
func main(){
	var a int = 10
	var sum = add(a,5)

    fmt.Println("Sum is ", sum)
}