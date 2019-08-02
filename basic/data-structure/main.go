package main

import "fmt"

func main() {
	type Person struct {
		name string
		age int
	}
	type Human struct {
		name     string
		behavior string
	}
	type User struct {
		Human
		Person
		name string
		address string
		phone string
	}

	p := Person{"Tyler Swift", 18}
	u := User{Human{"Json Mic", "walk"},Person{"Tyler Swift", 18}, "Jack Ma", "American", "18765434567"}

	fmt.Println("Name: ", p.name, "Age: ", p.age)
	fmt.Println(u.Human.name, u.Person.name, u.name)
}