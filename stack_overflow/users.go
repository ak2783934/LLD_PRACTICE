package stackoverflow

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	Id    string
	Name  string
	Email string
	Votes int
}

type UserManager struct {
	Users map[string]*User
}

func (u *UserManager) CreateUserManager() {
	u.Users = make(map[string]*User, 0)
}

func (u *UserManager) CreateUser(name string, email string) (string, error) {
	newUser := User{
		Id:    uuid.New().String(),
		Name:  name,
		Email: email,
		Votes: 0,
	}

	// check if this user id already exists
	_, exists := u.Users[newUser.Email]
	if !exists {
		u.Users[newUser.Email] = &newUser
	} else {
		return "", errors.New("user already exists for this email")
	}
	return newUser.Id, nil
}

func (u *UserManager) UpdateUser(email string, user *User) error {
	// check if the user exists.
	user, exists := u.Users[email]
	if !exists {
		return errors.New("user with this email doesn't exist")
	}
	u.Users[email] = user
	return nil
}

func (u *UserManager) UpdateVotes(email string, count int) error {
	// check if the user exists.
	user, exists := u.Users[email]
	if !exists {
		return errors.New("user with this email doesn't exist")
	}
	user.Votes += count
	return nil
}
