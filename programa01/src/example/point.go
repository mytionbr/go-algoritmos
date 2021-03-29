package main

func main() {
	i := 7
	inc(&i)
	println(&i)
	println(i)
}

func inc(x *int) {
	*x++
}
