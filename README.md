# sqlc-pgx-migrate-template
This is a boilerplate template for the following combination of libraries: [sqlc](https://github.com/sqlc-dev/sqlc), [pgx/v5](https://github.com/jackc/pgx), [golang-migrate](https://github.com/golang-migrate/migrate). This is just one of the many examples of how to use these libraries together — don't consider this as the single right way of doing things and tinker around with them yourself.

**How to use this code?**

- Put your migrations into the `internal/db/migrations` directory.
- Put your queries into the `internal/db/queries` directory and add them to the `sqlc.yaml` configuration.
- Use the `db.Setup` function to run the migrations and get the pgx connection pool.

## Running the example

```sh
# To start the example application
$ docker-compose up -d

# To stop the example application
$ docker-compose down
```

```sh
$ curl localhost:8080/authors

[{"ID":"7e96f8fc-c2c6-4844-a588-9bd14b066ffc","Name":"Robert Jordan","Description":null,"AverageRating":4.225}]
```

## Cleaning up the template

```sh
# To remove the example application, its Docker files, this README file, and initialize a new git repository.
# Do not forget to delete this rule after you've run it.
$ make cleanup-template

# To change the module name.
$ go mod edit -module github.com/you/your-repository
```
