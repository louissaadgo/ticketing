package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

func (user *User) ValidateUserModel() bool {
	if len(user.Email) < 6 {
		return false
	}
	if len(user.Password) < 7 {
		return false
	}
	return true
}

func (user *User) HashPassword() bool {
	sb, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(sb)
	return err == nil
}

func (user *User) CompareHashAndPassword(hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(user.Password))
	return err == nil
}
