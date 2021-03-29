package main

//pelo que parece não precisa ficar importando fmt

func main() {
	forNormal()
	forAnormal()
	forArray()
	forMap()

}

func soma(valor int, total int) int {
	total += valor
	return total
}

func forNormal() {
	total := 0

	for i := 0; i < 12; i++ {
		println(i)
		total = soma(i, total)
		println(total)
	}
}

func forAnormal() {
	i := 0
	for i < 10 {
		println(i)
		i++
	}
}

func forArray() {
	arr := [4]int{1, 2, 4, 5}

	for index, value := range arr {
		println("index", index, "value", value)
	}
}

func forMap() {
	arr := make(map[string]int)

	arr["dsfsaf"] = 1
	arr["dsfdfsfsaf"] = 2
	arr["fdsdf"] = 3

	for key, value := range arr {
		println("key", key, "value", value)
	}
}
