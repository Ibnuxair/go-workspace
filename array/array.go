package main

import "fmt"

func main() {

	/* var x [5]int

	x[2] = 6

	fmt.Printf("%d", x[2]) */

	//var x [5] = [5]{1,2,3,4,5}

	x := [...]int{1, 2, 3, 4}

	for i, v := range x {

		fmt.Printf("Index: %d  Value: %d\n", i, v)
	}

}
