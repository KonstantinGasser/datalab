package errors

type Api interface {
	Code() int32
	Info() string
	Error() string
}

type API struct {
	Status int32
	Err    error
	Msg    string
}

func New(code int32, err error, msg string) API {
	return API{
		Status: code,
		Msg:    msg,
		Err:    err,
	}
}

func (err API) Error() string {
	if err.Err == nil {
		return "nil error!"
	}
	return err.Err.Error()
}
func (err API) Info() string { return err.Msg }
func (err API) Code() int32  { return err.Status }
