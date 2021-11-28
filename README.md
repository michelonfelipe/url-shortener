# URL Shortener

Do you know [bit.ly](https://bit.ly)? Well... Basically the same thing

## Setup the app

You can run the app with or without docker. Both instruction will be avaliable below

### With Docker
#### Running the app
Run:

```shell
docker-compose up
```

Or, if you wanna docker running on the background

```shell
docker-compose up -d
```

#### Testing
With the app running, run:

```shell
docker exec url_shortener go test ./...
```

**NOTE:** To run tests, your local database should be empty. ~~sorry, still working on that~~

### Without docker

#### Setup

1. You should create the database and make it avaliable for connection before running/testing the app;
2. Create a `.env` file to be read by the app. You can do it running:
```shell
cp .env.sample .env
```
3. Fill the env file (instruction in [environment variables section](#environment-variables)).


#### Running the app
Run:

```shell
go run main.go
```

#### Testing
Run:

```shell
go test ./...
```

**NOTE:** To run tests, your local database should be empty. ~sorry, still working on that~

## Environment variables
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

#### `SHORTENED_URL_CHARS_NUMBER`
Number of characters on the shortened URL path. Example:

```
SHORTENED_URL_CHARS_NUMBER=5
host.com/XXXXX


SHORTENED_URL_CHARS_NUMBER=10
host.com/XXXXXXXXXX
```
