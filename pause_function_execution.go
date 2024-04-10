//Golang program to pause function execution for the specified duration
package main
import "fmt"
import "time"
 func main() {
    fmt.Println("Hello world")
	time.Sleep(500*time.Millisecond)
    fmt.Println("Hello India")
	time.Sleep(8*time.Second)
	fmt.Println("Hello,Welcome to golang")
 }