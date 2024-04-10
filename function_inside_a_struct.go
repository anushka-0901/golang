//function inside a struct in go
package main
import "fmt"
//initialize the function Rectangle
type Rectangle func(int,int) int
//create structure
type rectanglePara struct {
	length int
	breadth int
	color string
	//function as a field of struct
	rect Rectangle
}
func main() {
	//assign values to struct variables 
	result := rectanglePara{
		length:10,
		breadth:20,
		color:"red",
		rect:func(length int, breadth int) int {
			return length * breadth
		},
	}
	fmt.Println("color of rectangle:",result.color)
	fmt.Println("area of rectangle:\n",result.rect(result.length,result.breadth))
}