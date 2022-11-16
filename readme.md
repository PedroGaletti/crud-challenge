## Pismo Challenge

This project is a challenge from Pismo to store accounts and your transactions

![Build Badge](https://img.shields.io/github/workflow/status/PedroGaletti/pismo-challenge/pismo)
[![Codecov Coverage Badge](https://img.shields.io/codecov/c/gh/PedroGaletti/pismo-challenge)](https://img.shields.io/codecov/c/gh/PedroGaletti/pismo-challenge)

## Stack

- [Golang](https://go.dev) - Build fast, reliable, and efficient software at scale
- [Docker](https://www.docker.com) - Accelerate how you build, share, and run modern applications
- [MySQL](https://www.mysql.com) - Database management system, which uses the SQL language as an interface
- [Gin](https://github.com/gin-gonic/gin) - Gin is a HTTP web framework written in Go (Golang)

## Project structure

```
$PROJECT_ROOT
├── cmd
│   └── accounts         # Sequence and stats api logic
│   └── operations       # Operation model
│   └── transactions     # Transaction api logic
├── configs              # Dot env and log configs
└── db                   # Database logic and connection 
```

## Environment Variables

```
Variable                | Type    | Description                       | Default
----------------------- | ------- | --------------------------------- | ------------------------
GIN_MODE                | string  | Interval time of the cron         | debug
LOG_LEVEL               | string  | Leveled Logging                   | info
SQL_DB                  | string  | Database                          | pismo
SQL_HOST                | string  | Database Host                     | 127.0.0.1
SQL_PASSWORD            | string  | Database Password                 | root
SQL_PORT                | string  | Database Port                     | 3308
SQL_USER                | string  | Database user                     | root
```


## Make commands:

Assuming that you have already cloned the project and the [Go](https://golang.org/doc/install) is installed, ensure that all dependencies are vendored in the project:

```
make install
```

To build the application:

```
make build
```

To run the application local:

```
make run
```

## URL Requests

- [POST](http://localhost:8080/accounts) - /accounts
```
Body Request
{
  "document_number": "12345678900"
}
```
- [GET](http://localhost:8080/accounts/:accountId) - /accounts/:accountId
```
Body Response
{
  "id": 1,
  "document_number": "12345678900"
}
```
- [POST](http://localhost:8080/transactions) - /transactions
```
Body Request
{
  "account_id": 1,
  "operation_id": 3,
  "amount": 12345
}
amount need be int value without commas or dots
```

## Author

- [@pedrogaletti](https://www.github.com/PedroGaletti)
