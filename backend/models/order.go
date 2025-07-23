package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type OrderProduct struct {
	OrderID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProductID uuid.UUID `gorm:"type:uuid;primaryKey"`
	Quantity  int       `gorm:"not null"`
	Product   Product   `gorm:"foreignKey:ProductID"`
}

type Order struct {
	ID         uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
	ClientID   uuid.UUID      `gorm:"type:uuid" json:"clientid"`
	Client     Client         `gorm:"foreignKey:ClientID;references:ID"`
	Products   []OrderProduct `gorm:"foreignKey:OrderID" json:"products"`
	TotalPrice float64        `gorm:"default:0" json:"totalPrice"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	Status     string         `json:"status"`
	Processed  int            `json:"-" gorm:"default:0"`
}
