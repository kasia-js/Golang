//1. Basic solution
// package main

// import (
// 	"fmt"
// )

// func Sqrt(x float64) float64 {
//   z := 1.0
//   for i := 1; i <= 10; i++ {
//     z -= (z*z - x) / (2*z)
//   }
//   return z
// }

// func main() {
// 	fmt.Println(Sqrt(2))
// }

// Part 2.  the loop condition to stop once the value has stopped changing (or only changes by a very small amount 
	// package main

	// import (
	// 	"fmt"
	// 		"math"
	// )
	
	// func Sqrt(x float64) float64 {
	// 		 z := 1.0 
	// 	 c := 0
	// 		 var t float64
	// 		 for {
	// 			z, t = z - (z*z - x) / (2*z), z
	// 			c++
	// 			 if math.Abs(t-z) < 1e-6 {
	// 				 break
	// 	}
	//  }
	//  fmt.Printf("%d\n", c-1)
	//  return z
	// }
	
	// func main() {
	// 		guess := Sqrt(2)
	// 	expected := math.Sqrt(2)
	// 	fmt.Printf("Guess: %v, Expected: %v, Error: %v", guess, expected, math.Abs(guess - expected))
	// }
