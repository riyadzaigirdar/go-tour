package main

import "fmt"

func main(){
	// ids := []int{12, 32,23,12 ,45 ,1, 14}

	// for index, value := range ids{	
	// 	fmt.Printf("index: %d value: %d\n", index, value)
	// }

	// for _, id := range ids{
	// 	fmt.Printf("id: %d\n", id)
	// }

	// add all ids together
	// sum := 0
	// for _, value := range(ids){
	// 	sum = sum + value
	// }

	// fmt.Println(sum)
	
	// range with maps

	body := map[string]string{"name": "riyad", "age": "12"}

	for k,v := range body{
		fmt.Printf("key: %s value: %s\n", k, v)
	}

	// key value only need to be declared on slices not on maps

	for k := range(body){
		fmt.Printf("key_only: %s\n", k)
	}

}