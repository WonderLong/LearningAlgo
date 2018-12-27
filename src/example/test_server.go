package main

import (
	"../rpc/server"
	//"../service"
)

func main() {
	s := server.Server{8090, nil, nil}
	s.Start()

	//var exportedServices []interface{}
	//exportedServices = append(exportedServices, new(service.OrderService))
	//
	//s := server.RPCServer{":8090", "tcp", exportedServices}
	//s.Start()
}
