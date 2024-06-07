create table credit_limits
(
    id varchar(100) not null,
    user_id varchar(100) not null,
    tenor int not null,
    credit_limit decimal(10,2) not null,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    primary key (id),
    foreign key (user_id) references users(id)
) engine = InnoDB;