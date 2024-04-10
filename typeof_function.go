//find the variable using typeof() function
//golang program to find the variable type
//using reflect.typeof() function
package main
import(
	"fmt"
	"reflect"
)
func main() {
	//defining the variables
	a := 10
	b := 10.20
	c := "Hello"
	d := true
	e := []string{"India","USA","UK"}
	//Printing the values
	fmt.Println("a=",a) //integer variable
	fmt.Println("b=",b) //float variable
	fmt.Println("c=",c) //string variable
	fmt.Println("d=",d) //boolean variable
	fmt.Println("e=",e) //string array
	//Printing the type of the variables
	fmt.Println("type of a=",reflect.ValueOf(a).Kind())
	fmt.Println("type of b=",reflect.ValueOf(b).Kind())
	fmt.Println("type of c=",reflect.ValueOf(c).Kind())
    fmt.Println("type of d=",reflect.ValueOf(d).Kind())
	fmt.Println("type of e=",reflect.ValueOf(e).Kind())
}

