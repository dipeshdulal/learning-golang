package memoization

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

type outExp struct {
	out    []reflect.Value
	expiry time.Time
}

func buildInStruct(ft reflect.Type) (reflect.Type, error) {
	if ft.NumIn() == 0 {
		return nil, errors.New("Must have at least one param")
	}
	var sf []reflect.StructField
	for i := 0; i < ft.NumIn(); i++ {
		ct := ft.In(i)
		if !ct.Comparable() {
			return nil, fmt.Errorf("parameters are not comparable")
		}
		sf = append(sf, reflect.StructField{
			Name: fmt.Sprintf("F%d", i),
			Type: ct,
		})
	}
	s := reflect.StructOf(sf)
	return s, nil
}

// Cacher caches, memoizes the function input and outputs
func Cacher(f interface{}, expiration time.Duration) (interface{}, error) {
	// only for functions
	funcType := reflect.TypeOf(f)
	if funcType.Kind() != reflect.Func {
		return nil, errors.New("caching is only for functions")
	}

	inType, err := buildInStruct(funcType)
	if err != nil {
		return nil, err
	}

	if funcType.NumOut() == 0 {
		return nil, errors.New("must have at least one returned value")
	}

	m := map[interface{}]outExp{}
	funcValue := reflect.ValueOf(f)

	cacher := reflect.MakeFunc(funcType, func(args []reflect.Value) []reflect.Value {
		iv := reflect.New(inType).Elem()

		for k, v := range args {
			iv.Field(k).Set(v)
		}

		ivv := iv.Interface() // input actual value rather than pointer
		ov, ok := m[ivv]
		now := time.Now()

		// if not found on map or already expired
		if !ok || ov.expiry.Before(now) {
			ov.out = funcValue.Call(args)
			ov.expiry = now.Add(expiration)
			m[ivv] = ov
		}
		return ov.out
	})

	return cacher.Interface(), nil
}
