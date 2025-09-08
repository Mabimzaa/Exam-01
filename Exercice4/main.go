package main

import "fmt"

func PGCD(a, b int) int {
	if b == 0 {
		return a
	}
	return PGCD(b, a%b)
}

func main() {
	fmt.Println(PGCD(48, 18))
}
