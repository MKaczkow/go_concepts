package main

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			println("recovered:", v)
		}
	}()
	println(60 / i)
}

func main() {
	for _, val := range []int{1, 2, 0, 3} {
		div60(val)
	}
}
