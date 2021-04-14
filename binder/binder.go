package binder

import (
	"errors"
	"reflect"
	"regexp"
	"strings"
)

const (
	// tagPattern matches any thing in the tag starting with <bind:"> and ending with <">
	tagPattern = `bind:".+"`
)

var (
	ErrMissingValues = errors.New("not all required fields are set")
)

func MustBind(v interface{}) error {
	regexBind, err := regexp.Compile(tagPattern)
	if err != nil {
		return err
	}

	el := reflect.ValueOf(v).Elem()
	for i := 0; i < el.NumField(); i++ {
		field := el.Field(i)
		tag := el.Type().Field(i).Tag

		matches := regexBind.FindAllString(string(tag), -1)
		if len(matches) <= 0 {
			continue
		}
		yes, err := shoudlBind(matches[0])
		if err != nil {
			return err
		}
		if !yes {
			continue
		}
		if isNil(field) {
			return ErrMissingValues
		}
	}
	return nil
}

func shoudlBind(tag string) (bool, error) {
	vs := strings.Split(tag, ":")
	if len(vs) < 2 {
		return false, nil
	}
	re, err := regexp.Compile(`"yes"`)
	if err != nil {
		return false, err
	}
	return re.MatchString(vs[1]), nil
}

func isNil(field reflect.Value) bool {
	switch field.Type().Kind() {
	case reflect.String:
		return field.Interface().(string) == ""
	case reflect.Slice:
		return len(field.Interface().([]interface{})) > 0
	case reflect.Ptr:
		return field.Interface() == nil
	}
	return false
}
