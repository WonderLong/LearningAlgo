package server

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type RPCServer struct {
	Port           string
	Network        string
	ExportServices []interface{}
}

func (s *RPCServer) Start() {
	s.exportService()
	rpc.HandleHTTP()
	l, e := net.Listen(s.Network, s.Port)

	if e != nil {
		log.Fatal("cannot listen to port: ", s.Port)
	}

	http.Serve(l, nil)

	//in := make(chan int64)
	//
	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	i, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	//	in <- i
	//}
	//
	//select {
	//	case i := <- in:
	//		fmt.Println("close", i)
	//}
}

func (s *RPCServer) exportService() error {
	if len(s.ExportServices) == 0 {
		return errors.New("No exported service")
	}
	for _, srvc := range s.ExportServices {
		rpc.Register(srvc)
	}
	return nil
}
