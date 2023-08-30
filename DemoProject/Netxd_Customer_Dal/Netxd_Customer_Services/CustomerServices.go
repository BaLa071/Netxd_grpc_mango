package netxdcustomerservices

import (
	netxdcustomerinterfaces "DemoProject/Netxd_Customer_Dal/Netxd_Customer_interfaces"
	netxdcustomermodels "DemoProject/Netxd_Customer_Dal/Netxd_Customer_models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}

func InitCustomerService(collection *mongo.Collection, ctx context.Context) netxdcustomerinterfaces.ICustomer {
	return &CustomerService{collection, ctx}
}

func (c *CustomerService) CreateCustomer(user *netxdcustomermodels.Customer) (*netxdcustomermodels.DBResponse, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.IsActive = false
	// user.FirstName="jp"

	res, err := c.CustomerCollection.InsertOne(c.ctx, &user)
	if err != nil {
		return nil, errors.New("Sorry! cannot create your profile...")
	}

	var newUser *netxdcustomermodels.DBResponse
	query := bson.M{"_id": res.InsertedID}

	err = c.CustomerCollection.FindOne(c.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
