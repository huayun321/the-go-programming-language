package main

import "fmt"

var es []*Employee

type Employee struct {
	ID int
	Name string
}

func EmployeeID(id int) *Employee {
	for _, e := range es {
		if e.ID == id {
			return e
		}
	}
	return nil
}

func init()  {
	es = append(es, &Employee{1, "a"})
	es = append(es, &Employee{2, "b"})
}

func main() {
	fmt.Println(EmployeeID(1))
	//如果返回值不是指针类型 则不可以链式赋值
	//EmployeeID(1).Name = "hello"
}
