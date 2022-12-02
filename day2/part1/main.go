package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calcPoints(oponent string, you string) int {
	result := 0
	switch you {
	case "X":
		fmt.Println("You choose Rock")
		result = result + 1
		switch oponent {
		case "A":
			fmt.Println("The Oponent choose Rock")
			fmt.Println("")
			result = result + 3
		case "B":
			fmt.Println("The Oponent choose Paper")
			result = result + 0
		case "C":
			fmt.Println("The Oponent choose Scissors")
			result = result + 6
		default:
			fmt.Println("Unexpected input", oponent)
		}
	case "Y":
		fmt.Println("You choose Paper")
		result = result + 2
		switch oponent {
		case "A":
			fmt.Println("The Oponent choose Rock")
			result = result + 6
		case "B":
			fmt.Println("The Oponent choose Paper")
			result = result + 3
		case "C":
			fmt.Println("The Oponent choose Scissors")
			result = result + 0
		default:
			fmt.Println("Unexpected input", oponent)
		}
	case "Z":
		fmt.Println("You choose Scissors")
		result = result + 3
		switch oponent {
		case "A":
			fmt.Println("The Oponent choose Rock")
			result = result + 0
		case "B":
			fmt.Println("The Oponent choose Paper")
			result = result + 6
		case "C":
			fmt.Println("The Oponent choose Scissors")
			result = result + 3
		default:
			fmt.Println("Unexpected input", oponent)
		}
	default:
		fmt.Println("Unexpected input", you)
	}
	fmt.Println("Round score:", result)
	return result
}

func main() {
	total := 0
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		charts := strings.Fields(scanner.Text())
		total = total + calcPoints(charts[0], charts[1])
	}
	fmt.Println("TOTAL SCORE:", total)
}
