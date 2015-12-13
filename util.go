package clickatell

import (
    "bytes"
    "strings"
    "net/http"
    "io/ioutil"
)

// CreateRequest() function creates a new HTTP get request based on
// the URL passed to it. It will attach the user agent defined by
// the "userAgent" constant variable.
func CreateRequest(url string) (*http.Request, error) {

    req, err := http.NewRequest("GET", url, nil)

    if err != nil {
        // Attach user agent for tracking purposes
        req.Header.Add("User-Agent", userAgent)
    }

    return req, err
}

// GetContent() function will return the http.Response content
// as a string.
func GetContent(r *http.Response) string {
    defer r.Body.Close()

    // Read the body from the HTTP response object
    s, _ := ioutil.ReadAll(r.Body)

    return strings.Trim(string(s), "\n");
}

// Concat() function will concat all arguments into
// one string.
func Concat(args ...string) string {

    var buffer bytes.Buffer

    for _, arg := range args {
        buffer.WriteString(arg)
    }

    return buffer.String()
}