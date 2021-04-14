package required

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var (
	ErrBadSyntax  = errors.New("error: required: syntax for tag requires is wrong")
	ErrNotANumber = errors.New("error: required: option for tag is not a number")
)

// `not:"nil, max=12, min=6"` -> for string max/min len for int well obvious right?

// All checks if all tag structs field are not the default values
// if the tag field specifies special boundaries they will be checked
func All(v interface{}) error {
	_struct := reflect.ValueOf(v).Elem()
	fmt.Println(_struct.NumField())
	for i := 0; i < _struct.NumField(); i++ {
		// f := _struct.Field(i)
		tag := _struct.Type().Field(i).Tag
		tagParams, ok := tag.Lookup("required")
		if !ok {
			continue
		}
		actions := strings.Split(tagParams, ",")
		if len(actions) < 1 || actions[0] != "yes" {
			return ErrBadSyntax
		}
		opts, err := parse(actions[1:]...)
		if err != nil {
			return err
		}
		fmt.Println(opts["max"])
		// TODO: implement boundary checks if present
	}
	return nil
}

func isNil(field reflect.Value) bool {

	v := field.Type().Kind()
	switch v {
	case reflect.String:
		return field.Interface() == ""
	default:
		return false
	}
}

func parse(opts ...string) (map[string]int, error) {

	var vals map[string]int = make(map[string]int)
	for _, opt := range opts {
		v := strings.Split(opt, "=")
		if len(v) < 2 {
			return nil, ErrBadSyntax
		}
		i, err := strconv.Atoi(v[1])
		if err != nil {
			return nil, ErrNotANumber
		}
		vals[v[0]] = i
	}
	return vals, nil
}

// parse parses a field tag (`yes, max=12, min=6`)
// to => string:nil, [max]{12}
//					 [min]{ 6}
// func parse(tag string) (string, map[string]int, error) {

// 	parts := strings.Split(tag, ",")
// 	if len(parts) < 1 && len(parts) > 0 {
// 		return parts[0], nil, nil
// 	}

// 	var opts map[string]int = make(map[string]int)
// 	var opt1, opt2 []string
// 	if len(parts) > 1 && len(parts) < 3 {
// 		opt1 = strings.Split(parts[1], "=")
// 		if len(opt1) != 2 {
// 			return "", nil, ErrBadSyntax
// 		}
// 		opt2 = strings.Split(parts[2], "=")
// 		if len(opt2) != 2 {
// 			return "", nil, ErrBadSyntax
// 		}
// 	}
// 	v1, err := strconv.Atoi(opt1[1])
// 	if err != nil {
// 		return "", nil, ErrNotANumber
// 	}
// 	v2, err := strconv.Atoi(opt2[1])
// 	if err != nil {
// 		return "", nil, ErrNotANumber
// 	}

// 	opts[opt1[0]] = v1
// 	opts[opt2[0]] = v2

// 	return parts[0], opts, nil
// }

// isNil checks if a field has a nil/default value
