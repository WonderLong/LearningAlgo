package main

import (
	"../rpc/client"
	"../rpc/protocol"
)

func main() {
	c := &client.Client{"tcp4", "localhost:8090", nil}
	c.Init()

	params := []string{";ads;dfskdddddddsgjddfdsafadsfadsffksl;", ";adsfadsfasdfdfsdfdsdfsdfrewewrweew;", ";adsgfjdfsjg;sdfsdfeqwrtioeroijadslkdmflkadsjnfkladmsfklmkgjdfksl;"}
	req := &protocol.CallRequest{"aaa", params}
	c.Call(*req)
	c.Call(*req)
	c.Call(*req)
	c.Call(*req)
	c.Call(*req)

	//c := client.RPCClient{Addr:"localhost:8090", Network:"tcp"}
	//c.Start()
	//
	//
	//o := &service.Order{}
	//c.Invoke("OrderService.GetById", 10, o)
	//fmt.Println(o)
	//
	//for i := 0; i < 100000; i++ {
	//	k := i
	//	go func() {
	//		j := &service.Input{}
	//		ro := &service.Output{}
	//		c.Invoke("OrderService.GetList", j, ro)
	//		fmt.Println(k, " == ", ro.Pa)
	//	}()
	//}

	//
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
}
