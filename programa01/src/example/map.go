package main

func main() {
	produtos := make(map[string]int)

	produtos["abajur"] = 100
	produtos["mesa"] = 200
	produtos["cafeteira"] = 200

	delete(produtos, "cafeteira")

	println(produtos)
	println(produtos["abajur"])

}
