create table users
(
    id         varchar(100) not null,
    nik       varchar(100) not null,
    password   varchar(100) not null,
    token      varchar(100) null,
    created_at timestamp       not null,
    updated_at timestamp       not null,
    primary key (id)
) engine = InnoDB;