package models

import (
	"github.com/jinzhu/gorm"
)

// Structs for data in mongodb

// Order struct
type Order struct {
	ID         string `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt  string `json:"created_at,omitempty" bson:"created_at,omitempty"`
	OrderName  string `json:"order_name,omitempty" bson:"order_name,omitempty"`
	CustomerID string `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
}

// Customer struct
type Customer struct {
	ID          string `json:"_id,omitempty" bson:"_id,omitempty"`
	Login       string `json:"login,omitempty" bson:"login,omitempty"`
	Password    string `json:"password,omitempty" bson:"password,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	CompanyID   string `json:"company_id,omitempty" bson:"company_id,omitempty"`
	CreditCards string `json:"credit_cards,omitempty" bson:"credit_cards,omitempty"`
}

// CustomerCompany struct
type CustomerCompany struct {
	ID          string `json:"_id,omitempty" bson:"_id,omitempty"`
	CompanyName string `json:"company_name,omitempty" bson:"company_name,omitempty"`
}

// Structs for data in postgres

// Delivery struct
type Delivery struct {
	gorm.Model
	DeliveryID        int
	OrderItemID       int
	DeliveredQuantity int
}

// OrderItem struct
type OrderItem struct {
	gorm.Model
	OrderItemID  int
	OrderID      int
	PricePerUnit float64
	Quantity     int
	Product      string
}
