package golangDsa

import "fmt"

// prints a list of integers
func PrintListInteger(myList []int) {
	for i := range myList {
		fmt.Println(myList[i])
	}
}
