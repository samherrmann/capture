package configuration

import "errors"

// HelpError is the error that is returned by the Load function of the -h or
// --help flag was invoked and no such flag is defined. The error includes the
// usage information.
type HelpError struct {
	usage string
}

func (e *HelpError) Error() string {
	return e.usage
}

// IsHelpError returns true of err is of type *HelpError and returns false
// otherwise.
func IsHelpError(err error) bool {
	var helpErr *HelpError
	return errors.As(err, &helpErr)
}
