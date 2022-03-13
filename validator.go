package validator

import (
	"reflect"
)

type Scheme map[string]Validation

type Validation struct {
	Validators []*func(string, interface{}) (string, bool, bool)
}

type ValidationScheme struct {
	validation                map[string]Validation
	forbidFieldsIfNotInScheme bool
}

// Create new scheme validator
// forbidFieldsIfNotInScheme - if true, then validator will throw
// an error if scheme doesn't has some field from struct
func CreateValidationScheme(forbidFieldsIfNotInScheme bool, scheme map[string]Validation) *ValidationScheme {
	return &ValidationScheme{
		validation:                scheme,
		forbidFieldsIfNotInScheme: forbidFieldsIfNotInScheme,
	}
}

// Create new validation ruler
func GetValidation() *Validation {
	return &Validation{}
}

// Returns true if field is string
func (v *Validation) String(message string) *Validation {
	checkFunction := func(fieldType string, fieldValue interface{}) (string, bool, bool) {
		return message, fieldType == "string", true
	}

	v.Validators = append(v.Validators, &checkFunction)
	return v
}

// Returns true if string length is lower or equal provided length
func (v *Validation) LenLte(length int, message string) *Validation {
	checkFunction := func(fieldType string, fieldValue interface{}) (string, bool, bool) {
		return message, fieldType == "string" && len(fieldValue.(string)) <= length, false
	}

	v.Validators = append(v.Validators, &checkFunction)
	return v
}

// Returns true if string length is greeter or equal provided length
func (v *Validation) LenGte(length int, message string) *Validation {
	checkFunction := func(fieldType string, fieldValue interface{}) (string, bool, bool) {
		return message, fieldType == "string" && len(fieldValue.(string)) >= length, false
	}

	v.Validators = append(v.Validators, &checkFunction)
	return v
}

// Returns true if string length is equal @param length
func (v *Validation) Len(length int, message string) *Validation {
	checkFunction := func(fieldType string, fieldValue interface{}) (string, bool, bool) {
		return message, fieldType == "string" && len(fieldValue.(string)) == length, false
	}

	v.Validators = append(v.Validators, &checkFunction)
	return v
}

// Returns true if string is not empty
func (v *Validation) NotEmpty(message string) *Validation {
	checkFunction := func(fieldType string, fieldValue interface{}) (string, bool, bool) {
		return message, fieldType == "string" && len(fieldValue.(string)) > 0, false
	}

	v.Validators = append(v.Validators, &checkFunction)
	return v
}

// Add custom validation function.
// validateFunc must return 3 parameters (string (message of an error), bool (isValid),
// bool (immediateStop (if true then validation will stop checking other validation functions)))
func (v *Validation) CustomValidator(validateFunc *func(string, interface{}) (string, bool, bool)) *Validation {
	v.Validators = append(v.Validators, validateFunc)
	return v
}

func addError(fieldName, msg string, errors map[string][]string) map[string][]string {
	if errors == nil {
		errors = make(map[string][]string)
	}

	errors[fieldName] = append(errors[fieldName], msg)

	return errors
}

func ValidateVar(val interface{}, scheme Validation) (map[string][]string, bool) {
	var errors map[string][]string

	structVal := reflect.TypeOf(val)

	fieldType := structVal.String()
	fieldValue := val

	for _, validator := range scheme.Validators {
		msg, valid, immediateStop := (*validator)(fieldType, fieldValue)

		if !valid {
			errors = addError("val", msg, errors)

			if immediateStop {
				break
			}
		}

	}

	return errors, errors == nil
}

// Validate scheme with given rules
// First return value - map of field with an array of error messages | nil
// Seconds return value - bool, true if no errors
func (vs *ValidationScheme) Validate(structure interface{}) (map[string][]string, bool) {
	var errors map[string][]string

	structVal := reflect.ValueOf(structure)

	if structVal.Kind() != reflect.Ptr {
		panic("structure parameter must be a pointer to struct (*)")
	}

	e := structVal.Elem()

	for i := 0; i < e.NumField(); i++ {
		field := e.Type().Field(i)
		fieldName := field.Name
		fieldType := field.Type.String()
		fieldValue := e.Field(i).Interface()

		validators, fieldExists := vs.validation[fieldName]

		if !fieldExists {
			if vs.forbidFieldsIfNotInScheme {
				errors = addError(fieldName, "not allowed in this struct", errors)
			}
			continue
		}

		for _, validator := range validators.Validators {
			msg, valid, immediateStop := (*validator)(fieldType, fieldValue)

			if !valid {
				errors = addError(fieldName, msg, errors)

				if immediateStop {
					break
				}
			}
		}
	}

	return errors, errors == nil
}
