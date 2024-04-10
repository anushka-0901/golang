//sorting a slice of integer in ascending order in golang
package main
import "fmt"
import "sort"
func main() {
	slice:=[]string{"honesty","is","the","best","policy"}
	sort.Strings(slice)
	fmt.Println("sorted slice:")
	for val,item:=range slice{
		fmt.Printf("%s",item)
	}
}