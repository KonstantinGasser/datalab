package apptoken

import (
	"time"
)

func (apt apptoken) canBeOverriden(storedExp time.Time) bool {
	currDate := time.Now().Unix()
	return currDate >= storedExp.Unix()
}
