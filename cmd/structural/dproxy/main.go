package main

import (
	"fmt"

	"github.com/marksartdev/go-patterns/pkg/structural/dproxy"
)

func main() {
	pr := dproxy.NewPersonBean()
	h := dproxy.NewOwnerInvocationHandler(pr)
	p := dproxy.NewProxyPersonBean(h)
	p.SetName("test Name")
	fmt.Println(p.GetName())
}
