//change value of a map in golang
package main
import "fmt"
func main() {
	//create a map
	capital:=map[string]string{"Nepal":"Kathamandu","US":"New York"}
	fmt.Println("initial map:",capital)
	//change value of US to Washington DC
	capital["US"]="Washington DC"
	fmt.Println("Updated map:",capital)
} 