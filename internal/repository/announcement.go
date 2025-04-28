package repository

import (
	"hazar_tracking/internal/model"

	"github.com/jmoiron/sqlx"
)

type AnnouncementRepository struct {
	db *sqlx.DB
}

func NewAnnouncementRepository(db *sqlx.DB) *AnnouncementRepository {
	return &AnnouncementRepository{db: db}
}

func (r *AnnouncementRepository) Create(input model.AnnouncementInput, userId int) (int, error) {
	var id int
	query :=
		`INSERT INTO announcement (category, time, from_where, where_to, text, phone_number, name) values ($1,$2,$3,$4,$5,$6,$7) RETURNING id`
	row := r.db.QueryRow(query, input.Category, input.Time, input.From_where, input.Where_to, input.Text, input.PhoneNumber, input.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AnnouncementRepository) GetAll() ([]model.AnnouncementGet, error) {
	var data []model.AnnouncementGet
	query :=
		`SELECT 
			a.id,
			a.category,
			a.time,
			pf.name AS from_where,
			pt.name AS where_to,
			a.text
			FROM announcement a
			LEFT JOIN points pf ON pf.id = a.from_where
			LEFT JOIN points pt ON pt.id = a.where_to
			`
	err := r.db.Select(&data, query)
	return data, err
}

func (r *AnnouncementRepository) GetById(announcementId int) (model.AnnouncementGetById, error) {
	var data model.AnnouncementGetById
	query :=
		`SELECT 
			a.id,
			a.category,
			a.time,
			pf.name AS from_where,
			pt.name AS where_to,
			a.text,
			a.phone_number,
			a.name
			FROM announcement a
			LEFT JOIN points pf ON pf.id = a.from_where
			LEFT JOIN points pt ON pt.id = a.where_to
			WHERE a.id=$1`
	err := r.db.Get(&data, query, announcementId)
	return data, err
}
