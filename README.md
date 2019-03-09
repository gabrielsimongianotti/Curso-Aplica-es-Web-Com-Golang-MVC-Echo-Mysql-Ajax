# Curso-Aplica-es-Web-Com-Golang-MVC-Echo-Mysql-Ajax
Atividades do curso da Aplicações Web Com Golang MVC + Echo + Mysql + Ajax

link: https://www.udemy.com/aplicacoes-web-com-golang/

comando: dep ensure

base de dados:
    create database membros;

create table usuarios (
    id integer not null auto_increment,
    nome varchar(50) null,
    email varchar(255) not null,
    primary key(id)
);