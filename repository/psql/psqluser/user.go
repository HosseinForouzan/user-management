package psqluser

import (
	"context"
	"fmt"

	"github.com/HosseinForouzan/user-management/entity"
	"github.com/jackc/pgx/v5"
)

func (d *DB) Register(u entity.User)(entity.User, error) {
	var id uint
	err := d.conn.Conn().QueryRow(context.Background(), `INSERT INTO users(name, phone_number, email, password) VALUES($1,$2,$3,$4) RETURNING id`,
															u.Name, u.PhoneNumber, u.Email, u.Password).Scan(&id)
	if err != nil {
		return entity.User{}, fmt.Errorf("can't execute query: %w", err)
	}

	u.ID = id

	return u, nil

}

func (d *DB) IsEmailUnique(email string) (bool, error) {
	var id uint
	err := d.conn.Conn().QueryRow(context.Background(), `SELECT id FROM users WHERE email = $1`, email).Scan(&id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return true, nil
		}

		return false, fmt.Errorf("something wrong: %w", err)
	}

	return false, nil
}

func (d *DB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	var id uint
	err := d.conn.Conn().QueryRow(context.Background(), `SELECT id FROM users WHERE phone_number = $1`, phoneNumber).Scan(&id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return true, nil
		}

		return false, fmt.Errorf("something wrong: %w", err)
	}

	return false, nil
}


func (d *DB) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User

	err := d.conn.Conn().QueryRow(context.Background(), "SELECT * FROM users WHERE email=$1", email).
	Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Email, &user.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("Query Failed. %w", err)
	}

	return user, nil

}

func (d *DB) GetUserByPhoneNumber(phone_number string) (entity.User, error) {
	var user entity.User

	err := d.conn.Conn().QueryRow(context.Background(), "SELECT * FROM users WHERE phone_number=$1", phone_number).
	Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Email, &user.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("Query Failed. %w", err)
	}

	return user, nil

}