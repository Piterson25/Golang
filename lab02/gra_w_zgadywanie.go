package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type score struct {
	tries int
	name  string
	date  string
	guess int
}

func nameExist(name string, gameScores []score) bool {
	for _, s := range gameScores {
		if s.name == name {
			return true
		}
	}
	return false
}

func bestScore(newScore score, gameScores []score) bool {
	for _, s := range gameScores {
		if s.tries < newScore.tries {
			return false
		}
	}
	return true
}

func replaceScore(newScore score, gameScores []score) {
	for _, s := range gameScores {
		if s.name == newScore.name && s.tries > newScore.tries {
			s = newScore

			if bestScore(newScore, gameScores) {
				fmt.Println("Pobiles rekord globalny!")
			} else {
				fmt.Println("Pobiles swoj rekord osobisty!")
			}
			break
		}
	}
}

func game(gameScores *[]score) bool {
	var answer string

	number := rand.Intn(101)

	var proby int

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

	game_score := score{tries: proby, name: imie, date: time.Now().Format("01-02-2006"), guess: number}

	if !nameExist(imie, *gameScores) {
		*gameScores = append(*gameScores, game_score)
	} else {
		replaceScore(game_score, *gameScores)
	}

	return false
}

func main() {

	fmt.Println("Witaj w grze w zgadywanie!")

	fmt.Println("Teraz będziesz zgadywać liczbę, którą wylosowałem (0-100)")
	fmt.Println("Jeśli w jakimś momencie chcesz zakończyć, napisz (koniec) ")

	var answer string

	gameScores := []score{}

	file, err := os.Open("hall_of_fame.txt")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	i := 0
	readScore := score{}
	for scanner.Scan() {

		switch i {
		case 0:
			text, _ := strconv.Atoi(scanner.Text())
			readScore.tries = text
		case 1:
			readScore.name = scanner.Text()
		case 2:
			readScore.date = scanner.Text()
		case 3:
			text, _ := strconv.Atoi(scanner.Text())
			readScore.guess = text
			gameScores = append(gameScores, readScore)
		}
		if i == 3 {
			i = 0
		} else {
			i++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

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
		f := fmt.Sprintf("%d %s %s %d\n", s.tries, s.name, s.date, s.guess)
		file.WriteString(f)
		fmt.Println(s.tries, s.name, s.date, s.guess)
	}

	file.Close()

}
