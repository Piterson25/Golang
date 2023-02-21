package main

import "fmt"

func main() {

	fmt.Println("Ile masz lat?")

	var age float64
	fmt.Scanln(&age)
	fmt.Println("Twój wiek na innych planetach:")

	merkury := 0.24085
	wenus := 0.61521
	mars := 1.88089
	jowisz := 11.8622
	saturn := 29.4577
	uran := 85.0153
	neptun := 164.788

	fmt.Println("Masz tyle lat na merkurym:", age/merkury)
	fmt.Println("Masz tyle lat na wenus:", age/wenus)
	fmt.Println("Masz tyle lat na marsie:", age/mars)
	fmt.Println("Masz tyle lat na jowiszu:", age/jowisz)
	fmt.Println("Masz tyle lat na saturnie:", age/saturn)
	fmt.Println("Masz tyle lat na uranie:", age/uran)
	fmt.Println("Masz tyle lat na neptun:", age/neptun)
}
