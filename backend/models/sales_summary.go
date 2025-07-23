package models

type SalesSummary struct {
	ID              uint    `gorm:"primaryKey;autoIncrement:false"`
	TotalSalesMade  int     `json:"totalSalesMade"`
	InvoicedAmount  float64 `json:"invoicedSmount"`
	TotalOrdersSold int     `json:"totalOrdersSold"`
}
