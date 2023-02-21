package main

import (
	"fmt"
	"math/rand"
)

func main() {

	number := rand.Intn(100)

	fmt.Println("Witaj! Podaj liczbe")

	var answer int

	for {
		fmt.Scan(&answer)
		if answer == number {
			fmt.Println("Brawo! Udało Ci się zgadnąć liczbę!")
			break
		} else if answer < number {
			fmt.Println("Twoja liczba jest mniejsza od mojej")
		} else if answer > number {
			fmt.Println("Twoja liczba jest wieksza od mojej")
		}
	}

}
