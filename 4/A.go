package main

import "fmt"

func main() {
	var (
		n              int
		word_1, word_2 string
	)
	sino := make(map[string]string)

	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%s%s", &word_1, &word_2)
		sino[word_1] = word_2
		sino[word_2] = word_1
	}

	fmt.Scanln(&word_1)
	fmt.Println(sino[word_1])
}
