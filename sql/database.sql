CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
  id int auto_increment primary key,
  name varchar(50) not null ,
  nick varchar(50) unique not null,
  email varchar(50) unique not null,
  password varchar(50) not null,
  createdAt timestamp default current_timestamp 

)

