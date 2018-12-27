package server

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type Server struct {
	Port      int
	Conn      *net.TCPConn
	ConnCache map[net.TCPConn]bool
}

func (s *Server) Start() error {
	service := fmt.Sprint(":", s.Port)
	s.ConnCache = make(map[net.TCPConn]bool)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)

	if err != nil {
		return err
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)

	var ec = 0
	for {
		//accept new connection
		nc, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("err when listening, ", err)
			if ec < 10 {
				continue
			} else {
				time.Sleep(3 * time.Second)
				ec = 0
			}
			ec++
		}
		fmt.Println("......aaa.....", ec)
		go s.handleWithScanner(nc)
	}

	return nil
}

func (s Server) handleWithScanner(conn *net.TCPConn) {
	//解决粘包问题，找到分隔点，这个方法很不错
	//以下两个关键点：
	// 1.scanner的split
	scanner := bufio.NewScanner(conn)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		//对长度8的判断是因为，我们的一段数据是从标志字：MINI（长度是4）和 4个byte的length字段开始。
		//所以不同协议规则的分界点，需要重新定义
		if !atEOF && len(data) >= 8 && data[0] == 'M' && data[1] == 'I' && data[2] == 'N' && data[3] == 'I' {
			length := uint32(0)
			binary.Read(bytes.NewReader(data[4:8]), binary.BigEndian, &length)
			if int(length)+8 <= len(data) {
				return int(length) + 8, data[:int(length)+8], nil
			}
		}
		return
	})
	// 2.scanner的阻塞式监听
	for scanner.Scan() {
		buf := scanner.Bytes()
		clen := uint32(0)
		flag := buf[0:4]
		if string(flag) == "MINI" {
			clen = binary.BigEndian.Uint32(buf[4:8])
		}
		fmt.Println(string(buf[8:]), clen)
	}
}

func (s Server) handleConnection(conn *net.TCPConn) {
	buf := make([]byte, 128)
	r := bufio.NewReader(conn)
	//todo handle received data
	//0. content flag: MINI
	//1. get the content size
	//2. collect the bytes loop by loop
	//3. transform bytes to object
	content := make([]byte, 2048)

	for {
		//直接断开客户端，conn并不会马上销毁，调用会返回error
		n, err := r.Read(buf)
		if err != nil {
			//conn 被强行断开后，会有error，但同时Close操作时也会有error
			//应该在尝试Close之后，退出逻辑
			conn.Close()
			return
		}
		if n > 0 {
			parseData(buf, content, n)
		}

	}
}

func parseData(buf []byte, content []byte, n int) {
	clen := uint32(0)
	flag := buf[0:4]
	if string(flag) == "MINI" {
		clen = binary.BigEndian.Uint32(buf[4:8])
	}
	b := buf[8:]
	content = append(content, b...)
	clen = clen - uint32(n)
	fmt.Println(string(content), clen)
}
