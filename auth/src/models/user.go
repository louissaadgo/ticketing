package models

type User struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

func (user *User) ValidateUserModel() bool {
	//Add regexp later
	return true
}
