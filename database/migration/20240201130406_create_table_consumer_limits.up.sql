create table consumer_limits
(
    id         varchar(100) not null,
    nik       varchar(100) not null,
    tenor       varchar(255) not null,
    limits   bigint not null,
    created_at timestamp       not null,
    updated_at timestamp       not null,
    primary key (id)
) engine = InnoDB;