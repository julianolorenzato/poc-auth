package domain

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Permissions uint8

const (
	Read Permissions = iota
	Write
	ReadWrite
)

func (p Permissions) String() string {
	fmt.Println()
	return [...]string{"Read", "Write", "ReadWrite"}[p]
}

type User struct {
	ID          string
	Username    string
	Email       string
	password    string
	Permissions Permissions
}

func New(username, email, password string, permissions Permissions) (*User, error) {
	user := &User{
		ID:          uuid.NewString(),
		Username:    username,
		Email:       email,
		password:    password,
		Permissions: permissions,
	}

	err := user.hashPassword()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) hashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.password), 10)
	if err != nil {
		return err
	}

	u.password = string(hash)

	return nil
}

func (u *User) ComparePasswords(pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(pass))
	return err == nil
}
