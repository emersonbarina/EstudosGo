insert into usuarios (nome, nick, email, senha)
values
("Usuário 01", "user01", "usuario01@gmail.com", "$2a$10$0b8KoR8nZSrsk/XS92kC/u5KXMDOuxEXN1b0GgOTcJr1tyTYHGmzq"),
("Usuário 02", "user02", "usuario02@gmail.com", "$2a$10$0b8KoR8nZSrsk/XS92kC/u5KXMDOuxEXN1b0GgOTcJr1tyTYHGmzq"),
("Usuário 03", "user03", "usuario03@gmail.com", "$2a$10$0b8KoR8nZSrsk/XS92kC/u5KXMDOuxEXN1b0GgOTcJr1tyTYHGmzq"),
("Usuário 04", "user04", "usuario04@gmail.com", "$2a$10$0b8KoR8nZSrsk/XS92kC/u5KXMDOuxEXN1b0GgOTcJr1tyTYHGmzq"),
("Usuário 05", "user05", "usuario05@gmail.com", "$2a$10$0b8KoR8nZSrsk/XS92kC/u5KXMDOuxEXN1b0GgOTcJr1tyTYHGmzq"),
("Usuário 06", "user06", "usuario06@gmail.com", "$2a$10$0b8KoR8nZSrsk/XS92kC/u5KXMDOuxEXN1b0GgOTcJr1tyTYHGmzq");
commit;
insert into seguidores (usuario_id, seguidor_id)
values
(1,2),
(1,3),
(1,4),
(2,1),
(2,3),
(3,4),
(3,5),
(3,6);