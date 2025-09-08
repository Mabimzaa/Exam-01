package main

import "fmt"

func Binairedecimal(n int) string {
	switch n {
	case 0:
		return "0"
	case 1:
		return "1"
	}
	return Binairedecimal(n/2) + fmt.Sprint(n%2)
}

func main() {
	fmt.Println(Binairedecimal(0))
	fmt.Println(Binairedecimal(5))
	fmt.Println(Binairedecimal(10))
}
