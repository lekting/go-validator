package main

import (
	"fmt"

	"github.com/lekting/validator"
)

type RegisterRequest struct {
	UserName string
	LastName string
}

func main() {
	// will return errors
	errors, valid := validator.ValidateVar("bad value", *validator.GetValidation().LenLte(5, "must be lower than 5 symbols"))

	if valid {
		fmt.Println("scheme is valid")
		return
	}

	for fieldName, errorMsg := range errors {
		fmt.Printf("Error: %s %s", fieldName, errorMsg)
	}
}
