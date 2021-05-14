package main

import "fmt"

func adderInit() func(int) int{
	sum := 0
	return func (value int) int {
		sum += value
		return sum
	}
}

func main(){
	adder := adderInit()
	// adder(5)
	// result := adder(9)
	// fmt.Println(result)
	sum := 0
	for i := 0; i<10; i++{
		// sum += adder(i)
		sum = adder(i)
		// fmt.Println(adder(i))
	}
	fmt.Println(sum)
}