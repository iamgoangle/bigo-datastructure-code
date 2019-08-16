package main

import (
	"fmt"
)

func main() {
	proxy := NewProxy()
	proxy.GetEmployee()
}

type realService struct{}

func (rs *realService) GetEmployee() {
	fmt.Println("Return from real service")
}

type Employee interface {
	GetEmployee()
}

type Proxy struct {
	Emp *realService
}

func NewProxy() Employee {
	realService := &realService{}
	return &Proxy{
		Emp: realService,
	}
}

func (p *Proxy) GetEmployee() {
	p.Emp.GetEmployee()
}
