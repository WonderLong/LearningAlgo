package main

import (
	"../rpc/protocol"
	"bufio"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func main() {
	errorss()
}

type RuntimeError struct {
	Err error
}

func (re RuntimeError) Error() string {
	return re.Err.Error()
}

func throwError() *RuntimeError {
	return &RuntimeError{
		Err: errors.New("runtime error....."),
	}
}

func errorss() {
	if e := throwError(); e != nil {
		fmt.Println(e.Error())
	}
}

func concurrency() {
	c := make(chan int)

	go func() {
		d, ok := <-c
		fmt.Println("w data: ", d, "   ", ok)
		for j := range c {
			fmt.Println(j)
		}
	}()

	for i := 1; i < 10; i++ {
		c <- i
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}

func reflection() {

}

func test_json_marshal() {
	b := []byte{77, 73, 78, 73, 0, 0, 0, 32, 123, 34, 77, 101, 116, 104, 111, 100, 34, 58, 34, 97, 97, 97, 34, 44, 34, 65, 114, 103, 115, 34, 58, 91, 34, 116, 116, 116, 116, 34, 93, 125}
	fmt.Println(string(b[0:4]))
	fmt.Println(binary.BigEndian.Uint32(b[4:8]))
	d := &protocol.CallRequest{}
	json.Unmarshal(b[8:], d)
	fmt.Println(d)
}

func test_marshal() {
	i := Input{"zhang", 20, []Thing{{"a", 1, "aa"}, {"b", 2, "bb"}}}
	fmt.Println(i.toString())
	data, _ := json.Marshal(i)
	var r Input
	if e := json.Unmarshal(data, &r); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(r.toString())
	}
}

type Thing struct {
	What string
	When int
	How  string
}

func (in Thing) String() string {
	return fmt.Sprint("what: ", in.What, ", When: ", in.When, ", How: ", in.How)
}

type Input struct {
	Name   string
	Age    int
	Things []Thing
}

func (in Input) toString() string {
	return fmt.Sprint("name: ", in.Name, ", age: ", in.Age, ", things: ", in.Things)
}

func testImplementationTypeMethod() {
	a := &TT{"aaa"}
	a.show()
	a.update("bbb")
	a.show()

	a.tryUpdate("ccc")
	a.show()
}

type TT struct {
	a string
}

func (t TT) show() {
	fmt.Println("the value", t.a)
}

func (t *TT) update(n string) {
	t.a = n
}

func (t TT) tryUpdate(n string) {
	t.a = n
}
