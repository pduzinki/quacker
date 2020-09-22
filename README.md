## About:
This is a simple 'twitter clone' web app, but instead of tweeting, it's quacking! It's meant to be
a learning exercise for me, and hopefully for you as well. New features and improvements to come soon...

## Requirements:
* Go v1.14.2 (or higher)
* PostgreSQL v10.14 (or higher)

## Additional packages:
* GORM v1 (https://v1.gorm.io/)
* gorilla/mux (https://github.com/gorilla/mux)
* gorilla/schema (https://github.com/gorilla/schema)
* bcrypt (golang.org/x/crypto/bcrypt)

## Setup:
Before running the app for the first time, you need to create database in PostgreSQL. Login to Postgres and type:
```sh
CREATE DATABASE quacker;
```

Next you need to prepare configuration file 'config.json'. It should look similar to that, just remember to fill all the fields with your data:

```json
{
    "dbConfig":{
        "dialect":"postgres",
        "host":"localhost",
        "port":5432,
        "user":"your-postgres-user",
        "password":"your-password",
        "name":"quacker"
    },
    "passwordPepper":"a-secret-pepper",
    "hmacKey":"a-secret-hmac-key"
}
```

## Running:
```sh
go run *.go
```

## License:
Quacker is licensed under MIT License. See LICENSE file.
