create table users
(
    id         varchar(100) not null,
    email   varchar(100) not null,
    password   varchar(100) not null,
    role varchar(100) not null,
    token      varchar(255) null,
    token_expired_at timestamp null,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    primary key (id)
) engine = InnoDB;