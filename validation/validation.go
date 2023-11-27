package validation

import "fmt"

func checkCard(cardNumber string) string {
	var cardNumberArr []int
	
	// parsing card number string to int and appending it to a slice
	for _, n := range cardNumber {
		cardNumberArr = append(cardNumberArr, int(n-'0'))
	}

	// iterating over slice from the end and multiplying every other number by two 
	for i := len(cardNumberArr) - 1; i >= 0; i-- {
		if i % 2 == 0 {
			cardNumberArr[i] *= 2
		}
	}

	// making every two digit number into single digit by adding the digits of the number together 
	for i := range cardNumberArr {
		if cardNumberArr[i] > 9 {
			cardNumberArr[i] -= 9
		}
	}

	// summing all the digits of card number 
	sum := 0
	for i := range cardNumberArr {
		sum += cardNumberArr[i]
	}

	return isValid(sum)
}

func isValid(cardNumberSum int) string {
	if cardNumberSum % 10 == 0 {
		return fmt.Sprintln("This card is valid")
	} else {
		return fmt.Sprintln("This card is not valid")
	}
}
