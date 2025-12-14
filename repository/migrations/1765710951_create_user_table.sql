-- +migrate Up
CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    phone_number VARCHAR(11),
    email VARCHAR(255),
    password VARCHAR(255)
);

-- +migrate Down
DROP TABLE users;