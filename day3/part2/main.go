package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	INPUT    = "../input.txt"
	ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func createMapOfItems(items string) (set map[rune]bool) {
	set = make(map[rune]bool)
	for _, item := range items {
		set[item] = true
	}
	return
}

func getBadge(group []string) rune {
	map1 := createMapOfItems(group[0])
	map2 := createMapOfItems(group[1])
	for _, v := range group[2] {
		if map1[v] && map2[v] {
			return v
		}
	}
	return 0
}

func getPriority(input string) int {
	l := strings.Split(ALPHABET, "")
	for i, v := range l {
		if input == v {
			return i + 27
		} else if input == strings.ToLower(v) {
			return i + 1
		}
	}
	return 0
}

func main() {
	//Read input file
	input, _ := os.Open(INPUT)
	defer input.Close()
	sc := bufio.NewScanner(input)

	var priority int

	var group []string
	for sc.Scan() {
		rs := sc.Text()
		group = append(group, rs)
		if len(group) == 3 {
			badge := string(getBadge(group))
			p := getPriority(badge)
			priority += p
			fmt.Println("Badge:", badge, "Priority:", p, "Total Priority:", priority)
			group = nil
		}
	}
	fmt.Print(priority)
}
