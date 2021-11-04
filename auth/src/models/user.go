package models

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=7,max=30"`
}

func (user *User) ValidateUserModel() bool {
	//Add regexp later
	return true
}
