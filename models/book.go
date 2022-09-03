package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	Id        int            `json:"id" gorm:"primary_key"`
	Titulo    string         `json:"título"`
	Categoria string         `json:"categoria"`
	Autor     string         `json:"autor"`
	Sinopse   string         `json:"sinopse"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}
