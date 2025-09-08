package main

import "fmt"

func BinaireDecimal(binaire string) int {
	if len(binaire) == 1 {
		if binaire == "0" {
			return 0
		}
		return 1
	}

	nb := int(binaire[len(binaire)-1] - '0')
	return BinaireDecimal(binaire[:len(binaire)-1])*2 + nb
}

func main() {
	fmt.Println(BinaireDecimal("1011"))
	fmt.Println(BinaireDecimal("1001"))
	fmt.Println(BinaireDecimal("1111"))
}
