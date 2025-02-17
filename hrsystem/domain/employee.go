package domain

import "errors"

type Employee struct {
	firstName string
	lastName  string
}

func NewEmployee(firstName, lastName string) (*Employee, error) {
	if firstName == "" {
		return nil, errors.New("first name is required")
	}
	if lastName == "" {
		return nil, errors.New("last name is required")
	}
	return &Employee{
		firstName: firstName,
		lastName:  lastName,
	}, nil
}

func (e *Employee) GetFullName() string {
	return e.firstName + " " + e.lastName
}
