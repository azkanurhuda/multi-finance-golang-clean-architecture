drop table users;

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

drop table merchants;

create table merchants
(
    id varchar(100) not null,
    name varchar(100) not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    primary key (id)
) engine = InnoDB;

drop table credit_limits;

create table credit_limits
(
    id varchar(100) not null,
    user_id varchar(100) not null,
    tenor int not null,
    credit_limit decimal(10,2) not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    primary key (id),
    foreign key (user_id) references users(id)
) engine = InnoDB;

drop table consumers;

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

drop table assets;

create table assets
(
    id varchar(100) not null,
    merchant_id varchar(100) not null,
    name varchar(255) not null,
    otr decimal(10,2) not null,
    admin_fee decimal(10,2) not null,
    description varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    primary key (id),
    foreign key (merchant_id) references merchants(id)
) engine = InnoDB;

drop table transactions;

create table transactions
(
    id         varchar(100) not null,
    user_id       varchar(100) not null,
    asset_id       varchar(100) not null,
    contract_number       varchar(255) not null,
    total_payment       decimal(10,2) not null,
    payment_method       varchar(30) not null,
    created_at timestamp       not null,
    updated_at timestamp       not null,
    credit_id varchar(100) null,
    primary key (id),
    foreign key (user_id) references users(id),
    foreign key (asset_id) references assets(id)
) engine = InnoDB;

drop table credits;

create table credits
(
    id         varchar(100) not null,
    transaction_id         varchar(100) not null,
    tenor       int not null,
    credit_limit       decimal(10,2) not null,
    monthly_installment       decimal(10,2) not null,
    interest_amount       decimal(10,2) not null,
    created_at timestamp       not null,
    updated_at timestamp       not null,
    primary key (id),
    foreign key (transaction_id) references transactions(id)
) engine = InnoDB;

drop table credit_payments;

create table credit_payments
(
    id         varchar(100) not null,
    credit_id       varchar(100) not null,
    payment_amount       decimal(10,2) not null,
    created_at timestamp       not null,
    updated_at timestamp       not null,
    primary key (id),
    foreign key (credit_id) references credits(id)
) engine = InnoDB;

delete from merchants;

INSERT INTO merchants (id, name, created_at, updated_at)
VALUES ('e548aa9f-86ff-49ba-900f-9d4834d90cd1', 'PT Ahass Maju', '2024-05-19 21:56:23.000000', '2024-05-19 21:56:23.000000'),
       ('da605750-a29b-4e3d-9f4e-8154e1ab883a', 'Tokopedia', '2024-05-19 21:56:23.000000', '2024-05-19 21:56:23.000000'),
       ('4cb3afb7-e551-4b99-9c26-a7d70c439e83', 'Blibli', '2024-05-19 21:56:23.000000', '2024-05-19 21:56:23.000000');

DELETE FROM assets;

INSERT INTO assets (id, merchant_id, name, otr, admin_fee, description, created_at, updated_at)
VALUES ('70dc95b1-a873-4294-a2a0-d13eedef015e', 'e548aa9f-86ff-49ba-900f-9d4834d90cd1', 'Motor CBR 250 CBS', 30000000, 1200000, 'Motor Cepat', '2024-05-19 21:56:23.000000', '2024-05-19 21:56:23.000000'),
       ('7247735f-5436-4c3a-800b-605aa81fd375', 'e548aa9f-86ff-49ba-900f-9d4834d90cd1', 'Motor CBR 250 ABS', 32000000, 1400000, 'Motor Cepat', '2024-05-19 21:56:23.000000', '2024-05-19 21:56:23.000000'),
       ('d3e44ae0-3ff1-40e9-a988-d9aaddcd1958', 'e548aa9f-86ff-49ba-900f-9d4834d90cd1', 'Motor Honda Beat 250 ABS', 18000000, 720000, 'Motor Kota', '2024-05-19 21:56:23.000000', '2024-05-19 21:56:23.000000'),
       ('d25bea00-236a-45d0-866b-d50c464b9fdd', 'da605750-a29b-4e3d-9f4e-8154e1ab883a', 'Laptop Macbook M1 Pro 256', 20000000, 800000, 'Apple Macbook M1 Laptop', '2024-05-19 21:56:23.000000', '2024-05-19 21:56:23.000000'),
       ('7e345910-de2b-4dfb-9dc9-d1675c533c53', 'da605750-a29b-4e3d-9f4e-8154e1ab883a', 'Laptop Macbook M2 Pro 256', 21000000, 800000, 'Apple Macbook M1 Laptop', '2024-05-19 21:56:23.000000', '2024-05-19 21:56:23.000000'),
       ('6ac0adf4-ffa0-4eb6-b36e-07e1351c2736', '4cb3afb7-e551-4b99-9c26-a7d70c439e83', 'Samsung Galaxy S24 Ultra', 24000000, 1000000, 'Samsung Punya Minat', '2024-05-19 21:56:23.000000', '2024-05-19 21:56:23.000000');

DELETE FROM users;

INSERT INTO users (id, email, password, role, created_at, updated_at)
VALUES ('61a928dd-105b-46d3-b7f8-7558f75ca20e', 'nurhudaazka@gmail.com', '$2a$10$o.qvSBrnA7zlf6ewVMj7celzKBcRDeP.4o7eiTMCiEd5aPRyODPZS', 'admin', '2024-05-19 21:56:23.000000', '2024-05-19 21:56:23.000000');

DELETE FROM consumers;

INSERT INTO consumers (id, user_id, nik, full_name, legal_name, phone_number, address, place_of_birth, date_of_birth, salary, created_at, updated_at)
VALUES ('3a9a2e40-0b74-46ba-81e5-ccad458d8874', '61a928dd-105b-46d3-b7f8-7558f75ca20e', '8492938990111112', 'Azka Nurhuda', 'Azka Nurhuda', '08143924001', 'Sleman, Yogyakarta', 'Sleman', '2000-05-19', 20000000, '2024-05-19 21:56:23.000000', '2024-05-19 21:56:23.000000');