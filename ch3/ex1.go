package ch3

import "fmt"

func main() {

	var x = [...]int{10, 20, 30}
	var y = [3]int{1, 2, 3}

	if x == y {
		fmt.Println(true)
	}

}
