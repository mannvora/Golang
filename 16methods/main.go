package main

import (
	"fmt"
	"time"
)

// User struct with additional fields
type User struct {
	Name      string
	Email     string
	Age       int
	Status    bool
	CreatedAt time.Time
}

// Method to update the user's status
func (u *User) UpdateStatus(status bool) *User {
	u.Status = status
	fmt.Printf("Status updated for user %s to %t\n", u.Name, u.Status)
	return u // Returning the user allows method chaining
}

// Method to update email
func (u *User) UpdateEmail(newEmail string) *User {
	u.Email = newEmail
	fmt.Printf("Email updated for user %s to %s\n", u.Name, u.Email)
	return u
}

// Method to calculate years between user creation and current time
func (u User) GetAccountAge() int {
	years := time.Now().Year() - u.CreatedAt.Year()
	fmt.Printf("User %s has had an account for %d years\n", u.Name, years)
	return years
}

// Method to calculate age difference between two users
func (u User) AgeDifference(other User) int {
	diff := u.Age - other.Age
	if diff < 0 {
		diff = -diff
	}
	fmt.Printf("Age difference between %s and %s is %d years\n", u.Name, other.Name, diff)
	return diff
}

func main() {
	// Creating two users
	user1 := User{
		Name:      "Mann",
		Email:     "mannvora@gmail.com",
		Age:       23,
		Status:    true,
		CreatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	user2 := User{
		Name:      "Anil",
		Email:     "anilvora@example.com",
		Age:       28,
		Status:    false,
		CreatedAt: time.Date(2015, time.February, 15, 0, 0, 0, 0, time.UTC),
	}

	// Using method chaining to update status and email
	user1.UpdateStatus(false).UpdateEmail("newmann@example.com")

	// Getting account age for both users
	user1.GetAccountAge()
	user2.GetAccountAge()

	// Calculating the age difference between two users
	user1.AgeDifference(user2)

	// Printing user1 details in structured format
	fmt.Printf("%+v\n", user1)
}
