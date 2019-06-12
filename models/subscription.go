package models

import (
	"regexp"
)

var validPhoneRexExp = regexp.MustCompile(`^\+\s\d{2}\s\d{2}\s\d{6}$`)

type Subscription struct {
	PlanId *int    `json:"planId"`
	Name   *string `json:"name"`
	Email  *string `json:"email"`
	Phone  *string `json:"phone"`
}

func (s Subscription) Valid() (bool, map[string]string) {
	errorsList := make(map[string]string)

	if s.PlanId == nil {
		errorsList["planID"] = "is empty"
	}

	if s.Name == nil {
		errorsList["name"] = "is empty"
	}

	if s.Email == nil {
		errorsList["email"] = "is empty"
	}

	if s.Phone == nil {
		errorsList["phone"] = "is empty"
	}

	if s.Phone != nil && !validPhoneRexExp.MatchString(*s.Phone) {
		errorsList["phone"] = "is invalid"
	}

	if len(errorsList) > 0 {
		return false, errorsList
	}

	return true, nil
}
