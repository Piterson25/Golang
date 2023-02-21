package main

import "fmt"

func main() {

	fmt.Println("Ile masz lat?")

	var age int

	fmt.Scanln(&age)

	var months int
	var days int

	for i := 0; i < age; i++ {
		months += 12
		days += 365
	}

	fmt.Println("Zyjesz", months, "miesiecy i", days, "dni")
}
