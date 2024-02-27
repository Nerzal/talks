package main

func drinkBeer(count ...int) {
	if len(count) == 0 {
		println("bleibst wohl heute nüchtern")
	}

	for _, i := range count {
		println("komm drink nomma", i, "Bier")
	}
}

func main() {
	drinkBeer()
}