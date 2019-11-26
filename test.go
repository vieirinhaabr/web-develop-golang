package main

import "fmt"

type Person struct {
	Fname string
	Lname string
}

type SecretAgent struct {
	Person
	Licensetokill bool
}

func (p Person) speak(){
	fmt.Println(p.Fname, p.Lname, `says, "Good morning, James."`)
}

func main(){
	xi := []int{11, 22, 33, 42}
	fmt.Println(xi[0])

	m := map[string]int{
		"Tom": 40,
		"Kate": 25,
	}
	fmt.Println(m["Tom"])

	p1 := Person {
		"Money",
		"Batch",
	}
	fmt.Println(p1)

	p1.speak()

	James := SecretAgent{
		Person{
			"James",
			"Bond",
		},
		true,
	}
	fmt.Println(James)
}