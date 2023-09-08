// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: query.sql

package postgresql

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createOrder = `-- name: CreateOrder :one

INSERT INTO
    barista.barista_orders (
        id,
        item_type,
        item_name,
        time_up,
        created,
        updated
    )
VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, item_type, item_name, time_up, created, updated
`

type CreateOrderParams struct {
	ID       uuid.UUID    `json:"id"`
	ItemType int32        `json:"item_type"`
	ItemName string       `json:"item_name"`
	TimeUp   time.Time    `json:"time_up"`
	Created  time.Time    `json:"created"`
	Updated  sql.NullTime `json:"updated"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (BaristaBaristaOrder, error) {
	row := q.db.QueryRowContext(ctx, createOrder,
		arg.ID,
		arg.ItemType,
		arg.ItemName,
		arg.TimeUp,
		arg.Created,
		arg.Updated,
	)
	var i BaristaBaristaOrder
	err := row.Scan(
		&i.ID,
		&i.ItemType,
		&i.ItemName,
		&i.TimeUp,
		&i.Created,
		&i.Updated,
	)
	return i, err
}
