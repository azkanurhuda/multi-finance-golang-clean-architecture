create table consumers
(
    id         varchar(100) not null,
    user_id         varchar(100) not null,
    nik       varchar(100) not null unique,
    full_name   varchar(255) not null,
    legal_name   varchar(255) not null,
    phone_number   varchar(30) not null,
    address   varchar(255) not null,
    place_of_birth   varchar(100) not null,
    date_of_birth   timestamp not null,
    salary decimal(10,2) not null,
    id_card_photo LONGBLOB null,
    selfie_photo LONGBLOB null,
    created_at timestamp       not null,
    updated_at timestamp       not null,
    primary key (id),
    foreign key (user_id) references users(id)
) engine = InnoDB;