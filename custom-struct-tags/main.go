package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

const (
	validateTag string = "validate"
	minKey      string = "min"
	minAgeValue int    = 18
	requiredKey        = "required"
)

type User struct {
	Email string `json:"email" validate:"required"`
	Age   int    `json:"age" validate:"required,min=18"`
}

func isValid(input interface{}) (bool, error) {
	val := reflect.ValueOf(input)
	typ := reflect.TypeOf(input)

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i)
		validateTag := field.Tag.Get(validateTag)

		if strings.Contains(validateTag, requiredKey) && fieldValue.IsZero() {
			return false, fmt.Errorf("field %s is required", field.Name)
		}

		if strings.Contains(validateTag, fmt.Sprintf("%s=", minKey)) {
			minAge := extractAgeVal(validateTag)
			if fieldValue.Int() < int64(minAge) {
				return false, fmt.Errorf("field %s should be at least %d", field.Name, minAge)
			}
		}
	}

	return true, nil
}

func extractAgeVal(tag string) int {
	parts := strings.Split(tag, ",")
	for _, part := range parts {
		if strings.HasPrefix(part, fmt.Sprintf("%s=", minKey)) {
			var minAge int
			fmt.Sscanf(part, "min=%d", &minAge)
			return minAge
		}
	}
	return 0
}

func main() {
	user := User{Email: "xyz@abc.co", Age: 90}
	log.Printf("%+v", user)

	_, err := isValid(user)

	if err != nil {
		log.Fatal(err.Error())
	}
}
