package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderID     uint64
	CustomerID  uuid.UUID  `json:"order_id"`
	LineItems   []LineItem `json:"line_items"`
	createdAt   *time.Time `json:"create_at"`
	shippedAt   *time.Time `json:"shipped_at"`
	completedAt *time.Time `json:"completed_at"`
}

type LineItem struct {
	ItemID   uuid.UUID `json:"item_id"`
	Quantity uint      `json:"quantity"`
	Price    uint      `json:"price"`
}
