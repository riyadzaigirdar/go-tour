package main

import (
	"fmt"
)


func main(){
	// var arr [2]string
	// arr = append(arr, "sdasd")
	// arr = append(arr, "asd")
	// arr = append(arr, "asdasd")

	arr := []string{"apple", "orange"}

	arr = append(arr, "grape")
	fmt.Println(arr)
}