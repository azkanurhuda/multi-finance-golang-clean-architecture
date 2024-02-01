## Tech Stack

- Golang : https://github.com/golang/go
- MySQL (Database) : https://github.com/mysql/mysql-server

## Architecture

![Clean Architecture](architecture.png)

## Entity Relationship Diagram

![ERD](erd.png)

## Framework & Library

- GoFiber (HTTP Framework) : https://github.com/gofiber/fiber
- GORM (ORM) : https://github.com/go-gorm/gorm
- Viper (Configuration) : https://github.com/spf13/viper
- Golang Migrate (Database Migration) : https://github.com/golang-migrate/migrate
- Go Playground Validator (Validation) : https://github.com/go-playground/validator
- Logrus (Logger) : https://github.com/sirupsen/logrus

## Configuration

All configuration is in `config.json` file.

## Database Migration

All database migration is in `database/migration` folder.

### Create Migration

```shell
migrate create -ext sql -dir database/migration create_table_xxx
```

### Run Migration

```shell
migrate -database "mysql://root:root@tcp(localhost:3306)/multi_finance_golang_clean_architecture?charset=utf8mb4&parseTime=True&loc=Local" -path database/migration up
```

## Run Application

### Run app

```bash
go run cmd/web/main.go
```