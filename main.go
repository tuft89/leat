package main

import "fmt"

func main() {
	fmt.Println(romanToInt("IX"))
}

func romanToInt(s string) int {
	runes := []rune(s)
	sum := 0
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	for i := 0; i < len(runes); i++ {
		if i == 0 {
			switch {
			case runes[i] == rune('I'):
				sum = sum + 1
			case runes[i] == rune('V'):
				sum = sum + 5
			case runes[i] == rune('X'):
				sum = sum + 10
			case runes[i] == rune('L'):
				sum = sum + 50
			case runes[i] == rune('C'):
				sum = sum + 100
			case runes[i] == rune('D'):
				sum = sum + 500
			case runes[i] == rune('M'):
				sum = sum + 1000
			}

		} else {
			if runes[i] == rune('I') && (runes[i-1] != rune('V') && runes[i-1] != rune('X')) {
				sum = sum + 1
			} else if runes[i] == rune('I') && (runes[i-1] == rune('V') || runes[i-1] == rune('X')) {
				sum = sum - 1
			} else if runes[i] == rune('X') && (runes[i-1] != rune('L') && runes[i-1] != rune('C')) {
				sum = sum + 10
			} else if runes[i] == rune('X') && (runes[i-1] == rune('L') || runes[i-1] == rune('C')) {
				sum = sum - 10
			} else if runes[i] == rune('C') && (runes[i-1] != rune('D') && runes[i-1] != rune('M')) {
				sum = sum + 100
			} else if runes[i] == rune('C') && (runes[i-1] == rune('D') || runes[i-1] == rune('M')) {
				sum = sum - 100
			} else if runes[i] == rune('V') {
				sum = sum + 5
			} else if runes[i] == rune('L') {
				sum = sum + 50
			} else if runes[i] == rune('D') {
				sum = sum + 500
			}
		}
	}
	return sum
}
