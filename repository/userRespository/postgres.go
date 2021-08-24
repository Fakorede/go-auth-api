package userrespository

import (
	"database/sql"
	"goauthapi/models"
)

type UserRepository struct {}

func (u UserRepository) Signup(db *sql.DB, user models.User) (models.User, error) {
	stmt := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return user, err
	}

	user.Password = ""
	return user, nil
}

func (u UserRepository) Login(db *sql.DB, user models.User) (models.User, error) {
	row := db.QueryRow(`SELECT * FROM users WHERE email = $1`, user.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}