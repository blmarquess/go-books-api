package seeds

import (
	"log"

	"github.com/blmarquess/go-books-api/models"
	"gorm.io/gorm"
)

func Books(db *gorm.DB) {
	books := []models.Book{
		{
			Titulo:    "programação funcional e concorrente em rust",
			Categoria: "rust",
			Autor:     "julia naomi boeira",
			Sinopse:   "rust é uma linguagem de nível de sistema com algumas características que são grandes diferenciais, como ter segurança de memória sem coletor de lixo, possibilitar concorrência sem corrida de dados, abstração sem overhead e um compilador que garante segurança no alocamento de recursos. com rust, pensar em programação funcional será algo muito intuitivo no desenvolvimento.",
		},
		{
			Titulo:    "rust",
			Categoria: "rust",
			Autor:     "marcelo castellani e willian molinari",
			Sinopse:   "concorrência e alta performance com segurança",
		},
		{
			Titulo:    "orientação a objetos",
			Categoria: "poo",
			Autor:     "thiago leite e carvalho",
			Sinopse:   "raprenda seus conceitos e suas aplicabilidades de forma efetiva",
		},
		{
			Titulo:    "blockchain ethereum",
			Categoria: "blockchain",
			Autor:     "joão kuntz",
			Sinopse:   "fundamentos de arquitetura, desenvolvimento de contratos e aplicações",
		},
		{
			Titulo:    "certificação linux",
			Categoria: "linux",
			Autor:     "juliano ramos",
			Sinopse:   "guia prático para a prova lpic-1 101",
		},
		{
			Titulo:    "roadmap back-end",
			Categoria: "back-end",
			Autor:     "victor osório",
			Sinopse:   "conhecendo o protocolo http e arquiteturas rest",
		},
		{
			Titulo:    "o guia de dart",
			Categoria: "dart",
			Autor:     "julio bitencourt",
			Sinopse:   "fundamentos, prática, conceitos avançados e tudo mais",
		},
		{
			Titulo:    "desbravando solid",
			Categoria: "poo",
			Autor:     "alexandre aquiles",
			Sinopse:   "práticas avançadas para códigos de qualidade em java moderno",
		},
		{
			Titulo:    "elixir",
			Categoria: "elixir",
			Autor:     "tiago davi",
			Sinopse:   "do zero à concorrência",
		},
		{
			Titulo:    "scrum",
			Categoria: "scrum",
			Autor:     "rafael sabbagh",
		},
	}

	for _, book := range books {
		if err := book.Validate(); err != nil {
			log.Fatal(err)
		}
		if err := db.Create(&book).Error; err != nil {
			log.Fatal("Error on seeding database" + err.Error())
		}
	}
}
