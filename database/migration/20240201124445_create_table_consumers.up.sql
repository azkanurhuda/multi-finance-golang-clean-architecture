create table consumers
(
    id         varchar(100) not null,
    nik       varchar(100) not null unique,
    full_name   varchar(255) not null,
    legal_name   varchar(255) not null,
    place_of_birth   varchar(255) not null,
    date_of_birth   timestamp not null,
    salary bigint not null,
    id_card_photo LONGBLOB not null,
    selfie_photo LONGBLOB not null,
    created_at timestamp       not null,
    updated_at timestamp       not null,
    primary key (id)
) engine = InnoDB;