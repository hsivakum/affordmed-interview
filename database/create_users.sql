create table users
(
    id         bigserial
        constraint users_pk
            primary key,
    first_name varchar(50),
    last_name  varchar(50),
    email      varchar(50) not null,
    password   char(64)    not null
);

create unique index users_email_uindex
    on users (email);

