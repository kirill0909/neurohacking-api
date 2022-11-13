package models

type Word struct {
	Id           int
	UID          int    `db:"user_id"`
	CategoryId   int    `db:"category_id"`
	Name         string `json:"name" binding:"required"`
	DateCreation string `db:"date_creation"`
	LastUpdate   string `db:"last_update"`
}

type WordUpdateInput struct {
	Name string `json:"name" binding:"required"`
}
