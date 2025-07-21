package utils

import "fmt"

func AppendError(existError, newErr error) error {
	if existError == nil {
		return newErr
	}
	return fmt.Errorf("%v, %v", existError, newErr)
}
