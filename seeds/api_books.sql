create table if not exists "books" (
  id serial primary key,
  titulo text not null unique,
  categoria text not null,
  autor text not null,
  sinopse text not null,
  created_at timestamp not null default now(),
  updated_at timestamp not null default now(),
  deleted_at timestamp
);

insert into "books" ("autor", "categoria", "created_at", "deleted_at", "id", "sinopse", "titulo", "updated_at") values ('julia naomi boeira', 'rust', '2022-09-05 00:01:23.891504+00', NULL, '2', 'rust é uma linguagem de nível de sistema com algumas características que são grandes diferenciais, como ter segurança de memória sem coletor de lixo, possibilitar concorrência sem corrida de dados, abstração sem overhead e um compilador que garante segurança no alocamento de recursos. com rust, pensar em programação funcional será algo muito intuitivo no desenvolvimento.', 'programação funcional e concorrente em rust', '2022-09-05 00:01:23.891504+00');
insert into "books" ("autor", "categoria", "created_at", "deleted_at", "id", "sinopse", "titulo", "updated_at") values ('marcelo castellani e willian molinari', 'rust', '2022-09-04 23:59:02.619135+00', NULL, '1', 'concorrência e alta performance com segurança', 'rust', '2022-09-05 00:02:07.157129+00');
insert into "books" ("autor", "categoria", "created_at", "deleted_at", "id", "sinopse", "titulo", "updated_at") values ('thiago leite e carvalho', 'poo', '2022-09-05 00:03:05.502064+00', NULL, '3', 'raprenda seus conceitos e suas aplicabilidades de forma efetiva', 'orientação a objetos', '2022-09-05 00:03:05.502064+00');
insert into "books" ("autor", "categoria", "created_at", "deleted_at", "id", "sinopse", "titulo", "updated_at") values ('joão kuntz', 'blockchain', '2022-09-05 00:03:47.597863+00', NULL, '4', 'fundamentos de arquitetura, desenvolvimento de contratos e aplicações', 'blockchain ethereum', '2022-09-05 00:03:47.597863+00');
insert into "books" ("autor", "categoria", "created_at", "deleted_at", "id", "sinopse", "titulo", "updated_at") values ('juliano ramos', 'linux', '2022-09-05 00:04:34.952617+00', NULL, '5', 'guia prático para a prova lpic-1 101', 'certificação linux', '2022-09-05 00:04:34.952617+00');
insert into "books" ("autor", "categoria", "created_at", "deleted_at", "id", "sinopse", "titulo", "updated_at") values ('victor osório', 'back-end', '2022-09-05 00:05:22.326462+00', NULL, '6', 'conhecendo o protocolo http e arquiteturas rest', 'roadmap back-end', '2022-09-05 00:05:22.326462+00');
insert into "books" ("autor", "categoria", "created_at", "deleted_at", "id", "sinopse", "titulo", "updated_at") values ('julio bitencourt', 'dart', '2022-09-05 00:06:01.91601+00', NULL, '7', 'fundamentos, prática, conceitos avançados e tudo mais', 'o guia de dart', '2022-09-05 00:06:01.91601+00');
insert into "books" ("autor", "categoria", "created_at", "deleted_at", "id", "sinopse", "titulo", "updated_at") values ('alexandre aquiles', 'poo', '2022-09-05 00:06:52.649041+00', NULL, '8', 'práticas avançadas para códigos de qualidade em java moderno', 'desbravando solid', '2022-09-05 00:06:52.649041+00');
insert into "books" ("autor", "categoria", "created_at", "deleted_at", "id", "sinopse", "titulo", "updated_at") values ('tiago davi', 'elixir', '2022-09-05 00:07:28.26566+00', NULL, '9', 'do zero à concorrência', 'elixir', '2022-09-05 00:07:28.26566+00');
insert into "books" ("autor", "categoria", "created_at", "deleted_at", "id", "sinopse", "titulo", "updated_at") values ('rafael sabbagh', 'scrum', '2022-09-05 00:08:05.230724+00', NULL, '10', 'gestão ágil para produtos de sucesso', 'scrum', '2022-09-05 00:08:05.230724+00');
