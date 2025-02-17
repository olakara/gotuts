package domain

import (
	"testing"
)

func TestNewEmployee_WhenAllDataIsProvided_CreateEmployee(t *testing.T) {
	// Create a new employee
	employee, err := NewEmployee("John", "Doe")
	if err != nil {
		t.Fatal(err)
	}
	if employee != nil {
		t.Logf("Employee created: %v", employee)
	}
}

func TestNewEmployee_WhenFirstNameIsEmpty_ReturnError(t *testing.T) {
	// Create a new employee
	employee, err := NewEmployee("", "Doe")
	if err != nil {
		t.Logf("Error: %v", err)
	}
	if employee == nil {
		t.Logf("Employee is nil")
	}
}

func TestNewEmployee_WhenLastNameIsEmpty_ReturnError(t *testing.T) {
	// Create a new employee
	employee, err := NewEmployee("John", "")
	if err != nil {
		t.Logf("Error: %v", err)
	}
	if employee == nil {
		t.Logf("Employee is nil")
	}
}

func TestEmployee_GetFullName_ReturnValidFullName(t *testing.T) {
	// Create a new employee
	employee, err := NewEmployee("John", "Doe")
	if err != nil {
		t.Fatal(err)
	}

	fullName := employee.GetFullName()
	//assert the value of fullName
	if fullName != "John Doe" {
		t.Fatalf("Expected full name is John Doe, but got %s", fullName)
	}
}
