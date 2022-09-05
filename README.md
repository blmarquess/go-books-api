# go-books-api

## Problema a ser resolvido com este projeto

Esperamos que você desenvolva uma API Rest que contemple os livros de um usuário.

---

Requisitos:

- Como usuário eu gostaria de cadastrar um livro.
- Como usuário eu gostaria de atualizar um livro.
- Como usuário eu gostaria de listar um livro.
- Como usuário eu gostaria de listar todos os livros.
- Como usuário eu gostaria de deletar um livro.

Os livros devem ter:

    - Título
    - Categoria.
    - Autor.
    - Sinopse.

## O que foi utilizado no projeto

- Linguagem Go
- GORM: Para evitar queries manuais no código e fazer a sanitização das mesmas
- Gin Web Framework: Por conta do ganho de performance tanto no momento de escrever o código quanto no resultado pos build
- PostgresSQL: Por ser um banco de dados Open Source e flexível para usar num MVP e possibilitar um fácil escalonamento posterior

## Para testar a aplicação

Pre-requisitos:
Docker e uma ferramenta de requisições http exemplo curl, httpie, postman ou insomnia

clone o projeto para sua maquina:

```shell
git clone https://github.com/blmarquess/go-books-api.git
```

Entre na pasta do projeto e executo o docker compose:

```shell
cd go-books-api
docker compose up --build
```

O projeto esta configurado para iniciar na porta 8000 da maquina hospedeira

Exemplo no terminal com curl para ver a listagem dos livros:

```shell
curl --location --request GET 'http://localhost:8000/books' \
--data-raw ''
```

A [documentação](https://documenter.getpostman.com/view/20082367/VUxXJhqA) completa da aplicação com exemplos pronto para usar no postman estão nessa pagina:
<https://documenter.getpostman.com/view/20082367/VUxXJhqA>

![print da documentação](public/Captura%20de%20Tela%202022-09-04%20%C3%A0s%2020.10.31.png)
