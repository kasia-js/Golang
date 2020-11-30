package main
import "fmt"

func main() {
	// in map order is not important. if order important use array
	//accessing values in map is faster than in arrays
	// map[typeOfKey]typeOfValue
	var mp map[string]int = map[string]int{
		"apple": 5,
		"pear": 6,
		"orange": 9,
	}
	//val - value of the key. 
	// ok - true or false depending if the key exists 
	// if the key does not exist the default value of 0 is assigned 
	val,ok := mp["apple"]
	fmt.Println(val,ok)
	fmt.Println(mp)
}