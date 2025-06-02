package gosol

import (
	"errors"
	"fmt"
)

var ErrRequireFalse = errors.New("require false")

func Requirer(condition bool, err error) error {
	if err == nil {
		return nil
	}
	if !condition {
		return fmt.Errorf("%s: %w", ErrRequireFalse, err)
	}
	return nil
}
