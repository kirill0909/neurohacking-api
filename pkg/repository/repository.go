package repository

import ()

type User interface{}

type Category interface{}

type Word interface{}

type Repository struct {
	User
	Category
	Word
}

func NewRepository() *Repository {
	return &Repository{}
}
