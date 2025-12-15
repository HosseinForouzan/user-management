package psqluser

import (
	"context"
	"fmt"

	"github.com/HosseinForouzan/user-management/entity"
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

func (d *DB) GetUserByID(id uint) (entity.User, error) {
	var user entity.User

	err := d.conn.Conn().QueryRow(context.Background(), "SELECT * FROM users WHERE id=$1", 1).
	Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Email, &user.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("Query Failed. %w", err)
	}

	return user, nil

}