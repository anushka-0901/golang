//golang program to convert celsius to farenheit
package main
import "fmt"
func main() {
	var  ftemp float64 = 0
	var ctemp float64 = 0
	fmt.Printf("Enter temperature in celsius:")
	fmt.Scanf("%f",&ctemp)
	ftemp=(ctemp*1.8)+32
	fmt.Printf("Temperature in farenheit: %.2f",ftemp)
}