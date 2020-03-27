# GraphQL Server Proof of Concept Implemented In Go

## Dependencies

* Go 1.14
* PostgreSQL

## Config

Server configuration is done via the following Environment Variables:

* `DBURL` Database connection string. Required.
* `PORT` Service Port. Optional, default 8080.

You can find example config in `.envrc.example`. 

You can also use [`direnv`](https://direnv.net/) to better manage envvars.

## Run

1. Make sure your Environment Variables are configured.
2. Create the postgres database. e.g. `createdb <your_db>`
3. Run `go run server.go` to boot the server.
4. Go to [http://localhost:8080](http://localhost:8080) to use [GraphQL Playground](http://localhost:8080)

## Tech Stack

* [gqlgen](https://gqlgen.com/) - go generate based graphql server library.
* [gorm](https://gorm.io/docs/) - The fantastic ORM library for Golang.
