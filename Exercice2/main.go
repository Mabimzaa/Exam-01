package main
import "fmt"
func CountLetters(s string) int {
	for _, char := range s {
		if char < '' || char > 'z' {
