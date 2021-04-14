package not

import (
	"reflect"
	"strings"
)

// `not:"nil, max=12, min=6"` -> for string max/min len for int well obvious right?

// Void checks if all tag structs field are not the default values
// if the tag field specifies special boundaries they will be checked
func Void(v interface{}) bool {

	return false
}

// parse parses a field tag (`nil, max=12, min=6`)
// to => string:nil, [max]{12}
//					 [min]{ 6}
func parse(tag string) (string, []map[string]int) {

	parts := strings.Split(tag, ",")
	if len(parts) < 1 && len(parts) > 0 {
		return parts[0], nil
	}

	var opts map[string]int = make(map[string]int)
	if len(parts) > 1 && len(parts) < 3 {
		opt1 := strings.Split(parts[1], "=")
		opt2 := strings.Split(parts[2], "=")

	}
	return "", nil
}

// isNil checks if a field has a nil/default value
func isNil(field reflect.Value) bool {
	return false
}
