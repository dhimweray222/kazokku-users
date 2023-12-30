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
