package repository

import (
	"errors"
	"hazar_tracking/internal/model"

	"github.com/jmoiron/sqlx"
)

type EmailRepository struct {
	db *sqlx.DB
}

func NewEmailsRepository(db *sqlx.DB) *EmailRepository {
	return &EmailRepository{db: db}
}

func (r *EmailRepository) Validate(email string) (int, error) {
	var id int
	query :=
		`SELECT id FROM users WHERE email=$1`
	row := r.db.QueryRow(query, email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *EmailRepository) UpdateForgotCode(userId int, code string) error {
	query :=
		`UPDATE users SET forgot_pas_code=$1 WHERE id=$2`
	_, err := r.db.Exec(query, code, userId)
	return err
}

func (r *EmailRepository) TestCode(code string, userId int) error {
	query :=
		`SELECT id FROM users WHERE id=$1 AND forgot_pas_code=$2`
	result, err := r.db.Exec(query, userId, code)
	if err != nil {
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if num <= 0 {
		return errors.New("invalid code")
	} else {
		return nil
	}
}

func (r *EmailRepository) UpdateUsersPassword(input model.UpdatePassword, userId int) error {
	query :=
		`UPDATE users SET password=$1, password_confirm=$2 WHERE id=$3`
	result, err := r.db.Exec(query, input.Password, input.Password_confirm, userId)
	if err != nil {
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if num <= 0 {
		return errors.New("password not updated")
	} else {
		return nil
	}
}
