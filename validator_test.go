package validator

import (
	"testing"
)

type RegisterRequest struct {
	UserName string
}

type RegisterRequest2 struct {
	UserName int
}

func testScheme() *ValidationScheme {
	return CreateValidationScheme(false, Scheme{
		"UserName": *GetValidation().Len(1, "len = 1"),
	})

}

func TestValidationLenLte(t *testing.T) {

	goodStruct := &RegisterRequest{
		UserName: "good",
	}

	scheme := CreateValidationScheme(false, Scheme{
		"UserName": *GetValidation().LenLte(5, "len >= 5"),
	})

	errors, valid := scheme.Validate(goodStruct)

	if !valid {
		t.Fatalf("Scheme is right, something went wrong")
	}

	if errors != nil {
		t.Fatalf("errors val must be empty, something went wrong")
	}

	badStruct := &RegisterRequest{
		UserName: "bad val",
	}

	errors, valid = scheme.Validate(badStruct)

	if valid {
		t.Fatalf("Scheme is bad, but here is valid = true, something went wrong")
	}

	if errors == nil {
		t.Fatalf("errors val must be empty, something went wrong")
	}

	if errors["UserName"] == nil || len(errors["UserName"]) == 0 {
		t.Fatalf("Error structure is incorrect %v", errors)
	}

	if errors["UserName"][0] != "len >= 5" {
		t.Fatalf("Message on an error is incorrect `%s`, must be `len >= 5`", errors["UserName"][0])
	}

	badStruct2 := &RegisterRequest2{
		UserName: 5,
	}

	errors, valid = scheme.Validate(badStruct2)

	if valid {
		t.Fatalf("Scheme is bad, but here is valid = true, something went wrong")
	}

	if errors == nil {
		t.Fatalf("errors val must be empty, something went wrong")
	}

	if errors["UserName"] == nil || len(errors["UserName"]) == 0 {
		t.Fatalf("Error structure is incorrect %v", errors)
	}

	if errors["UserName"][0] != "must be a string" {
		t.Fatalf("Message on an error is incorrect `%s`, must be `must be a string`", errors["UserName"][0])
	}
}

func TestValidationLenGte(t *testing.T) {

	goodStruct := &RegisterRequest{
		UserName: "goods",
	}

	scheme := CreateValidationScheme(false, Scheme{
		"UserName": *GetValidation().LenGte(5, "len <= 5"),
	})

	errors, valid := scheme.Validate(goodStruct)

	if !valid {
		t.Fatalf("Scheme is right, something went wrong")
	}

	if errors != nil {
		t.Fatalf("errors val must be empty, something went wrong")
	}

	badStruct := &RegisterRequest{
		UserName: "bad",
	}

	errors, valid = scheme.Validate(badStruct)

	if valid {
		t.Fatalf("Scheme is bad, but here is valid = true, something went wrong")
	}

	if errors == nil {
		t.Fatalf("errors val must be empty, something went wrong")
	}

	if errors["UserName"] == nil || len(errors["UserName"]) == 0 {
		t.Fatalf("Error structure is incorrect %v", errors)
	}

	if errors["UserName"][0] != "len <= 5" {
		t.Fatalf("Message on an error is incorrect `%s`, must be `len <= 5`", errors["UserName"][0])
	}
}

func TestValidationLen(t *testing.T) {

	goodStruct := &RegisterRequest{
		UserName: "goods",
	}

	scheme := CreateValidationScheme(false, Scheme{
		"UserName": *GetValidation().Len(5, "len == 5"),
	})

	errors, valid := scheme.Validate(goodStruct)

	if !valid {
		t.Fatalf("Scheme is right, something went wrong")
	}

	if errors != nil {
		t.Fatalf("errors val must be empty, something went wrong")
	}

	badStruct := &RegisterRequest{
		UserName: "bad",
	}

	errors, valid = scheme.Validate(badStruct)

	if valid {
		t.Fatalf("Scheme is bad, but here is valid = true, something went wrong")
	}

	if errors == nil {
		t.Fatalf("errors val must be empty, something went wrong")
	}

	if errors["UserName"] == nil || len(errors["UserName"]) == 0 {
		t.Fatalf("Error structure is incorrect %v", errors)
	}

	if errors["UserName"][0] != "len == 5" {
		t.Fatalf("Message on an error is incorrect `%s`, must be `len == 5`", errors["UserName"][0])
	}
}

func TestValidationNotEmpty(t *testing.T) {

	goodStruct := &RegisterRequest{
		UserName: "good",
	}

	scheme := CreateValidationScheme(false, Scheme{
		"UserName": *GetValidation().NotEmpty("cant be empty"),
	})

	errors, valid := scheme.Validate(goodStruct)

	if !valid {
		t.Fatalf("Scheme is right, something went wrong")
	}

	if errors != nil {
		t.Fatalf("errors val must be empty, something went wrong")
	}

	badStruct := &RegisterRequest{
		UserName: "",
	}

	errors, valid = scheme.Validate(badStruct)

	if valid {
		t.Fatalf("Scheme is bad, but here is valid = true, something went wrong")
	}

	if errors == nil {
		t.Fatalf("errors val must be empty, something went wrong")
	}

	if errors["UserName"] == nil || len(errors["UserName"]) == 0 {
		t.Fatalf("Error structure is incorrect %v", errors)
	}

	if errors["UserName"][0] != "cant be empty" {
		t.Fatalf("Message on an error is incorrect `%s`, must be `cant be empty`", errors["UserName"][0])
	}
}
func TestValidationStructCustom(t *testing.T) {

	goodStruct := &RegisterRequest{
		UserName: "good",
	}

	customValidator := func(fieldType string, fieldValue interface{}) (string, bool, bool) {
		return "cant be `bad`", fieldValue != "bad", false
	}

	scheme := CreateValidationScheme(false, Scheme{
		"UserName": *GetValidation().CustomValidator(&customValidator),
	})

	errors, valid := scheme.Validate(goodStruct)

	if !valid {
		t.Fatalf("Scheme is right, something went wrong")
	}

	if errors != nil {
		t.Fatalf("errors val must be empty, something went wrong")
	}

	badStruct := &RegisterRequest{
		UserName: "bad",
	}

	scheme = CreateValidationScheme(false, Scheme{
		"UserName": *GetValidation().CustomValidator(&customValidator),
	})

	errors, valid = scheme.Validate(badStruct)

	if valid {
		t.Fatalf("Scheme is bad, but here is valid = true, something went wrong")
	}

	if errors == nil {
		t.Fatalf("errors val must be empty, something went wrong")
	}

	if errors["UserName"] == nil || len(errors["UserName"]) == 0 {
		t.Fatalf("Error structure is incorrect %v", errors)
	}

	if errors["UserName"][0] != "cant be `bad`" {
		t.Fatalf("Message on an error is incorrect `%s`, must be `cant be `bad``", errors["UserName"][0])
	}
}

func TestValidationValidateVarCustom(t *testing.T) {

	userName := "good"

	customValidator := func(fieldType string, fieldValue interface{}) (string, bool, bool) {
		return "cant be `bad`", fieldValue != "bad", false
	}

	errors, valid := ValidateVar(userName, *GetValidation().CustomValidator(&customValidator))

	if !valid {
		t.Fatalf("Scheme is right, something went wrong")
	}

	if errors != nil {
		t.Fatalf("errors val must be empty, something went wrong")
	}

	userName = "bad"

	errors, valid = ValidateVar(userName, *GetValidation().CustomValidator(&customValidator))

	if valid {
		t.Fatalf("Scheme is bad, but here is valid = true, something went wrong")
	}

	if errors == nil {
		t.Fatalf("errors val must be empty, something went wrong")
	}

	if errors["val"] == nil || len(errors["val"]) == 0 {
		t.Fatalf("Error structure is incorrect %v", errors)
	}

	if errors["val"][0] != "cant be `bad`" {
		t.Fatalf("Message on an error is incorrect `%s`, must be `cant be `bad``", errors["val"][0])
	}
}

func TestValidationValidateVar(t *testing.T) {

	userName := "good"

	errors, valid := ValidateVar(userName, *GetValidation().NotEmpty("cant be empty"))

	if !valid {
		t.Fatalf("Scheme is right, something went wrong")
	}

	if errors != nil {
		t.Fatalf("errors val must be empty, something went wrong")
	}

	userName = ""

	errors, valid = ValidateVar(userName, *GetValidation().NotEmpty("cant be empty"))

	if valid {
		t.Fatalf("Scheme is bad, but here is valid = true, something went wrong")
	}

	if errors == nil {
		t.Fatalf("errors val must be empty, something went wrong")
	}

	if errors["val"] == nil || len(errors["val"]) == 0 {
		t.Fatalf("Error structure is incorrect %v", errors)
	}

	if errors["val"][0] != "cant be empty" {
		t.Fatalf("Message on an error is incorrect `%s`, must be `cant be empty`", errors["val"][0])
	}
}
