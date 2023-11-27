package validation

import "fmt"

func CheckCard(cardNumber string) (bool, string) {
	var cardNumberArr []int

	// parses card number string to int and appending it to a slice
	for _, n := range cardNumber {
		cardNumberArr = append(cardNumberArr, int(n-'0'))
	}

	// iterates over slice from the end and multiplying every other number by two
	for i := len(cardNumberArr) - 1; i >= 0; i-- {
		if i%2 == 0 {
			cardNumberArr[i] *= 2
		}
	}

	// makes every two digit number into single digit by adding the digits of the number together
	for i := range cardNumberArr {
		if cardNumberArr[i] > 9 {
			cardNumberArr[i] -= 9
		}
	}

	// summs all the digits of card number
	sum := 0
	for i := range cardNumberArr {
		sum += cardNumberArr[i]
	}
	fmt.Println(sum)
	return isValid(sum), creditCardIssuer(cardNumber)
}

func isValid(cardNumberSum int) bool {
	return cardNumberSum%10 == 0
}

func creditCardIssuer(cardNumber string) string {
	if cardNumber[0] == '4' {
		return "Visa"
	}

	return "Unknown"
}