package main

import "fmt"

type myJson struct{
	Name string `json:"name"`
	Message string `json:"message"`
}

// func getMessage(name string) []myJson{
// 	var data = myJson{
// 		Name: name,
// 		Message: "hello " + name,
// 	}
// 	var arr []myJson
// 	arr = append(arr, data)
// 	return arr
// }

// func getSum(num1, num2 int)int{
// 	return num1 + num2
// }

func main(){
	// fmt.Println(getMessage("riyad").Name)
	// fmt.Println(getMessage("riyad").Message)
	// fmt.Println(getMessage("riyad"))
	// fmt.Println(getSum(2,3))
	person := myJson{
		"riyad",
		"hello riyad",
	}
	fmt.Printf("%+v\n",person)
}