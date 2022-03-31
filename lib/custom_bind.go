package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

// CustomBind custom bind the data
func CustomBind(source *http.Request, dest interface{}) error {
	err := source.ParseMultipartForm(1000)
	if err != nil {
		return err
	}
	if source == nil {
		return nil
	}

	formSource := source.MultipartForm.Value

	destValue := reflect.ValueOf(dest)
	destType := reflect.TypeOf(dest)

	if destType.Kind() != reflect.Ptr {
		return errors.New("only pointers can be binded")
	}

	if destType.Elem().Kind() != reflect.Struct {
		return errors.New("only struct pointers are allowed")
	}

	for i := 0; i < destType.Elem().NumField(); i++ {
		currentField := destType.Elem().Field(i)
		fieldValue := destValue.Elem().Field(i)

		if currentField.Type == reflect.TypeOf(ModelBase{}) {
			if err := CustomBind(source, fieldValue.Addr().Interface()); err != nil {
				return err
			}
		}

		key := currentField.Tag.Get("form")
		if key == "" {
			continue
		}

		targetValue, ok := formSource[key]
		if !ok {
			continue
		}

		if len(targetValue) == 0 {
			continue
		}

		kind := currentField.Type.Kind()

		switch kind {
		case reflect.String:
			fieldValue.SetString(targetValue[0])
		case reflect.Bool:
			b, err := strconv.ParseBool(targetValue[0])
			if err != nil {
				return err
			}
			fieldValue.SetBool(b)
		case reflect.Int:
			i, err := strconv.Atoi(targetValue[0])
			if err != nil {
				return err
			}
			fieldValue.SetInt(int64(i))
		default:

			_, ok := fieldValue.Interface().(time.Time)
			if ok {
				val, _ := time.Parse("2006-01-02 15:04:05", targetValue[0])
				fieldValue.Set(reflect.ValueOf(val))
				continue
			}

			if fieldValue.CanInterface() && fieldValue.Type().NumMethod() > 0 {
				val, ok := fieldValue.Addr().Interface().(json.Unmarshaler)
				if !ok {
					return fmt.Errorf("data type %s doesn't implement unmarshaler interface", fieldValue.Type())
				}
				err := val.UnmarshalJSON([]byte(targetValue[0]))
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
