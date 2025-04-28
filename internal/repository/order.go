package repository

import (
	"errors"
	"fmt"
	"hazar_tracking/internal/model"
	"hazar_tracking/pkg/database"
	"os"
	"path/filepath"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type OrdersRepository struct {
	db *sqlx.DB
}

func NewOrdersRepository(db *sqlx.DB) *OrdersRepository {
	return &OrdersRepository{db: db}
}

func (r *OrdersRepository) Create(ordersInput model.OrdersInput, userId int, filename string) (int, error) {
	var id int
	qrcode_url := filepath.Join("uploads", filename)
	seria_id := filename[:len(filename)-4]
	tx, err := r.db.Begin()
	if err != nil {
		err1 := os.Remove(qrcode_url)
		if err1 != nil {
			logrus.Println("Can not deleted a creating QRCode")
		}
		tx.Rollback()
		return 0, err
	}
	query := fmt.Sprintf("INSERT INTO %s (senders_name, buyers_name, from_where, where_to, type_of_service, weight, users_id, seria_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id", database.OrdersTable)
	row := tx.QueryRow(query, ordersInput.Senders_name, ordersInput.Buyers_name, ordersInput.From_where, ordersInput.Where_to, ordersInput.Type_of_service, ordersInput.Weight, userId, seria_id)
	if err := row.Scan(&id); err != nil {
		err1 := os.Remove(qrcode_url)
		if err1 != nil {
			logrus.Println("Can not deleted a creating QRCode")
		}
		tx.Rollback()
		return 0, err
	}

	qrcodeQuery := fmt.Sprintf("INSERT INTO %s (urls,data,orders_id) VALUES ($1,$2,$3)", database.QRCodeTable)
	_, err = tx.Exec(qrcodeQuery, qrcode_url, seria_id, id)
	if err != nil {
		err1 := os.Remove(qrcode_url)
		if err1 != nil {
			logrus.Println("Can not deleted a creating QRCode")
		}
		tx.Rollback()
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		err1 := os.Remove(qrcode_url)
		if err1 != nil {
			logrus.Println("Can not deleted a creating QRCode")
		}
		return 0, err
	}
	return id, err
}

func (r *OrdersRepository) GetAll(userId int) ([]model.OrdersGet, error) {
	var orders []model.OrdersGet
	query :=
		`SELECT 
			o.id,
			o.senders_name,
			o.buyers_name, 
			pf.name AS from_where,
			pt.name AS where_to, 
			o.type_of_service, 
			o.weight, 
			o.status,
			o.seria_id,
			o.started_time,
			o.finished_time,
			pl.name AS name,

			(SELECT COUNT(*) 
			FROM order_tracking_steps s 
			WHERE s.order_id = o.id) AS total_steps,

			(SELECT COUNT(*) 
			FROM order_tracking_steps s 
			WHERE s.order_id = o.id AND s.step_date IS NOT NULL) AS current_step_number

			FROM orders o
			LEFT JOIN points pf ON pf.id = o.from_where
			LEFT JOIN points pt ON pt.id = o.where_to

			LEFT JOIN LATERAL (
				SELECT p.name
				FROM order_tracking_steps s
				JOIN points p ON p.id = s.location
				WHERE s.order_id = o.id AND s.step_date IS NOT NULL
				ORDER BY s.step_date DESC, s.created_at DESC
				LIMIT 1
			) pl ON true

				WHERE o.users_id = $1`

	err := r.db.Select(&orders, query, userId)
	return orders, err
}

func (r *OrdersRepository) GetById(userId, orderId int) (model.OrdersGet, []model.OrderTrackingSteps, error) {
	var orders model.OrdersGet
	query :=
		`SELECT 
			o.id,
			o.senders_name,
			o.buyers_name, 
			pf.name AS from_where,
			pt.name AS where_to, 
			o.type_of_service, 
			o.weight, 
			o.status,
			o.seria_id,
			o.started_time,
			o.finished_time,
			pl.name AS name,

			(SELECT COUNT(*) 
			FROM order_tracking_steps s 
			WHERE s.order_id = o.id) AS total_steps,

			(SELECT COUNT(*) 
			FROM order_tracking_steps s 
			WHERE s.order_id = o.id AND s.step_date IS NOT NULL) AS current_step_number

			FROM orders o
			LEFT JOIN points pf ON pf.id = o.from_where
			LEFT JOIN points pt ON pt.id = o.where_to

			LEFT JOIN LATERAL (
				SELECT p.name
				FROM order_tracking_steps s
				JOIN points p ON p.id = s.location
				WHERE s.order_id = o.id AND s.step_date IS NOT NULL
				ORDER BY s.step_date DESC, s.created_at DESC
				LIMIT 1
			) pl ON true

				WHERE o.users_id = $1 AND o.id=$2`

	err := r.db.Get(&orders, query, userId, orderId)
	var location []model.OrderTrackingSteps
	if err == nil {
		locationQuery :=
			`SELECT 
				p.name,
				s.step_date
				FROM order_tracking_steps s
				LEFT JOIN points p ON p.id=s.location WHERE s.order_id=$1
				ORDER BY s.id`

		err = r.db.Select(&location, locationQuery, orderId)
	}
	return orders, location, err
}

func (r *OrdersRepository) Search(userId int, input string) (model.OrdersGet, error) {
	var data model.OrdersGet
	query :=
		`SELECT * FROM orders WHERE users_id=$1 AND seria_id=$2`

	err := r.db.Get(&data, query, userId, input)
	return data, err
}

func (r *OrdersRepository) GetAllPoints() ([]model.Points, error) {
	var data []model.Points
	query :=
		`SELECT * FROM points`
	err := r.db.Select(&data, query)
	if err != nil {
		return data, errors.New("can not get all points")
	}
	return data, nil
}

func (r *OrdersRepository) Update(orderId int, input model.UpdateOrderInput) error {
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argsId := 1

	if len(input.Senders_name) != 0 {
		setValue = append(setValue, fmt.Sprintf("senders_name=$%d", argsId))
		args = append(args, input.Senders_name)
		argsId++
	}
	if len(input.Buyers_name) != 0 {
		setValue = append(setValue, fmt.Sprintf("buyers_name=$%d", argsId))
		args = append(args, input.Buyers_name)
		argsId++
	}
	if len(input.From_where) != 0 {
		setValue = append(setValue, fmt.Sprintf("from_where=$%d", argsId))
		args = append(args, input.From_where)
		argsId++
	}
	if len(input.Where_to) != 0 {
		setValue = append(setValue, fmt.Sprintf("where_to=$%d", argsId))
		args = append(args, input.Where_to)
		argsId++
	}
	if input.Type_of_service != 0 {
		setValue = append(setValue, fmt.Sprintf("type_of_service=$%d", argsId))
		args = append(args, input.Type_of_service)
		argsId++
	}
	if len(input.Weight) != 0 {
		setValue = append(setValue, fmt.Sprintf("weight=$%d", argsId))
		args = append(args, input.Weight)
		argsId++
	}
	if input.Status != 0 {
		setValue = append(setValue, fmt.Sprintf("status=$%d", argsId))
		args = append(args, input.Status)
		argsId++
	}
	if len(input.Seria_id) != 0 {
		setValue = append(setValue, fmt.Sprintf("seria_id=$%d", argsId))
		args = append(args, input.Seria_id)
		argsId++
	}
	if len(*input.Started_time) != 0 {
		setValue = append(setValue, fmt.Sprintf("started_time=$%d", argsId))
		args = append(args, input.Started_time)
		argsId++
	}
	if len(*input.Finished_time) != 0 {
		setValue = append(setValue, fmt.Sprintf("finished_time=$%d", argsId))
		args = append(args, input.Finished_time)
		argsId++
	}

	setQuery := strings.Join(setValue, ", ")
	query := fmt.Sprintf("UPDATE orders SET %s WHERE id=$%d", setQuery, argsId)
	args = append(args, orderId)

	result, err := r.db.Exec(query, args...)
	san, _ := result.RowsAffected()
	if san == 0 {
		return errors.New("data not found")
	}
	return err
}

func (r *OrdersRepository) CreatePoints(input model.OrderTrackingStepsInput) (int, error) {
	var id int
	query :=
		`INSERT INTO order_tracking_steps (order_id, location) VALUES ($1,$2) RETURNING id`
	row := r.db.QueryRow(query, input.OrderId, input.Location)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *OrdersRepository) UpdatePoints(input model.UpdateOrderTrackingStepsInput) error {
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argsId := 1

	if input.OrderId != 0 {
		setValue = append(setValue, fmt.Sprintf("order_id=$%d", argsId))
		args = append(args, input.OrderId)
		argsId++
	}
	if input.Location != 0 {
		setValue = append(setValue, fmt.Sprintf("location=$%d", argsId))
		args = append(args, input.Location)
		argsId++
	}
	if len(input.StepDate) != 0 {
		setValue = append(setValue, fmt.Sprintf("step_date=$%d", argsId))
		args = append(args, input.StepDate)
		argsId++
	}

	setQuery := strings.Join(setValue, ", ")
	query := fmt.Sprintf("UPDATE order_tracking_steps SET %s WHERE order_id=$%d AND location=$%d", setQuery, argsId, argsId+1)
	args = append(args, input.OrderId, input.Location)

	result, err := r.db.Exec(query, args...)
	san, _ := result.RowsAffected()
	if san == 0 {
		return errors.New("data not found")
	}
	return err
}
