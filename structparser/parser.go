package structparser

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetShouldBindToStruct binds the get request params into struct
func GetShouldBindToStruct(c *gin.Context, dataHolder interface{}) error {
	var err error
	dhType := reflect.TypeOf(dataHolder)

	if err := typeShouldBeStruct(dhType); err != nil {
		return err
	}

	dhVal := reflect.ValueOf(dataHolder)
	for i := 0; i < dhType.Elem().NumField(); i++ {

		field := dhType.Elem().Field(i)
		key := field.Tag.Get("getParser")
		kind := field.Type.Kind()

		getVal := c.Query(key)

		fbn := dhVal.Elem().Field(i)

		switch kind {

		case reflect.Int:
			intVal, err := strconv.ParseInt(getVal, 10, 64)
			if err != nil {
				return err
			}
			if fbn.CanSet() {
				fbn.SetInt(intVal)
			}
			break

		case reflect.String:
			if fbn.CanSet() {
				fbn.SetString(getVal)
			}
			break

		case reflect.Struct:
			if fbn.CanSet() {
				val := reflect.New(field.Type)
				err := json.Unmarshal([]byte(getVal), val.Interface())
				if err != nil {
					return err
				}
				fbn.Set(val.Elem())
			}
			break
		}

	}

	return err
}

func typeShouldBeStruct(refType reflect.Type) error {
	var err error
	if refType.Elem().Kind() != reflect.Struct {
		return errors.New("input should be struct type")
	}
	return err
}
