CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

COMMIT;

CREATE TABLE usuarios(
  id int primary key AUTO_INCREMENT,
  nome varchar(50) not null,
  nick varchar(50) not null unique,
  email varchar(50) not null unique,
  senha varchar(100) not null,
  criadoEm timestamp not null default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE seguidores(
  usuario_id int not null,
  FOREIGN KEY (usuario_id)
  REFERENCES usuarios(id)
  ON DELETE CASCADE,

  seguidor_id int not null,
  FOREIGN KEY (seguidor_id)
  REFERENCES usuarios(id)
  ON DELETE CASCADE,

  primary key(usuario_id, seguidor_id)
)