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