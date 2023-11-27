package main

import (
	"fmt"
	"github.com/tomprogramuje/card_validator/validation"
)

func main() {
	cardNumber := "4539059644464842"
	fmt.Println(validation.checkCard(cardNumber))
}
