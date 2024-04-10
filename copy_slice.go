//program to copy one slice to another
package main
import "fmt"
func main() {
	//create two slices 
	primeNumbers:=[]int{2,3,5,7}
	numbers:=[]int{1,2,3}
	//copy lements of prime numbers to numbers
	copy{numbers,primeNumbers}
	//prime numbers
	fmt.Println("numbers:",numbers)
}