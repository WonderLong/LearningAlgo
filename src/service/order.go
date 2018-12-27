package service

import (
	"fmt"
)

type Order struct {
	Id     uint32
	Number string
}

func (o Order) String() string {
	return fmt.Sprint("id: ", o.Id, " number: ", o.Number)
}

type OrderService struct {
	//fields
}

func (o *OrderService) GetById(id int, repl *Order) error {
	repl.Id = 10
	repl.Number = "010001010"

	return nil
}

type Input struct {
	ScopeBegin int
	ScopeEnd   int
}

type Output struct {
	Pa     string
	Orders []Order
}

func (o *OrderService) GetList(input *Input, repl *Output) error {
	repl.Orders = append(repl.Orders, Order{1, "010001010"})
	repl.Orders = append(repl.Orders, Order{2, "010001011"})
	repl.Orders = append(repl.Orders, Order{3, "010001012"})
	repl.Orders = append(repl.Orders, Order{4, "010001013"})
	repl.Orders = append(repl.Orders, Order{5, "010001014"})
	repl.Pa = "aa"

	fmt.Println(repl)

	return nil
}
