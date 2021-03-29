package main

import (
	"fmt"
)

func main() {
	x := 5
	bro := "Bro, isso é "

	if x > 6 {
		bro += "mais que seis"
	} else {
		bro += "menos que seis"
	}

	fmt.Println(bro)
}
