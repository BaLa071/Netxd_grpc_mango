package netxdcustomercontroller

import (
	// netxdcustomerservices "DemoProject/Netxd_Customer_Dal/Netxd_Customer_Services"
	netxdcustomerinterfaces "DemoProject/Netxd_Customer_Dal/Netxd_Customer_interfaces"
	netxdcustomermodels "DemoProject/Netxd_Customer_Dal/Netxd_Customer_models"
	pro "DemoProject/Netxd_Customer_proto/Netxd_Customer"
	"context"
	"fmt"
)

type RPCServer struct {
	pro.UnimplementedNetxd_DalServiceServer
}

var (
	CustomerService netxdcustomerinterfaces.ICustomer
)

func (c *RPCServer) CreateCustomer(ctx context.Context, req *pro.CustomerProfile) (*pro.CustomerResponse, error) {
		dbCustomer := &netxdcustomermodels.Customer{CustomerId: req.CustomerId, FirstName: req.FirstName, LastName: req.LastNamae, Balance: float64(req.Balance)}
	if err != nil {
		return nil, err
	} else {
		responseProfile := &pro.CustomerResponse{
			CustomerId: 1,

			// LastName:  result.LastName,
		}
		return responseProfile, nil
	}
}
