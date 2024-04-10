package main
import "fmt"
type Books struct {
	title string
	author string
	subject string
	book_id int
}
func main() {
	var Book1 Books //declare a book1 of type book
	var Book2 Books //declare a book2 of type book
	//Book1 specification
	Book1.title="Go programming"
	Book1.author="Mahesh Kumar"
	Book1.subject="Go programming tutorial"
	Book1.book_id=6495407
	//Book2 specification
	Book2.title="Telecom billing"
	Book2.author="Zara Ali"
	Book2.subject="Telecom billing tutorial"
	Book2.book_id=6495700
	//Print book1 info
	fmt.Printf("Book1 title:%s\n",Book1.title)
	fmt.Printf("Book1 auhtor:%s\n",Book1.author)
	fmt.Printf("Book1 subject:%s\n",Book1.subject)
	fmt.Printf("Book1 book_id:%d\n",Book1.book_id)
	//Print book2 info
	fmt.Printf("Book2 title:%s\n",Book2.title)
	fmt.Printf("Book2 auhtor:%s\n",Book2.author)
	fmt.Printf("Book2 subject:%s\n",Book2.subject)
	fmt.Printf("Book2 book_id:%d\n",Book2.book_id)
}