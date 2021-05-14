package main

import "fmt"


func main()  {
	a := 5
	b := &a

	fmt.Println(a) // value of a 
	fmt.Println(&b) // memory location of a 
	fmt.Println(*b) // value stored in the memory location
	fmt.Printf("%T\n", b)

	*b = 10

	fmt.Printf("changed: %d\n", a)

}