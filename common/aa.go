package common

import "fmt"

func NewLoginer() Loginer{
	return defaultLogin(5)
}
type Loginer interface {
	Login()
}
type defaultLogin string

func (d defaultLogin) Login(){
	fmt.Println("login in...")
}
