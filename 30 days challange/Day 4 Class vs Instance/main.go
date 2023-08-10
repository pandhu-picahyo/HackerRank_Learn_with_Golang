package main

import "fmt"

type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	//Add some more code to run some checks on initialAge
	// start of code
	if initialAge < 0 {
		fmt.Println("Age is not valid, setting age to 0.")
		p.age = 0
	} else {
		p.age = initialAge
	}

	// end of code
	return p
}

func (p person) amIOld() {
	//Do some computation in here and print out the correct statement to the console
	// start of code
	if p.age < 13 {
		fmt.Println("You are young.")
	} else if p.age >= 13 && p.age < 18 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are old.")
	}

	// end of code
}

func (p person) yearPasses() person {
	//Increment the age of the person in here
	// start of code

	p.age++

	// end of code

	return p
}

func main() {
	var T, age int

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&age)
		p := person{age: age}
		p = p.NewPerson(age)
		p.amIOld()
		for j := 0; j < 3; j++ {
			p = p.yearPasses()
		}
		p.amIOld()
		fmt.Println()
	}
}
