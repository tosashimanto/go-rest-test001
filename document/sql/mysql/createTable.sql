
-- MySQL

drop table IF EXISTS test001;


create table  IF NOT EXISTS test001 (
    id int AUTO_INCREMENT NOT NULL PRIMARY KEY,
    username varchar(255),
    email varchar(255),
    password char(30),
    createdBy int not null,
    createdAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedBy int not null,
    updatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);





