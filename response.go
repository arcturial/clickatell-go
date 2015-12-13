package clickatell

import (
    "errors"
)

// The Clickatell API can return certain errors and these errors are represented
// by an ErrorResponse object.
type ErrorResponse struct {
    Description string `json:"description"`
    Code        string `json:"code"`
}

// HasError() function will check if an error response has been populated or not.
func (e *ErrorResponse) HasError() bool {
    return e.Description != ""
}

// GetError() function will return a golang error if the ErrorResponse object
// has been populated.
func (e *ErrorResponse) GetError() *ClickatellError {

    if e.HasError() {
        return &ClickatellError{errors.New(e.Description), e.Code}
    } else {
        return nil
    }
}