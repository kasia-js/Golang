// &x memory location
// *y de-reference

// Part 1
package main
import "fmt"

func main() {
	x := 7
	y := &x
	fmt.Println(x,y)
	*y = 8 
	fmt.Println(x,y)
}


