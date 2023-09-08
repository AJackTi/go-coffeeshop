// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package postgresql

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type OrderLineItem struct {
	ID             uuid.UUID     `json:"id"`
	ItemType       int32         `json:"item_type"`
	Name           string        `json:"name"`
	Price          string        `json:"price"`
	ItemStatus     int32         `json:"item_status"`
	IsBaristaOrder bool          `json:"is_barista_order"`
	OrderID        uuid.NullUUID `json:"order_id"`
	Created        time.Time     `json:"created"`
	Updated        sql.NullTime  `json:"updated"`
}

type OrderOrder struct {
	ID              uuid.UUID    `json:"id"`
	OrderSource     int32        `json:"order_source"`
	LoyaltyMemberID uuid.UUID    `json:"loyalty_member_id"`
	OrderStatus     int32        `json:"order_status"`
	Updated         sql.NullTime `json:"updated"`
}
