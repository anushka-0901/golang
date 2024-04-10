//len() and cap() functions
package main
import "fmt"
func main() {
	var numbers=  make([]int,3,5)
	numbers[0]=30
	numbers[2]=98
	printSlice(numbers)
}
func printSlice(x[]int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}