package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calcPoints2(oponent string, you string) int {
	result := 0
	switch you {
	case "X":
		fmt.Println("You need to lose")
		result = result + 0
		switch oponent {
		case "A":
			fmt.Println("The Oponent choose Rock")
			fmt.Println("Yoy need choose Scissors")
			result = result + 3
		case "B":
			fmt.Println("The Oponent choose Paper")
			fmt.Println("Yoy need choose Rock")
			result = result + 1
		case "C":
			fmt.Println("The Oponent choose Scissors")
			fmt.Println("Yoy need choose Paper")
			result = result + 2
		default:
			fmt.Println("Unexpected input", oponent)
		}
	case "Y":
		fmt.Println("You need to draw")
		result = result + 3
		switch oponent {
		case "A":
			fmt.Println("The Oponent choose Rock")
			fmt.Println("Yoy need choose Rock")
			result = result + 1
		case "B":
			fmt.Println("The Oponent choose Paper")
			fmt.Println("Yoy need choose Paper")
			result = result + 2
		case "C":
			fmt.Println("The Oponent choose Scissors")
			fmt.Println("Yoy need choose Scissors")
			result = result + 3
		default:
			fmt.Println("Unexpected input", oponent)
		}
	case "Z":
		fmt.Println("You need to Win")
		result = result + 6
		switch oponent {
		case "A":
			fmt.Println("The Oponent choose Rock")
			fmt.Println("Yoy need choose Paper")
			result = result + 2
		case "B":
			fmt.Println("The Oponent choose Paper")
			fmt.Println("Yoy need choose Scissors")
			result = result + 3
		case "C":
			fmt.Println("The Oponent choose Scissors")
			fmt.Println("Yoy need choose Rock")
			result = result + 1
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
		total = total + calcPoints2(charts[0], charts[1])
	}
	fmt.Println("TOTAL SCORE:", total)
}
