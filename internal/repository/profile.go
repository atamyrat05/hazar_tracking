package repository

import (
	"errors"
	"fmt"
	"hazar_tracking/internal/model"
	"strings"

	"github.com/jmoiron/sqlx"
)

type ProfileRepository struct {
	db *sqlx.DB
}

func NewProfileRepository(db *sqlx.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (s *ProfileRepository) GetAll() ([]model.User, error) {
	var data []model.User
	query :=
		`SELECT * FROM users`
	err := s.db.Select(&data, query)
	return data, err
}

func (s *ProfileRepository) GetById(userId int) (model.UpdateProfileInput, error) {
	var data model.UpdateProfileInput
	query :=
		`SELECT 
		id,
		first_name,
		last_name,
		email,
		phone_number,
		image_urls
		FROM users WHERE id=$1`
	err := s.db.Get(&data, query, userId)
	return data, err
}

func (s *ProfileRepository) Delete(userId int) error {
	query :=
		`DELETE FROM users WHERE id=$1`
	result, err := s.db.Exec(query, userId)
	num, _ := result.RowsAffected()
	if num <= 0 {
		return errors.New("user not found")
	}
	return err
}

func (s *ProfileRepository) Update(userId int, input model.UpdateProfileInput) error {
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argsId := 1

	if len(input.First_name) != 0 {
		setValue = append(setValue, fmt.Sprintf("first_name=$%d", argsId))
		args = append(args, input.First_name)
		argsId++
	}
	if len(input.Last_name) != 0 {
		setValue = append(setValue, fmt.Sprintf("last_name=$%d", argsId))
		args = append(args, input.Last_name)
		argsId++
	}
	if len(input.Email) != 0 {
		setValue = append(setValue, fmt.Sprintf("email=$%d", argsId))
		args = append(args, input.Email)
		argsId++
	}
	if len(input.PhoneNumber) != 0 {
		setValue = append(setValue, fmt.Sprintf("phone_number=$%d", argsId))
		args = append(args, input.Email)
		argsId++
	}
	if len(*input.Image_Url) != 0 {
		setValue = append(setValue, fmt.Sprintf("image_urls=$%d", argsId))
		args = append(args, input.Image_Url)
		argsId++
	}

	setQuery := strings.Join(setValue, ", ")
	query := fmt.Sprintf("UPDATE users SET %s WHERE id=$%d", setQuery, argsId)
	args = append(args, userId)

	result, err := s.db.Exec(query, args...)
	san, _ := result.RowsAffected()
	if san == 0 {
		return errors.New("data not found")
	}
	return err
}
