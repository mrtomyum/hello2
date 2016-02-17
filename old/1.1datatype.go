package main
import "fmt"

func main() {
	var p int
//	var pn *int
	pn := new(*int)
	*pn = 10
	fmt.Println("Pointer1", p, &p, pn)
	fmt.Println("Pointer2", p, &p, *pn)

}
