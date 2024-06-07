create table assets
(
    id varchar(100) not null,
    merchant_id varchar(100) not null,
    name varchar(255) not null,
    otr decimal(10,2) not null,
    admin_fee decimal(10,2) not null,
    description varchar(255) not null,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    primary key (id),
    foreign key (merchant_id) references merchants(id)
) engine = InnoDB;