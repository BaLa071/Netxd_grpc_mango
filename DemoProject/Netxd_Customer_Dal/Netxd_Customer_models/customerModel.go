package netxdcustomermodels

import (
	"time"

	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	CustomerId int32     `json:"customerId" bson:"customerId"`
	FirstName  string    `json:"firstName" bson:"firstname"`
	LastName   string    `json:"lastName" bson:"lastname"`
	BankId     int32     `json:"bankId" bson:"bankId"`
	Balance    float64   `json:"balance" bson:"balance"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" bson:"updatedAt"`
	IsActive   bool      `json:"isActive" bson:"isActive"`
}

type DBResponse struct {
	// ID         primitive.ObjectID `json:"id" bson:"_id"`
	CustomerId int32              `json:"customerId" bson:"customerId"`
	FirstName  string             `json:"firstName" bson:"firstname"`
	// LastName   string             `json:"lastName" bson:"lastname"`
	// BankId     int32              `json:"bankId" bson:"bankId"`
	// Balance    float64            `json:"balance" bson:"balance"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt"`
	// UpdatedAt  time.Time          `json:"updatedAt" bson:"updatedAt"`
	// IsActive   bool               `json:"isActive" bson:"isActive"`
}
