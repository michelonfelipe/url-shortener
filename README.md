# URL Shortener

Do you know [bit.ly](https://bit.ly)? Well... Basically the same thing

## Setup the app
### Environment variables
To setup local environment variables, you can create a `.env` file. To create that, you can copy `.env.sample` with:
```shell
cp .env.sample .env
```
You can fill each variable with the following:

#### `DATABASE_URL`
Url to connect with database.

Template: `postgres://user:password@host:port/database_name`

Example: `postgres://postgres_user:password@localhost:5432/url_shortener`

**NOTE** Remember to create the database before running the app.

#### `GIN_MODE`
Mode that your app is running. It can be:
-  `debug`: local environment for developing;
-  `release`: production environment;
-  `test`: test environment.

## Running the app
Just run:

```shell
go run main.go
```

## Testing
Just run:

```shell
go test ./...
```
