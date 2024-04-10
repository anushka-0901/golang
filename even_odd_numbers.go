package main 
import "fmt"
func main() {
	var a int
	fmt.Printf("enter the number:\n")
	fmt.Scanf("%d",&a)
	if (a % 2 == 0){
		fmt.Println("a is an even number")
	}else{
		fmt.Println("a is an odd number")
	}
}