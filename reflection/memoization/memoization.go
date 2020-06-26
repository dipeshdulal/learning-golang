package memoization

import (
	"errors"
	"reflect"
	"time"
)

// Cacher caches, memoizes the function input and outputs
func Cacher(f interface{}, expiration time.Duration) (interface{}, error) {
	// only for functions
	funcType := reflect.TypeOf(f)
	if funcType.Kind() != reflect.Func {
		return nil, errors.New("caching is only for functions")
	}
	return f, nil
}
