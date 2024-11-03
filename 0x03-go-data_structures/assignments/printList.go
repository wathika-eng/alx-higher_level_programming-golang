package golangDsa

import "fmt"

func PrintListInteger(myList []int) {
	for i := range myList {
		fmt.Println(myList[i])
	}
}
