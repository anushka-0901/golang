package main 
import "fmt"
func main() {
	var i,j int
	var m3[3][3] int
    m1:=[3][3]int{{0,1},{4,5},{8,9}}
	m2:=[3][3]int{{5,6},{7,8},{1,2}}
	fmt.Println("first matrix is:\n")
	for i=0;i<3;i++{
		for j=0;j<3;j++{
          fmt.Print(m1[i][j],"\t")
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("the second matrix is:\n")
	for i=0;i<3;i++{
		for j=0;j<3;j++{
			fmt.Print(m2[i][j],"\t")
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("addition of the two matrices is:\n")
	for i=0;i<3;i++{
		for j=0;j<3;j++{
			m3[i][j]=m1[i][j]+m2[i][j]
		}
	}
	for i=0;i<3;i++{
		for j=0;j<3;j++{
			fmt.Print(m3[i][j],"\t")
		}
		fmt.Println()
	}
}