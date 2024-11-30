package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	Balance     float64   `json:"balance"`
	AffiliateID uuid.UUID `json:"affiliate_id"`
}

type Product struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Quantity int       `json:"quantity"`
	Price    float64   `json:"price"`
}

type Commission struct {
	ID          uuid.UUID `json:"id"`
	OrderID     uuid.UUID `json:"order_id"`
	AffiliateID uuid.UUID `json:"affiliate_id"`
	Amount      float64   `json:"amount"`
}

type Affiliate struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	MasterAffiliate uuid.UUID `json:"master_affiliate"`
	Balance         float64   `json:"balance"`
}

type Order struct {
	ID          uuid.UUID `json:"id"`
	AffiliateID uuid.UUID `json:"affiliate_id"`
	ProductID   uuid.UUID `json:"product_id"`
	TotalAmount float64   `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"` 
}
