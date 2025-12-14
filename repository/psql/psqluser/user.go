package psqluser

import (
	"context"
	"fmt"

	"github.com/HosseinForouzan/user-management/entity"
)

func (d *DB) GetUserByID(id uint) (entity.User, error) {
	var user entity.User

	err := d.conn.Conn().QueryRow(context.Background(), "SELECT * FROM users WHERE id=$1", 1).
	Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Email, &user.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("Query Failed. %w", err)
	}

	return user, nil

}