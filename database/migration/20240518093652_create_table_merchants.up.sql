create table merchants
(
    id varchar(100) not null,
    name varchar(100) not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    primary key (id)
) engine = InnoDB;