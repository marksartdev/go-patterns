package dproxy

import (
	"reflect"
	"strings"

	"github.com/marksartdev/go-patterns/pkg/common"
)

// InvocationHandler Интерфейс обработчика вызовов.
type InvocationHandler interface {
	invoke(method string, args ...interface{}) (reflect.Value, error)
}

type invocationHandler struct {
	person PersonBean
}

func (i invocationHandler) invokeReflect(method string, args ...interface{}) reflect.Value {
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
	invocationHandler
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

// NewOwnerInvocationHandler Создать обработчик для владельца анкеты.
func NewOwnerInvocationHandler(person PersonBean) InvocationHandler {
	h := ownerInvocationHandler{}
	h.person = person

	return h
}

type nonOwnerInvocationHandler struct {
	invocationHandler
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

// NewNonOwnerInvocationHandler Создать обработчик для не владельца анкеты.
func NewNonOwnerInvocationHandler(person PersonBean) InvocationHandler {
	h := nonOwnerInvocationHandler{}
	h.person = person

	return h
}
