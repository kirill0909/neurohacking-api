package models

type User struct {
	Id           int    `json:"-"`
	Name         string `json:"name"     binding:"required"`
	Email        string `json:"email"    binding:"required"`
	Password     string `json:"password" binding:"required"`
	DateCreation string
	LastUpdate   string
}

type UserSignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
