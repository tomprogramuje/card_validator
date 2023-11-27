package validation

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func cardNumberParser(cardNumber string) []int {
	var cardNumberArr []int
	for _, n := range cardNumber {
		cardNumberArr = append(cardNumberArr, int(n-'0'))
	}
	return cardNumberArr
}

func CheckCardNumber(cardNumber string) (bool, string) {
	cardNumberArr := cardNumberParser(cardNumber)

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
	cardNumberArr := cardNumberParser(cardNumber)

	ve := cardNumberArr[:4]
	if reflect.DeepEqual(ve, []int{4, 0, 2, 6}) ||
		reflect.DeepEqual(ve, []int{4, 5, 0, 8}) ||
		reflect.DeepEqual(ve, []int{4, 8, 4, 4}) ||
		reflect.DeepEqual(ve, []int{4, 9, 1, 3}) ||
		reflect.DeepEqual(ve, []int{4, 9, 1, 7}) {
		return "Visa Electron"
	}
	ve = cardNumberArr[:6]
	if reflect.DeepEqual(ve, []int{4, 1, 7, 5, 0, 0}) {
		return "Visa Electron"
	}

	if cardNumberArr[0] == 4 {
		return "Visa"
	}

	ma := cardNumberArr[:4]
	if reflect.DeepEqual(ma, []int{5, 0, 1, 8}) ||
		reflect.DeepEqual(ma, []int{5, 0, 2, 0}) ||
		reflect.DeepEqual(ma, []int{5, 0, 3, 8}) ||
		reflect.DeepEqual(ma, []int{5, 8, 9, 3}) ||
		reflect.DeepEqual(ma, []int{6, 3, 0, 4}) ||
		reflect.DeepEqual(ma, []int{6, 7, 5, 9}) ||
		reflect.DeepEqual(ma, []int{6, 7, 6, 1}) ||
		reflect.DeepEqual(ma, []int{6, 7, 6, 2}) ||
		reflect.DeepEqual(ma, []int{6, 7, 6, 3}) {
		return "Maestro"
	}

	mc := cardNumberArr[:2]
	if reflect.DeepEqual(mc, []int{5, 1}) ||
		reflect.DeepEqual(mc, []int{5, 2}) ||
		reflect.DeepEqual(mc, []int{5, 3}) ||
		reflect.DeepEqual(mc, []int{5, 4}) ||
		reflect.DeepEqual(mc, []int{5, 5}) {
		return "MasterCard"
	}
	firstSixDigits, err := strconv.Atoi(cardNumber[:6])
	if err != nil {
		log.Println(err)
	}
	if firstSixDigits >= 222100 && firstSixDigits <= 272099 {
		return "MasterCard"
	}

	return "Unknown"
}
