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
migrate create -ext sql -dir migrations [name]]
```

## Swagger docs

Use https://github.com/swaggo/swag for api documentations. Run 

```
swag init -g start.go
```

for generate docs. Open docs on http://localhost:9000/api/docs/index.html