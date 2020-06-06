package main

import "fmt"

// struct
type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person
	licenseToKill bool
}

// function - reciever is key point
func (p person) speak() {
	fmt.Println(p.firstName, p.lastName, "says, Good morning, James.")
}

func (sa secretAgent) speak() {
	fmt.Println(sa.firstName, sa.lastName, "says, James Bond.")
}

// interface
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	// slices demo - hloding list of things
	xi := []int{1, 2, 3, 8, 9}
	fmt.Println(xi)

	// map demo
	testMap := map[string]int{
		"John": 45,
		"Bob":  30,
	}
	fmt.Println(testMap)

	// struct demo
	p1 := person{
		"Miss",
		"Melody",
	}
	fmt.Println(p1)

	// function demo
	p1.speak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	sa1.speak()
	sa1.person.speak()

	// interface, polymorphism demo
	saySomething(p1)
	saySomething(sa1)
}
