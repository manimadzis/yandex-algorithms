package main

import "fmt"

func main() {

	var n int

	fmt.Scanf("%d", &n)

	buttons := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &buttons[i])
	}

	var k int

	fmt.Scanf("%d", &k)

	for i := 0; i < k; i++ {
		var key int
		fmt.Scanf("%d", &key)
		buttons[key-1]--
	}

	for _, value := range buttons {
		var ans string

		if value < 0 {
			ans = "YES"
		} else {
			ans = "NO"
		}

		fmt.Println(ans)
	}

}
