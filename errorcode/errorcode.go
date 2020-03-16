package errorcode

import (
	"github.com/pkg/errors"
	"strconv"
)

// Codes ecode error interface which has a code & message.
type Codes interface {
	// sometimes Error return Code in string form
	// NOTE: don't use Error in monitor report even it also work for now
	Error() string
	// Code get error code.
	Code() int
	// Message get code message.
	Message() string
	//Detail get error detail,it may be nil.
	Details() []interface{}
	// Equal for compatible.
	// Deprecated: please use ecode.EqualError.
	Equal(error) bool
}

func (e Code) Error() string {
	return strconv.FormatInt(int64(e), 10)
}

// Code return error code
func (e Code) Code() int { return int(e) }

// Message return error message
func (e Code) Message() string {
	if cm, ok := _codes[e.Code()]; ok {
		return cm
	}
	return e.Error()
}

// Details return details.
func (e Code) Details() []interface{} { return nil }

// Equal for compatible.
// Deprecated: please use ecode.EqualError.
func (e Code) Equal(err error) bool { return EqualError(e, err) }

type Code int

var (
	_codes             = map[int]string{}
	CodeMustBeThanZero = add(-1, "business error code must greater than zero")
	CodeNotExist       = add(-2, "error code not exist")
)

// Register register ecode message map.
func Register(e int, msg string) {
	_codes[e] = msg
}

// New new a ecode.Codes by int value.
// NOTE: ecode must unique in global, the New will check repeat and then panic.
func New(e int, msg string) Code {
	if e <= 0 {
		return CodeMustBeThanZero
	}
	return add(e, msg)
}

func add(e int, msg string) Code {
	if _, ok := _codes[e]; ok {
		// return self code if error code already exist
		return Int(e)
	}
	// register err message into message struct
	Register(e, msg)

	return Int(e)
}

// Int parse code int to error.
func Int(i int) Code { return Code(i) }

// String parse code string to error.
func String(e string) Code {
	if e == "" {
		return OK
	}
	// try error string
	i, err := strconv.Atoi(e)
	if err != nil {
		return ServerErr
	}
	return Code(i)
}

// Cause cause from error to ecode.
func Cause(e error) Codes {
	if e == nil {
		return OK
	}
	ec, ok := errors.Cause(e).(Codes)
	if ok {
		return ec
	}
	return String(e.Error())
}

// Equal equal a and b by code int.
func Equal(a, b Codes) bool {
	if a == nil {
		a = OK
	}
	if b == nil {
		b = OK
	}
	return a.Code() == b.Code()
}

// EqualError equal error
func EqualError(code Codes, err error) bool {
	return Cause(err).Code() == code.Code()
}
