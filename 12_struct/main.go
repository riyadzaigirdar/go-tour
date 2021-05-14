package main

import (
	"fmt"
)

// Define person struct
type Person struct{
	// firstName string
	// lastname string
	// city string
	// gender string
	// age int
	firstName, lastname, city, gender string
	age int
}

// normal way pass the value as a parameter
// func gretting(p Person) string{
// 	return "hello " + p.firstName + " " + p.lastname + " you are " + strconv.Itoa(p.age) + " years old"
// }

// dynamic way write a function 
// func (p Person) greeting() string{
// 	return "hello " + p.firstName + " " + p.lastname + " you are " + strconv.Itoa(p.age) + " years old"
// }

// increment age if has bithday (pointer receiver)
// func (p *Person) hasBirthday() {
// 	p.age++
// }

// change lastname if female when ge married(pointer receiver)
func (p *Person) gettingMarried(spouseName string){
	if p.gender == "Male"{
		return
	}else{
		p.lastname = spouseName
		return
	}
}

func main(){
	// Init a person using struct
	person1 := Person{
		firstName: "riyad",
		lastname: "zaigirdar",
		city: "Dhaka",
		gender: "FeMale",
		age: 26,
	}
	// memory_location := &person1.age
	// // Init a person without giving property names just like positional arguments in python
	// person2 := Person{
	// 	"arman",
	// 	"zaigirdar",
	// 	"narayanganj",
	// 	"Male",
	// 	28,
	// }

	// fmt.Println(person1)
	// fmt.Println(person2)

	// // getting properties
	// fmt.Println("firstname",person1.firstName)
	// fmt.Println("lastname",person1.lastname)
	// fmt.Println("city",person1.city)
	// fmt.Println("gender",person1.gender)
	// fmt.Println("age",person1.age)

	// // memory location checke after and before modifying
	// // before
	// fmt.Println("before modifying", memory_location)
	// // property can be modified
	// person1.age ++
	// fmt.Println("change",person1.age)
	// // after
	// fmt.Println("after modifying", memory_location)
		
	// greeting message
	// fmt.Println(gretting(person1))
	// person1.hasBirthday()
	// fmt.Println(person1.greeting())
	person1.gettingMarried("rama")

	fmt.Println(person1.lastname)
}	