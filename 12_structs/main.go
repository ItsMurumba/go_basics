package main

import ("fmt"
	"strconv")

// Person : Define person struct
type Person struct{
	// firstName string
	// lastName string
	// city string
	// gender string
	// age int

	firstName, lastName, city, gender string
	age int
}

//	Greeting method (value receiver)
func (p Person) greet() string  {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

//	hasBirthday method (pointer receiver)
func (p *Person) hasBirthday()  {
	p.age++
}

//	getMarried method (pointer receiver)
func (p *Person) getMarried(spouseLastName string)  {
	if p.gender == "m" {
		return
	}else{
		p.lastName = spouseLastName
	}
}

func main()  {
	person1 := Person{firstName: "Mercy", lastName: "Nomzano", city: "Nairobi", gender: "f", age: 20}

	person2 := Person{"Kelvin", "Murumba", "New York", "m", 26}

	// person1.age++
	// fmt.Println(person1)
	// fmt.Println(person1.firstName)
	person1.hasBirthday()
	person1.getMarried("Kelvin")
	person2.getMarried("Thompson")

	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}