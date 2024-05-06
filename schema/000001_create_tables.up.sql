CREATE TABLE  IF NOT EXISTS People
(
    id   SERIAL primary key,
    name    varchar(30) NOT NULL ,
    surname varchar(30) NOT NULL,
    patronymic varchar(50)
);

CREATE TABLE  IF NOT EXISTS Car
(
    regNum   varchar(25) PRIMARY KEY,
    mark     varchar(50) NOT NULL ,
    model    varchar(50) NOT NULL ,
    year     smallint,
    owner    SERIAL,
    foreign key (owner)   references  People(id) on delete cascade
);