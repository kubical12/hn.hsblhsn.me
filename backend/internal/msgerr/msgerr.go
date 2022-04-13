package msgerr

type Error struct {
	err error
	msg string
}

var _ error = (*Error)(nil)

func New(err error, msg string) *Error {
	return &Error{err, msg}
}

func (e *Error) Error() string {
	return e.err.Error()
}

func (e *Error) Msg() string {
	return e.msg
}
