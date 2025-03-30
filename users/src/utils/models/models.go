package models

type User struct {
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	PhoneNo  string `json:"phone_no" validate:"required,e164"`
	UserType string `json:"user_type" validate:"required,oneof=manager guest"`
}
