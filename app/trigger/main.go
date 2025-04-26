package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/doutokk/doutok/common/clientsuite"
	"rpc102/app/uid/kitex_gen/uid"
	"rpc102/app/uid/kitex_gen/uid/uidservice"
	"rpc102/app/user/kitex_gen/user"
	"rpc102/app/user/kitex_gen/user/userservice"
)

func main() {
	triggerUser()
}

func triggerUid() {

	fmt.Println("ttt")
	registryAddr := "127.0.0.1:8500"
	serviceName := "uid"
	commonSuite := client.WithSuite(clientsuite.CommonGrpcClientSuite{
		RegistryAddr:       registryAddr,
		CurrentServiceName: serviceName,
	})
	newClient, err := uidservice.NewClient("uid", commonSuite)
	if err != nil {
		fmt.Println(err)
	}
	r, err := newClient.UidGen(context.Background(), &uid.Empty{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(r.Uid)
}

func triggerUser() {

	registryAddr := "127.0.0.1:8500"
	serviceName := "user"
	commonSuite := client.WithSuite(clientsuite.CommonGrpcClientSuite{
		RegistryAddr:       registryAddr,
		CurrentServiceName: serviceName,
	})
	newClient, err := userservice.NewClient("user", commonSuite)
	if err != nil {
		fmt.Println(err)
	}
	r, err := newClient.Login(context.Background(), &user.LoginReq{
		Email:    "",
		Password: "",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(r.Token)
}
