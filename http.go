package clickatell

import (
    "github.com/google/go-querystring/query"
    "net/http"
    "strings"
    "strconv"
    "net/url"
    "fmt"
)

// HttpClient represents the Clickatell HTTP API
// and the credentials required to connect to it.
type HttpClient struct {
    // The net/http client
    client      *http.Client

    // API credentials
    apiId       int
    username    string
    password    string
}

// MakeQueryString() function will create a usable query string prepopulated
// with the HTTP API credentials.
func MakeQueryString(c *HttpClient, v url.Values) (url.Values) {

    // If no URL values have been set, we define an
    // empty value object
    if v == nil {
        v = url.Values{}
    }

    // Add API details from the HttpClient
    v.Add("api_id", fmt.Sprintf("%v", c.apiId))
    v.Add("user", c.username)
    v.Add("password", c.password)

    return v
}

// Http() function will create a new HttpClient based on the authentication
// credentials and HTTP connection provided. The http.Client is optional and will
// use a default client if not defined.
func Http(apiId int, username string, password string, client *http.Client) *HttpClient {

    // If the net/http client was not specified, a default
    // one shall be used instead
    if client == nil {
        client = http.DefaultClient
    }

    // Creates a new HTTP client
    return &HttpClient{
        apiId:      apiId,
        username:   username,
        password:   password,
        client:     client,
    }
}

// Send() function allows a user to send a message as defined
// by the Message object passed to it.
func (c *HttpClient) Send(m Message) (*SendResponse, error) {

    v, _ := query.Values(m)

    // Extend the request by adding the appropriate string parameters
    u := MakeQueryString(c, v)

    // Manually specify the destination to avoid bad query string encoding
    u.Add("to", strings.Join(m.Destination, ","))

    // Dispatch the API request
    req, _ := CreateRequest(Concat(apiEndpoint, "http/sendmsg", "?", u.Encode()))
    resp, err := c.client.Do(req);

    // Initialize the resulting structure before
    // adding the message responses
    result := &SendResponse{}

    if err == nil {

        // Populate the response schema with information unwrapped in the legacy API response
        for _, entry := range strings.Split(GetContent(resp), "\n") {

            // Unwrap the response
            unwrapped := UnwrapLegacyLine(entry)

            to := unwrapped.To

            // If the "To" entry does not exist and
            // we have an original destination. We will just use
            // that instead of relying on the API response
            if to == "" && len(m.Destination) > 0 {
                to = m.Destination[0]
            }

            // Push the resulting message into the message stack
            // so out response looks similar to the REST API
            result.Data.Message = append(result.Data.Message, SendResponseMessage{
                to,
                unwrapped.Data,
                ErrorResponse{unwrapped.Error, unwrapped.ErrorCode},
            })
        }
    }

    return result, err

}

// GetBalance() function returns the users balance as per
// Clickatell account status.
func (c *HttpClient) GetBalance() (*GetBalanceResponse, error) {

    // Extend the request by adding the appropriate string parameters
    u := MakeQueryString(c, nil)

    // Initialize the result
    result := &GetBalanceResponse{}

    // Dispatch the API request
    req, _ := CreateRequest(Concat(apiEndpoint, "http/getbalance", "?", u.Encode()))
    resp, err := c.client.Do(req);

    if err == nil {

        // Create an array of response objects with the error object nil
        content := GetContent(resp)
        unwrapped := UnwrapLegacyLine(content)
        err := unwrapped.GetError()

        if err == nil {

            // Convert the balance to float
            balance, err := strconv.ParseFloat(unwrapped.Data, 64)

            // If the balance could not be parse, this should be treated
            // as an API error. The API is suppose to return float
            if err == nil {
                result.Data.Balance = balance
            }
        }
    }

    // Return a blank response with the error variable populated
    return result, err
}