package main

import (
	"fmt"

	"github.com/lekting/go-validator"
)

type RegisterRequest struct {
	UserName string
	LastName string
}

func main() {
	registerRequst := &RegisterRequest{
		UserName: "Mikla",
		LastName: "Booana",
	}

	validatorScheme := validator.CreateValidationScheme(false, validator.Scheme{
		"UserName": *validator.GetValidation().NotEmpty("cant be empty"),
		"LastName": *validator.GetValidation().Len(2, ""),
	})

	errors, valid := validatorScheme.Validate(registerRequst)

	if valid {
		fmt.Println("scheme is valid")
		return
	}

	for fieldName, errorMsg := range errors {
		fmt.Printf("Error: %s %s", fieldName, errorMsg)
	}
}
