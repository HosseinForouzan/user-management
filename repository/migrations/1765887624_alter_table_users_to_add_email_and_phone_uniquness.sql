-- +migrate Up
ALTER TABLE users
ADD CONSTRAINT users_email_phone_number_unique UNIQUE(email, phone_number);

-- +migrate Down
ALTER TABLE users
DROP CONSTRAINT users_email_phone_number_unique;