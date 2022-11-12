package models

type Category struct {
	Id           int
	UID          int    `db:"user_id"`
	Name         string `json:"name" binding:"required"`
	DateCreation string `db:"date_creation"`
	LastUpdate   string `db:"last_update"`
}
