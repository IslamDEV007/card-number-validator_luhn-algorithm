package main

import "fmt"

func luhn(number string) bool {
	sum := 0
	for i := len(number) - 1; i >= 0; i-- {
		digit := int(number[i] - '0')
		if i%2 == 0 {
			digit *= 2
		}
		if digit > 9 {
			digit -= 9
		}
		sum += digit
	}
	return (sum % 10) == 0
}

func main() {
	number := "4485275742308329"

	if luhn(number) {
		fmt.Println("This is valid")
	} else {
		fmt.Println("This not valid")
	}
}
