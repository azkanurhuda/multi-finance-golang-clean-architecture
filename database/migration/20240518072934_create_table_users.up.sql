create table users
(
    id         varchar(100) not null,
    email   varchar(100) not null,
    password   varchar(100) not null,
    role varchar(100) not null,
    token      varchar(255) null,
    token_expired_at timestamp null,
    created_at timestamp not null,
    updated_at timestamp not null,
    primary key (id)
) engine = InnoDB;