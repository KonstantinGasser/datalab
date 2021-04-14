package binder

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const (
	// tagPattern matches any thing in the tag starting with <bind:"> and ending with <">
	tagPattern = `bind:".+"`
)

type ErrNoFullBind struct {
	invalid []string
}

func (err ErrNoFullBind) Error() string {
	return fmt.Sprintf("%v do not satisfy the bind conditions", err.invalid)
}

func MustBind(v interface{}) error {
	regexBind, err := regexp.Compile(tagPattern)
	if err != nil {
		return err
	}

	invalid := []string{}

	el := reflect.ValueOf(v).Elem()
	for i := 0; i < el.NumField(); i++ {
		field := el.Field(i)
		tag := el.Type().Field(i).Tag

		matches := regexBind.FindAllString(string(tag), -1)
		if len(matches) <= 0 {
			continue
		}
		ok := shoudlBind(matches[0])
		if !ok {
			continue
		}
		if isNil(field) {
			invalid = append(invalid, el.Type().Field(i).Name)
		}
	}
	if len(invalid) > 0 {
		return ErrNoFullBind{invalid: invalid}
	}
	return nil
}

func shoudlBind(tag string) bool {
	vs := strings.Split(tag, ":")
	if len(vs) < 2 {
		return false
	}
	re, _ := regexp.Compile(`"yes"`)
	return re.MatchString(vs[1])
}

func isNil(field reflect.Value) bool {
	switch field.Type().Kind() {
	case reflect.Uint:
		return field.Interface().(uint) == 0
	case reflect.Int8:
		return field.Interface().(int8) == 0
	case reflect.Uint8:
		return field.Interface().(uint8) == 0
	case reflect.Int16:
		return field.Interface().(int16) == 0
	case reflect.Uint16:
		return field.Interface().(uint16) == 0
	case reflect.Int32:
		return field.Interface().(int32) == 0
	case reflect.Uint32:
		return field.Interface().(uint32) == 0
	case reflect.Int64:
		return field.Interface().(int64) == 0
	case reflect.Uint64:
		return field.Interface().(uint64) == 0
	case reflect.Int:
		return field.Interface().(int) == 0
	case reflect.String:
		return field.Interface().(string) == ""
	case reflect.Slice:
		return field.Len() == 0
	case reflect.Ptr:
		return field.Interface() == nil
	}
	return false
}
