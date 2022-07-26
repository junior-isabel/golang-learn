package main

import (
	"fmt"

	"github.com/junior-isabel/vector"
)

func main() {

	// exercicio 1
	//fizzBuzz()

	// exercicio 2
	/*	fmt.Println("Prime numbers less than 20:")

		for number := 1; number <= 20; number++ {
			if findprimes(number) {
				fmt.Printf("%v ", number)
			}
		}
	*/

	// exercicio 3
	/*
		val := 0
		for {

			fmt.Print("Enter number: ")
			fmt.Scanf("%d", &val)

			if val < 0 {
				panic("Allow is not negative number")
			} else if val == 0 {
				fmt.Println("0 is neither negative nor positive")
			} else {
				fmt.Println("You entered:", val)
			}
		}
	*/
	// conversor de numerção romana para decimal
	/*
		var numberRoman string
		fmt.Print("Enter number Roman: ")
		fmt.Scanf("%s", &numberRoman)
		result := convRomanToDecimal(numberRoman)
		println(numberRoman, " => ", result)
	*/

	vetor := vector.Vector{3, 4}

	fmt.Println(vetor)
	fmt.Printf("mod: %v\n", vetor.Mod())
	fmt.Printf("normalize: %v\n", *vetor.Normalize())
}

func convRomanToDecimal(x string) int {

	simbolos := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}

	var acc int

	for i := 1; i <= len(x); i++ {
		beforeValue, exist := simbolos[string(x[i-1])]
		if !exist {
			panic("Number Roman Invalid!")
		}
		if i == len(x) {
			return acc + beforeValue
		}
		value := simbolos[string(x[i])]
		if beforeValue < value {
			acc -= beforeValue
		} else {
			acc += beforeValue
		}
	}
	return acc
}

func fibbonnaci(x int) ([]int, int) {

	if x < 2 {

		panic("Numero X < 2")
	}
	fibbo := []int{1, 1}

	for i := 2; i < x; i++ {

		fibbo = append(fibbo, fibbo[i-2]+fibbo[i-1])
	}

	return fibbo, fibbo[x-1]
}

func findprimes(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	if number > 1 {
		return true
	} else {
		return false
	}
}
func fizzBuzz() {

	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%5 == 0:
			fmt.Println("Buzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		default:
			fmt.Println(i)
		}
	}
}
