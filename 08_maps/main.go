package main

import "fmt"


func main(){
	body := make(map[string]string)

	body["name"] = "riyad"
	body["age"] = "26"
	body["occupation"] = "Software Developer"

	fmt.Println(body)

	// body2 := map[string]string{"name": "riyad", "age": "26"}

	// fmt.Println(body2)
	delete(body, "age")

	fmt.Println(body)
}