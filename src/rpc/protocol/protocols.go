package protocol

import (
	"container/list"
	"fmt"
)

type Param struct {
	name string
	age  string
}

func (p Param) String() string {
	return fmt.Sprint("Name: ", p.name, ", age: ", p.age)
}

type CallRequest struct {
	Method string
	Args   []string
}

func (r CallRequest) String() string {
	return fmt.Sprint("Method: ", r.Method, ", args: ", r.Args)
}

type CallResponse struct {
	Results *list.List
}

func (r CallResponse) String() string {
	return fmt.Sprint("Results: ", r.Results)
}
