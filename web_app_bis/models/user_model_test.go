package models

import (
	"testing"
)

func TestUser(t *testing.T) {
	// Create a sample User
	user := User{
		Name: "John Doe",
		Age:  30,
		Address: Address{
			State:   "California",
			City:    "Los Angeles",
			Pincode: 90001,
		},
	}

	// Test the fields of the User struct
	if user.Name != "John Doe" {
		t.Errorf("Expected Name to be 'John Doe', but got %s", user.Name)
	}

	if user.Age != 30 {
		t.Errorf("Expected Age to be 30, but got %d", user.Age)
	}

	// Test the fields of the Address struct within User
	if user.Address.State != "California" {
		t.Errorf("Expected Address.State to be 'California', but got %s", user.Address.State)
	}

	if user.Address.City != "Los Angeles" {
		t.Errorf("Expected Address.City to be 'Los Angeles', but got %s", user.Address.City)
	}

	if user.Address.Pincode != 90001 {
		t.Errorf("Expected Address.Pincode to be 90001, but got %d", user.Address.Pincode)
	}
}
