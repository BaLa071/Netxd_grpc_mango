package netxdcustomerinterfaces

import netxdcustomermodels "DemoProject/Netxd_Customer_Dal/Netxd_Customer_models"

type ICustomer interface{
	CreateCustomer(customer *netxdcustomermodels.Customer)(*netxdcustomermodels.DBResponse,error)
}