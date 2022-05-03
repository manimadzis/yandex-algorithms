package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func find_stress(s string) int {
	for i, char := range s {
		if unicode.IsUpper(char) {
			return i
		}
	}
	return -1
}

func upper_count(s string) int {
	count := 0
	for _, char := range s {
		if unicode.IsUpper(char) {
			count++
		}
	}
	return count
}

func contains(slice []int, elem int) bool {
	for _, value := range slice {
		if value == elem {
			return true
		}
	}
	return false
}

func readline() string {
	reader := bufio.NewReader(os.Stdin)

	line, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return line
}

func main() {
	var n int

	fmt.Scanf("%d", &n)

	stresses := make(map[string][]int)

	for i := 0; i < n; i++ {
		var word string
		fmt.Scanln(&word)
		stress := find_stress(word)
		word = strings.ToLower(word)
		stresses[word] = append(stresses[word], stress)
	}

	var text string

	text = readline()

	words := strings.Fields(text)

	errors := 0

	for _, word := range words {
		if word == strings.ToLower(word) {
			errors++
		} else if upper_count(word) == 1 {
			stress := find_stress(word)
			word = strings.ToLower(word)

			if len(stresses[word]) > 0 && !contains(stresses[word], stress) {
				errors++
			}
		} else {
			errors++
		}
	}

	fmt.Print(errors)
}
