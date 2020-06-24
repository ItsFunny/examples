package common

import "github.com/pkg/errors"

var (
	NilPointerError    = errors.New("Nil value")
	NoSuchElementError = errors.New("No Such Element")
	IndexOutOfRangeError=errors.New("Index out of range" )
	EmptyListError=errors.New("Empty List error")
)
