package main

import "fmt"

type ICallService interface {
	Call()
}
type IChildCallService interface {
	doCall()
}

func (this *abstractCallService) Call() {
	fmt.Println("parent call this func")
	this.doCall()
}

// 抽象类
type abstractCallService struct {
	IChildCallService
}

// 具体子类
type ConcreteCallServiceImpl struct {
	abstractCallService
}

func (*ConcreteCallServiceImpl) doCall() {
	fmt.Println("child call this func")
}

// 初始化方式:
func NewConcreteCallService() *ConcreteCallServiceImpl {
	serviceImpl := &ConcreteCallServiceImpl{}
	serviceImpl.IChildCallService = serviceImpl
	return serviceImpl
}
func main() {
	service := NewConcreteCallService()
	service.Call()
}
