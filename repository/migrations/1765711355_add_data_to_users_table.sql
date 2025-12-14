-- +migrate Up
INSERT INTO users(name, phone_number, email, password)
VALUES('hossein', '09391234567', 'h@h.com', '1234');

-- +migrate Down
// TODO - implement me