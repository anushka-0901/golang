//golang program to get current gate using date function
package main
import "fmt"
import "time"
func main() {
	YYYY,MM,DD := time.Now().Date()
	fmt.Println("%d/%s/%d",DD,MM,YYYY)
}