# EstudosGo


# basic mysql / docker

docker pull mysql:9.0.1  
docker run --name devbook -e MYSQL_ROOT_PASSWORD=root -d mysql 
docker ps    
docker exec -it devbook bash  

https://www.datacamp.com/pt/tutorial/set-up-and-configure-mysql-in-docker
https://hub.docker.com/_/mysql


mysql -u goolang -p
show databases;
CREATE DATABASE devbook;
USE devbook;

CREATE TABLE usuarios( id int auto_increment primary key, nome varchar(50) not null, email varchar(50) not null ) ENGINE=INNODB;
show tables;
GRANT ALL PRIVILEGES ON devbook.* TO 'golang'@'localhost'

exit

go mod init banco-de-dados
go get github.com/go-sql-driver/mysql
