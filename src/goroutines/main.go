package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type fibbonnaci struct {
	x int
	y float64
}

func main() {

	start := time.Now()
	ch := make(chan fibbonnaci, 4)

	for i := 1; i < 15; i++ {
		go Fib(ch, float64(i))
	}

	for i := 1; i < 15; i++ {
		data := <-ch
		fmt.Printf("Fib(%v): %v\n", data.x, data.y)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
func Fib(ch chan fibbonnaci, number float64) {
	x, y, z := 0.0, 1.0, 1.0

	for i := 0; i < int(number); i++ {
		x, y, z = y, x+y, y+z
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)
	ch <- fibbonnaci{x: int(number), y: x}
}

func send(ch chan string, message string) {

	ch <- message
}

func apiChecks(api string, ch chan string) {

	_, err := http.Get(api)

	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", err)
		return
	}
	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}
