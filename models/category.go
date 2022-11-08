package models

type Category struct {
	Id           int
	UID          int
	Name         string `json:"name" binding:"required"`
	DateCreation string
	LastUpdate   string
}
