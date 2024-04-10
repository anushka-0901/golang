//Absolute value of a float number
//Golang program to print the absolute value of float number
package main
import "fmt"
import "math"
func main() {
	var val float64 = -19.25
	fmt.Printf("Absolute value of %f is %f",val,math.Abs(val))
}