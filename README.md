# Kazokku Users API

## clone this repository
```bash
git clone https://github.com/your-username/kazokku-users.git
```

## install dependencies
```bash
go mod tidy
```

## run the sql file on folder db 
```bash
# db/db.sql

-- Create the database
CREATE DATABASE IF NOT EXISTS kazokku_users;

-- Switch to the new database
\c kazokku_users;

-- Create the users table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    key UUID,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    photos VARCHAR(255)[] NOT NULL,
    creditcard_type VARCHAR(255) NOT NULL,
    creditcard_number VARCHAR(255) NOT NULL,
    creditcard_name VARCHAR(255) NOT NULL,
    creditcard_expired VARCHAR(255) NOT NULL,
    creditcard_cvv VARCHAR(255) NOT NULL
);

```


## Configure the following environment variables for the server:

```bash
//server setting
SERVER_URI=localhost
SERVER_PORT=3000

//database setting
DB_HOST=localhost
DB_PORT=5432
DB_NAME=kazokku_users
DB_USERNAME=postgres
DB_PASSWORD=dimasslalu123
DB_POOL_MIN=10
DB_POOL_MAX=100
DB_TIMEOUT=10
DB_MAX_IDLE_TIME_SECOND=60

//cloudinary setting
CLOUD_NAME=dzpzhe5xs
CLOUD_API_KEY=824762265846988
CLOUD_API_SECRET=x0-Kpe2kLdFraXWwQb7vESf_Sq8
```

## Run the applications

```bash
go run main.go
```

## Postman Documentation

```bash
https://documenter.getpostman.com/view/23663611/2s9YsDjaB1
```
