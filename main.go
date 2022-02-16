package main

import "fmt"

func main() {
	var n int
	ans := 0.0
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		ans += float64(i) * 10000 / float64(n)
	}
	fmt.Println(ans)
}
