package main

import "fmt"

func main() {

	arr := [...]string{"a", "b", "c", "d", "e", "f"}
	s1 := arr[1:4]
	s2 := arr[2:5]

	fmt.Printf("s1: %s\n", s1)
	fmt.Printf("s2: %s\n", s2)

	s1[2] = "z"
	fmt.Printf("%s\n", s1)
	fmt.Printf("%s\n", s2)
	fmt.Printf("%s\n", arr)

}
