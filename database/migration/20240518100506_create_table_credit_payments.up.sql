create table credit_payments
(
    id         varchar(100) not null,
    credit_id       varchar(100) not null,
    payment_amount       decimal(10,2) not null,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    primary key (id),
    foreign key (credit_id) references credits(id)
) engine = InnoDB;