package errors

type ErrApi interface {
	Code() int32
	Info() string
	Error() string
}

type ErrAPI struct {
	Status int32
	Err    error
	Msg    string
}

func (err ErrAPI) Error() string { return err.Err.Error() }
func (err ErrAPI) Info() string  { return err.Msg }
func (err ErrAPI) Code() int32   { return err.Status }
