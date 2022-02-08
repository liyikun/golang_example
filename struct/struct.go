package main

import "fmt"

type Person struct {
	name string
	age  int
}

func Factory(name string) *Person {
	instance := &Person{
		name: name,
		age:  20,
	}

	return instance
}

func (r *Person) add_age() {
	r.age += 1
}

func (r Person) get_age() string {
	return fmt.Sprintf("age is %d", r.age)
}

func main() {

	peter := Person{
		name: "peter",
		age:  12,
	}

	fmt.Println(peter)
	fmt.Println(peter.age)

	packer := &Person{
		name: "packer",
		age:  23,
	}

	fmt.Println(packer.age)
	fmt.Println(packer.name)

	v := Factory("hello")

	fmt.Println(v.name)
	fmt.Println(v.age)

	v.add_age()

	fmt.Println(v.age)
	fmt.Println(v.get_age())

}
