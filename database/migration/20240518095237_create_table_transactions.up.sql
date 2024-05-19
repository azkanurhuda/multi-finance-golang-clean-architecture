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