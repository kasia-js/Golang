// &x memory location
// *y de-reference

// Part 1
// package main
// import "fmt"

// func main() {
// 	x := 7
// 	y := &x
// 	fmt.Println(x,y)
// 	*y = 8 
// 	fmt.Println(x,y)
// }

// //Part 2
// package main
// import "fmt"

// func changeValue(str *string) {
// 	*str = "changed" // de-reference
// }

// func changeValue2(str string) {
// 	str = "changed2"
// }

// func main() {
// 	toChange := "hello"
// 	changeValue(&toChange) // pass by memory location
// 	fmt.Println(toChange)

// 	toChange2 := "hello"
// 	changeValue2(toChange2)
// 	fmt.Println(toChange2)
// }