package main

import "fmt"

var quit = make(chan bool)

func main() {

	command := ""
	data := make(chan int)

	go Fib(data)

	for {
		num := <-data
		fmt.Println(num)
		fmt.Scanf("%s", &command)
		if command == "quit" {
			quit <- true
			break
		}
	}
}

func Fib(c chan int) {

	x, y := 1, 1

	for {
		select {

		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Done calculating Fibonacci!")
			return
		}

	}
}
