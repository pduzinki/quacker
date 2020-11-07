## About:
This is a simple 'twitter clone' web app, but instead of tweeting, it's quacking! It's meant to be
a learning exercise for me, and hopefully for you as well. New features and improvements to come soon...

## Requirements:
* Go v1.14.2 (or higher)
* PostgreSQL v10.14 (or higher)

## Additional packages:
To install additional packages needed, just run:
```sh
go get github.com/gorilla/mux github.com/gorilla/schema github.com/jinzhu/gorm github.com/lib/pq golang.org/x/crypto/bcrypt
```

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
