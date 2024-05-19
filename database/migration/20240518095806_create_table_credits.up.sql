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