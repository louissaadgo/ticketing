package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

func (user *User) ValidateUserModel() bool {
	//Add regexp later
	return true
}

func (user *User) HashPassword() bool {
	sb, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(sb)
	return err == nil
}

func (user *User) CompareHashAndPassword(hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(user.Password))
	return err == nil
}
