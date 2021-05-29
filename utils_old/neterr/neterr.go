package neterr

import "fmt"

// ReqErr contain all information required for an api
// to return a usefull response to the client
type ReqErr struct {
	// Status of the call http.Status*
	Status int
	// Msg is an optional message given more detailed
	// information on failture (if http.StatusBadeRequest -> tell me what is wrong)
	Msg string
	// Err is optional if Status is >= http.StatusInternalServerError
	Err error
}

func (err ReqErr) Error() string {
	return fmt.Sprintf("stauts: %d, msg: %s, err: %s", err.Status, err.Msg, err.Err.Error())
}
