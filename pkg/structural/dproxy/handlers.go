package dproxy

import (
	"reflect"
	"strings"

	"github.com/marksartdev/go-patterns/pkg/common"
)

// Интерфейс обработчика вызовов.
type invocationHandler interface {
	invoke(method string, args ...interface{}) (reflect.Value, error)
}

type baseInvocationHandler struct {
	person PersonBean
}

func (i baseInvocationHandler) invokeReflect(method string, args ...interface{}) reflect.Value {
	in := make([]reflect.Value, len(args))
	for i := range args {
		in[i] = reflect.ValueOf(args[i])
	}

	v := reflect.ValueOf(i.person)
	f := v.MethodByName(method)
	result := f.Call(in)

	var resp reflect.Value
	if len(result) > 0 {
		resp = result[0]
	}

	return resp
}

type ownerInvocationHandler struct {
	baseInvocationHandler
}

func (o ownerInvocationHandler) invoke(method string, args ...interface{}) (reflect.Value, error) {
	switch {
	case strings.HasPrefix(method, "Get"):
		return o.invokeReflect(method, args...), nil
	case method == "SetHotOrNotRating":
		return reflect.Value{}, common.IllegalAccessError{}
	case strings.HasPrefix(method, "Set"):
		return o.invokeReflect(method, args...), nil
	default:
		return reflect.Value{}, nil
	}
}

// Создать обработчик для владельца анкеты.
func newOwnerInvocationHandler(person PersonBean) invocationHandler {
	h := ownerInvocationHandler{}
	h.person = person

	return h
}

type nonOwnerInvocationHandler struct {
	baseInvocationHandler
}

func (n nonOwnerInvocationHandler) invoke(method string, args ...interface{}) (reflect.Value, error) {
	switch {
	case strings.HasPrefix(method, "Get"):
		return n.invokeReflect(method, args...), nil
	case method == "SetHotOrNotRating":
		return n.invokeReflect(method, args...), nil
	case strings.HasPrefix(method, "Set"):
		return reflect.Value{}, common.IllegalAccessError{}
	default:
		return reflect.Value{}, nil
	}
}

// Создать обработчик для не владельца анкеты.
func newNonOwnerInvocationHandler(person PersonBean) invocationHandler {
	h := nonOwnerInvocationHandler{}
	h.person = person

	return h
}
