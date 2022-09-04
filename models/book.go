package models

import (
	"errors"
	"strings"
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

func (b *Book) hasInvalidFields() error {
	if b.Titulo == "" {
		return errors.New("Título é um campo obrigatório")
	}
	if b.Categoria == "" {
		return errors.New("Categoria é um campo obrigatório")
	}
	if b.Autor == "" {
		return errors.New("Autor é um campo obrigatório")
	}
	if b.Sinopse == "" {
		return errors.New("Sinopse é um campo obrigatório")
	}
	return nil
}

func (b *Book) formatter() {
	b.Titulo = strings.TrimSpace(strings.ToLower(b.Titulo))
	b.Categoria = strings.TrimSpace(strings.ToLower(b.Categoria))
	b.Autor = strings.TrimSpace(strings.ToLower(b.Autor))
	b.Sinopse = strings.TrimSpace(strings.ToLower(b.Sinopse))
}

func (b *Book) Validate() error {
	if err := b.hasInvalidFields(); err != nil {
		return err
	}
	b.formatter()
	return nil
}
