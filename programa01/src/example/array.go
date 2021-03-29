package main

import (
	"fmt"
)

func main() {
	var a [100]int

	fmt.Println(a)

	a[1] = 405
	a[10] = 50
	a[3] = a[1] + a[10]

	fmt.Println(a)

	b := []int{1, 2, 3, 4}
	b = append(b, 2, 3, 4)

	fmt.Println(b)

	// Pelo que parece isso não funfa
	//c := []string{'matheus','sousa','de','jesus'}

	//Acho que é o mesmo esquema do c e do java,
	// aspas simples = caractere
	// aspas duplas =  string

	c := []string{"matheus", "sousa"}

	fmt.Println(c)
}
