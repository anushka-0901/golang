package main
import "fmt"
func countDown(number int){
	//display the number
	if number > 0{
	fmt.Println(number)
	//recursive call by decreasing number
    countDown(number - 1)
}
else {
	//ends the recursive function
	fmt.Println("countdown stopes")
   }
}
func main() {
	fmt.Println("countdown starts")
	countDown(3)
}