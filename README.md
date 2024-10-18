# EstudosGo


# basic mysql / docker

docker pull mysql:9.0.1  
docker run --name devbook -e MYSQL_ROOT_PASSWORD=root -d mysql 
docker ps    
docker exec -it devbook bash  

https://www.datacamp.com/pt/tutorial/set-up-and-configure-mysql-in-docker
https://hub.docker.com/_/mysql

mysql -u root -p

show databases;
CREATE DATABASE devbook;
USE devbook;

CREATE TABLE usuarios( id int auto_increment primary key, nome varchar(50) not null, email varchar(50) not null ) ENGINE=INNODB;
show tables;
CREATE USER 'golang'@'localhost'IDENTIFIED BY 'golang'
GRANT ALL PRIVILEGES ON devbook.* TO 'golang'@'localhost';

desc usuarios;

mysql -u goolang -p

exit

Parei o serviço e criei novamente informando a porta.
docker stop devbook
docker rm devbook
docker run -d --name devbook -e MYSQL_ROOT_PASSWORD=root -p 3307:3306 mysql  
docker port devbook

>>> resultado : 3306/tcp -> 0.0.0.0:3307

Dicas: 
https://www.connectionstrings.com/
https://www.youtube.com/watch?v=1VzhUEn2pCU&ab_channel=BrianMorrison


# usando no go - CRUD Básico
go mod init <projeto>
go get github.com/go-sql-driver/mysql

go get github.com/gorilla/mux

go build
./crud

# início do projeto com API
Optei por manter junto ao projeto de estudos.

go mod init api

go get github.com/gorilla/mux

go get github.com/joho/godotenv 

go get github.com/go-sql-driver/mysql

go get github.com/badoux/checkmail


