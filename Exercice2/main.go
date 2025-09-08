package main

import "fmt"

func CountLetters(s string) int {
	count := 0
	for _, char := range s {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountLetters("Melvin, le beau gosse!"))
	fmt.Println(CountLetters("1234!@#$"))
	fmt.Println(CountLetters("J'aime la promo"))

}
