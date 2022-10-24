package main

import "time"

func GgoroutineTest() string {
	return "Dzaka Eryan"
}

func Says(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		println(s)
	}
}

func main() {
	n := GgoroutineTest()
	println(n)

	go Says("Hello World")
	Says("Come to stay")
}
