package kit

import (
	"errors"
	"reflect"
	"sync"
)

var (
	reflectValuePool sync.Pool
	InvalidFunction  = errors.New("the function is not a Func type")
)

func init() {
	reflectValuePool = sync.Pool{New: func() interface{} { return make([]reflect.Value, 0) }}
}

func CallFunction(function interface{}, args ...interface{}) (results []interface{}, err error) {

	// args
	params := reflectValuePool.Get().([]reflect.Value)
	for _, arg := range args {
		params = append(params, reflect.ValueOf(arg))
	}

	// call
	fn := reflect.ValueOf(function)
	if fn.Kind() != reflect.Func {
		return nil, InvalidFunction
	}
	values := fn.Call(params)

	// get result
	for _, value := range values {
		results = append(results, value.Interface())
	}

	// release the pool resource
	reflectValuePool.Put(params[:0])

	return
}
