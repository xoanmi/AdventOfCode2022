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

func findDuplicate(l1 []string, l2 []string) string {
	items := make(map[string]int, 0)
	for _, itemList1 := range l1 {
		items[itemList1] = 0
		for _, itemlist2 := range l2 {
			if _, ok := items[itemlist2]; ok {
				return itemlist2
			}
		}
	}
	return ""
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

	for sc.Scan() {
		rs := strings.Split(sc.Text(), "")
		c1 := rs[0 : len(rs)/2]
		c2 := rs[len(rs)/2:]
		priority += getPriority(findDuplicate(c1, c2))
	}
	fmt.Print(priority)
}
