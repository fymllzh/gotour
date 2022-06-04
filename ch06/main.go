package main

import (
	"fmt"
	"gotour/fmtss"
)

func main()  {
	p := persion{
		age:30,
		name:"zhihao",
		address: addr{
			city:"河南",
		},

	}
	fmt.Println(p.address.aa,p.address.city)
	printString(p)
	p1 :=NewPersion("英美")
	fmt.Println(p1)

}

type persion struct {
	name string
	age uint
	address addr
}

type addr struct {
	city string
	aa int
}
//
func (p persion) Str() string  {
	return fmt.Sprintf("name is %s,age is %d",p.name,p.age)
}
func printString(s fmtss.Stringer)  {
	fmt.Println(s.Str())

}

func NewPersion(name string)persion{
	return persion{name: name}
}
