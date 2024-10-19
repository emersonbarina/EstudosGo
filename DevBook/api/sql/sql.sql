CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

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

