package main

import (
	"strings"
	"fmt"
)

func main() {
	J := "aA"
	S := "aAAbbbb"
	fmt.Println(numJewelsInStones(J, S))
}

func numJewelsInStones(J string, S string) int {
	JJ := strings.Split(J, "")
	count := 0
	for _, j := range JJ {
		count = count + strings.Count(S, j)
	}
	return count
}
