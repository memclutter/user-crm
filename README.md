# User CRM

CRM example. User database management.

## Features

- CRUD operations
- Import/Export list as csv
- Filter & sort

## Database migrations

Use https://github.com/golang-migrate/migrate. Download binary and run

```
migrate -path migrations -database 'postgres://user:secret@localhost:5432/test?sslmode=disable' up
```

Generate new migration file

```
migrate create -ext sql -dir migrations [name]
```

## Swagger docs

Swagger docs available on address http://localhost:9000/api/docs/index.html

## Watch changes with modd

Install https://github.com/cortesi/modd for best development experience. Run local server

```
modd 
```