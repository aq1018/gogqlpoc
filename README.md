# GraphQL Server Proof of Concept Implemented In Go

## Tech Stack

* [gqlgen](https://gqlgen.com/) - go generate based graphql server library.
* [gorm](https://gorm.io/docs/) - The fantastic ORM library for Golang.

## Step for Docker

### Dependencies

* Docker
* docker-compose

### Run

```
docker-compose up
```

Now open [http://localhost:8080](http://localhost:8080) in your browser.

## Setup for Bare Metal

### Dependencies

* Go 1.14
* PostgreSQL

### Config

Server configuration is done via the following Environment Variables:

* `DBURL` Database connection string. Required.
* `PORT` Service Port. Optional, default 8080.

You can find example config in `.envrc.example`. 

You can also use [`direnv`](https://direnv.net/) to better manage envvars.

### Run

Before you start, make sure you have:

* Make sure your Environment Variables are configured.
* Create the postgres database. e.g. `createdb <your_db>`

Run the following command:

```
go run server.go
```


Now open [http://localhost:8080](http://localhost:8080) in your browser.

