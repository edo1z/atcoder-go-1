package main

import (
	"fmt"
	"strings"
)

func main() {
	targets := "atcoder"
	ans := "You can win"
	var s, t string
	fmt.Scan(&s, &t)
	cs := []rune(s)
	ct := []rune(t)
	for i := 0; i < len(cs); i++ {
		if cs[i] == ct[i] {
			continue
		} else if cs[i] == '@' && strings.Contains(targets, string(ct[i])) {
			continue
		} else if ct[i] == '@' && strings.Contains(targets, string(cs[i])) {
			continue
		} else {
			ans = "You will lose"
			break
		}
	}
	fmt.Println(ans)
}
