package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

type score struct {
	name  string
	tries uint
}

func game(gameScores *[]score) bool {
	var answer string

	number := rand.Intn(101)

	var proby uint

	for {
		fmt.Print("Podaj liczbę: ")
		fmt.Scan(&answer)
		answerInt, err := strconv.Atoi(answer)
		if answer == "koniec" {
			fmt.Println("Żegnaj!")
			return true
		} else if err == nil {
			proby += 1
			if answerInt == number {
				fmt.Println("Gratulacje! Udało Ci się zgadnąć liczbę!")
				break
			} else if answerInt < number {
				fmt.Println("Twoja liczba jest za mała")
			} else if answerInt > number {
				fmt.Println("Twoja liczba jest za duża")
			}
		}
	}

	var imie string

	fmt.Print("Podaj swoję imię: ")
	fmt.Scan(&imie)

	game_score := score{name: imie, tries: proby}

	*gameScores = append(*gameScores, game_score)

	return false
}

func main() {

	fmt.Println("Witaj w grze w zgadywanie!")

	fmt.Println("Teraz będziesz zgadywać liczbę, którą wylosowałem (0-100)")
	fmt.Println("Jeśli w jakimś momencie chcesz zakończyć, napisz (koniec) ")

	var answer string

	gameScores := []score{}

	for {
		if game(&gameScores) {
			break
		}
		fmt.Print("Czy gramy jeszcze raz? [T/N] ")
		fmt.Scan(&answer)
		answer = strings.ToLower(answer)
		if answer == "t" {
			continue
		} else if answer == "n" || answer == "koniec" {
			fmt.Println("Żegnaj!")
			break
		}
	}

	sort.SliceStable(gameScores, func(i, j int) bool {
		return gameScores[i].tries < gameScores[j].tries
	})

	for _, s := range gameScores {
		fmt.Println(s.name, s.tries)
	}

}
