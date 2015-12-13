package clickatell

import (
    "fmt"
    "errors"
)

// The ClickatellError struct is an extension of
// normal golang errors that allows us to add a Clickatell
// API error code to error objects
type ClickatellError struct {
    error
    Code string
}

// Error() function formats the error object to include
// the API error code as part of the error string
func (e *ClickatellError) Error() string {
    return fmt.Sprintf("%s - %s", e.error, e.Code)
}

// MakeError takes an ErrorResponse object from the API and turns
// it into a golang error
func MakeError(err ErrorResponse) *ClickatellError {
    return &ClickatellError{errors.New(err.Description), err.Code}
}
