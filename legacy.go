package clickatell

import (
    "errors"
    "regexp"
    "strings"
)

// LegacyUnwrap struct represents the result
// that was parsed from the Clickatell HTTP API. The API is older
// and does not always return consistent results.
type LegacyUnwrap struct {
    Error       string
    ErrorCode   string
    Data        string
    To          string
}

// HasError() function will check if the unwrapped response was
// successful or not.
func (l *LegacyUnwrap) HasError() bool {
    return l.Error != ""
}

// GetError() function will return a golang error object (ClickatellError)
// if the response was unsuccessful.
func (l *LegacyUnwrap) GetError() *ClickatellError {

    if l.HasError() {
        return &ClickatellError{errors.New(l.Error), l.ErrorCode}
    } else {
        return nil
    }
}

// Since the HTTP API contains some inconsistent results, the UnwrapLegacyLine() function
// will parse these responses and return them as a struct with "known" properties. This function
// is used to create a bridge between the legacy response and a predictable response object.
func UnwrapLegacyLine(line string) *LegacyUnwrap {

    // Configure regex expressions
    toExp := regexp.MustCompile("To:\\s(.*)")
    errExp := regexp.MustCompile("ERR:\\s([0-9]+),(.*)")
    exp := regexp.MustCompile("([A-Za-z]+):\\s(.*)")

    // Pull the To: field from the line
    toMatch := toExp.FindStringSubmatch(line)
    line = toExp.ReplaceAllString(line, "")
    to := ""

    // Fetch the To: address from the API response if it's applicable
    if len(toMatch) > 0 {
        to = strings.TrimSpace(toMatch[1])
    }

    // Check if the request is an error
    errMatch := errExp.FindStringSubmatch(line)

    // If the request was an error, there's no need to run
    // more regex lookups
    if len(errMatch) > 0 {
        // Unpack error values
        code := strings.TrimSpace(errMatch[1])
        error := strings.TrimSpace(errMatch[2])

        return &LegacyUnwrap{Error: error, ErrorCode: code, Data: "", To: to}
    } else {
        // Unpack data values
        dataMatch := exp.FindStringSubmatch(line)
        data := strings.TrimSpace(dataMatch[2])

        return &LegacyUnwrap{Error: "", ErrorCode: "", Data: data, To: to}
    }
}