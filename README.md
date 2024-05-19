## Tech Stack

- Golang : https://github.com/golang/go
- MySQL (Database) : https://github.com/mysql/mysql-server

## Security
#### Supports the prevention of security attacks in the OWASP TOP 10

- Broken Access Control
- SQL Injection
- Security Misconfiguration
- Identification and Authentication Failures 
- Security Logging and Monitoring Failures

## Architecture

![Clean Architecture](architecture.png)

## Entity Relationship Diagram

![ERD](erd.png)

## Documentation

![Documentation](multi_finance.json)

## Configuration

All configuration is in `development.yaml` file.

## Database Migration

All database migration is in `database/migration` folder.

### Create Migration

```shell
migrate create -ext sql -dir database/migration create_table_xxx
```

### Run Migration

```shell
migrate -database "mysql://root:root@tcp(localhost:3306)/multi_finance?charset=utf8mb4&parseTime=True&loc=Local" -path database/migration up
```

## Run Application

### Run app

```bash
go run cmd/web/main.go
```

### Docker Compose

```bash
docker compose up -d
```