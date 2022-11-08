package models

type Word struct {
	Id           int
	UID          int
	CategoryId   int
	Name         string `json:"name" binding:"required"`
	DateCreation string
	LastUpdate   string
}
