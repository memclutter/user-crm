# User CRM

This is my hobby project. As an idea mini-CRM was taken. The purpose of this system is maintaining a user database. 

Initially, it is supposed to store such user characteristics as gender, date of birth, country, etc. Also added a directory of countries that can be edited.

The plans to deploy the project to some free hosting fan for (for example on [heroku](https://www.heroku.com)). 

## Features

- CRUD operations
- Import/Export list as csv
- Filter & sort

## Technology stack

- [Golang](https://golang.org) as backend 
- [Echo](https://echo.labstack.com) as http framework
- [PostgreSQL](https://www.postgresql.org) as database
- [Vue.js](https://vuejs.org) as frontend
- [Vuetify](https://vuetifyjs.com) as ui
- [Swagger](https://swagger.io) for docs
- [Heroku?](https://www.heroku.com) as cloud platform

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