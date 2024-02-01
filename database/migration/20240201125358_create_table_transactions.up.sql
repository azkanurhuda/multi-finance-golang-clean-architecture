create table transactions
(
    id         varchar(100) not null,
    nik       varchar(100) not null,
    contract_number       varchar(255) not null,
    otr       bigint not null,
    admin_fee       bigint not null,
    amount_of_interest       bigint not null,
    installment_amount       bigint not null,
    asset_name       varchar(255) not null,
    created_at timestamp       not null,
    updated_at timestamp       not null,
    primary key (id)
) engine = InnoDB;