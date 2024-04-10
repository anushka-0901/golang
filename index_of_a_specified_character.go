//Golang program to get the index of a specified character in the string.
package main
import "fmt"
import "strings"
func main() {
	var str string = "Hello world"
	var ind int = 0
	ind = strings.Index(str,"W")
	fmt.Println("Index is :",ind)
}