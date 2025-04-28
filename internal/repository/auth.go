package repository

import (
	"fmt"
	"hazar_tracking/internal/model"
	"hazar_tracking/pkg/database"

	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) Create(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (first_name, last_name, email, phone_number, password, password_confirm) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id",
		database.UsersTable)
	row := r.db.QueryRow(query, user.First_name, user.Last_name, user.Email, user.PhoneNumber, user.Password, user.Password_confirm)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	
	return id, nil
}

func (r *AuthRepository) Get(email, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE email=$1 AND password=$2", database.UsersTable)
	err := r.db.Get(&user, query, email, password)
	return user, err
}
