package models

import (
	"errors"
	"fmt"
)

// User is the container for user data
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers returns all users
func GetUsers() []*User {
	return users
}

// InitUsers initialize the user model state
func InitUsers(theUsers []*User) {
	if len(theUsers) > 0 {
		users = theUsers
		nextID = theUsers[len(theUsers)-1].ID + 1
	}
}

// AddUser adds a user to the list of users
func AddUser(u User) (User, error) {
	fmt.Println("models.AddUser")
	if u.ID != 0 {
		return User{}, errors.New("New User must not include id or it must be set to zero")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// GetUserByID returns the user with the provided id
func GetUserByID(id int) (User, error) {
	fmt.Println("models.GetUserByID")
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

// UpdateUser updates the user wtih the provided user id
func UpdateUser(u User) (User, error) {
	fmt.Println("models.UpdateUser")
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

// RemoveUserByID removes the user with the provided user id from the list of users
func RemoveUserByID(id int) error {
	fmt.Println("models.RemoveUserById")
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID '%v' not found", id)
}
