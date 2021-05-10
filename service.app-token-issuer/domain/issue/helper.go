package issue

import "time"

func override(exp time.Time) bool {
	return exp.Unix() <= time.Now().Unix()
}
