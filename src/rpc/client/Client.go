package client

import (
	. "../protocol"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type Client struct {
	Network    string
	ServerAddr string
	Conn       *net.TCPConn
}

func (c *Client) Init() error {
	if len(c.Network) == 0 {
		return errors.New("no designated Network, please choose tcp/udp")
	}

	raddr, err := net.ResolveTCPAddr("tcp4", c.ServerAddr)

	if err != nil {
		fmt.Println("err setup client:", err)
		return err
	}

	var connErr error = nil
	c.Conn, connErr = net.DialTCP(c.Network, nil, raddr)
	if connErr != nil {
		fmt.Println("err setup client:", connErr)
		return connErr
	}

	return nil
}

func (c Client) Call(req CallRequest) (*CallResponse, error) {
	bytes, err := json.Marshal(req)
	if err != nil {
		fmt.Println("err when marshal data", err)
		return nil, err
	}

	data := wrapSendingData(bytes)
	_, err = c.Conn.Write(data)

	if err != nil {
		fmt.Println("err when writing data", err)
		return nil, err
	}

	//rBytes, err := ioutil.ReadAll(c.Conn)
	//
	//if err != nil {
	//	fmt.Println("err when read response data", err)
	//	return nil, err
	//}
	//
	//result := &CallResponse{}
	//
	//json.Unmarshal(rBytes, result)

	return nil, nil

}

const DATA_FLAG = "MINI"

func wrapSendingData(content []byte) []byte {
	clen := len(content)
	r := []byte{}
	flagBytes := []byte(DATA_FLAG)
	r = append(r, flagBytes...)
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(clen))
	fmt.Println("len of content: ", uint32(clen))
	fmt.Println("len of content: ", binary.BigEndian.Uint32(b))
	r = append(r, b...)
	r = append(r, content...)
	fmt.Println("all data", r)
	return r
}
