package main

type pessoa struct {
	name string
	age  int
}

func main() {
	p := pessoa{name: "matheus", age: 20}

	println(p.name)
	println(p.age)
}
