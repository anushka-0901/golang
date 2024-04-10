//Golang program to print ASCII value of characters of string
package main
import "fmt"
func main() {
	str := "Hello World"
	fmt.Println("ascii value of string:")
	for i:=0;i<len(str);i++{
        fmt.Printf("%c %d\n",str[i],str[i])
	}
}