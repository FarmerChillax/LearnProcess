package main

import "fmt"

type Inter interface {
	Ping()
	Pang()
}

type Anter interface {
	Inter
	String()
}

type St struct {
	Name string
}

func (St) Ping() {
	fmt.Println("ping")
}

func (*St) Pang() {
	fmt.Println("pang")
}

func main() {
	st := &St{"farmer"}
	var i interface{} = st

	// 判断 i 绑定的实例是否实现了接口类型 Inter
	o := i.(Inter)
	o.Ping()
	o.Pang()

	// 下面的语句会引发 panic , 因为 i 没有实现接口 Anter
	// p := i.(Anter)
	// p.String()

	// 判断 i 绑定的实例是否就是具体类型 st
	s := i.(*St)
	fmt.Printf("%s\n", s.Name)

	// 通过 comma, ok 表达式模式进行判断

	// 判断 i 绑定的实例是否实现了接口类型 Inter
	if o, ok := i.(Inter); ok {
		o.Ping()
		o.Pang()
	}

	if p, ok := i.(Anter); ok {
		// i 没有实现接口 Anter, 所以程序不会执行这里
		p.String()
	}

	// 判断 i 绑定的实例是否就是具体类型 St
	if s, ok := i.(*St); ok {
		fmt.Println(s.Name)
	}
}
