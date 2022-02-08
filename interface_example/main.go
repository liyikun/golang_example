package main

import "fmt"

type person interface {
	add_age(int)
	get_age() int
}

type student struct {
	age int
}

type teacher struct {
	age int
}

func (r *student) add_age(n int) {
	r.age += n
}

func (r *teacher) add_age(n int) {
	r.age += n
}

func (r student) get_age() int {
	return r.age
}

func (r teacher) get_age() int {
	return r.age
}

func get_info(i person) {
	fmt.Println(i.get_age())
	i.add_age(1)
	fmt.Println(i.get_age())
}

type base struct {
	name string
	num  int
}

type container struct {
	base
	text string
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

func main() {

	li := student{
		age: 10,
	}
	wang := teacher{
		age: 22,
	}

	get_info(&li)
	get_info(&wang)

	c := container{
		base: base{
			name: "liyikun",
			num:  10,
		},
		text: "hello",
	}

	fmt.Println(c.base.name)
	fmt.Println(c.name)
	fmt.Println(c.text)

	d := c

	fmt.Println(d.describe())

}
