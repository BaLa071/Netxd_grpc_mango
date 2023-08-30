package main

import (
	netxdcustomerservices "DemoProject/Netxd_Customer_Dal/Netxd_Customer_Services"
	pro "DemoProject/Netxd_Customer_proto/Netxd_Customer"
	netxdcustomerconfig "DemoProject/Netxd_Customer_server/Netxd_Customer_Config"
	netxdcustomerconstants "DemoProject/Netxd_Customer_server/Netxd_Customer_Constants"
	netxdcustomercontroller "DemoProject/Netxd_Customer_server/Netxd_Customer_Controller"
	"context"
	"fmt"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	customerCollection := netxdcustomerconfig.GetCollection(client, "customer", "profiles")
	netxdcustomercontroller.CustomerService = netxdcustomerservices.InitCustomerService(customerCollection, context.Background())
}

func main() {
	mongoclient, err := netxdcustomerconfig.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	lis, err := net.Listen("tcp", netxdcustomerconstants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pro.RegisterNetxd_DalServiceServer(s, &netxdcustomercontroller.RPCServer{})

	fmt.Println("Server listening on", netxdcustomerconstants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
