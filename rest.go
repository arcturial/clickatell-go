package clickatell

import (
    "net/http"
    "encoding/json"
    "bytes"
)

// RestClient struct represents an instance of a REST API connection
// based on the API token provided.
type RestClient struct {
    // The net/http client
    client      *http.Client

    // API credentials
    apiToken    string
}

// Rest() funciton creates a new REST connection instance based on
// the token and http.Client. The http.Client is optional and will use
// a default implementation if nil.
func Rest(apiToken string, client *http.Client) *RestClient {

    // If the net/http client was not specified, a default
    // one shall be used instead
    if client == nil {
        client = http.DefaultClient
    }

    return &RestClient{
        client:     client,
        apiToken:   apiToken,
    }
}

// applyHeaders() function is a private function that will apply the
// required HTTP headers to a RESTful packet.
func (c *RestClient) applyHeaders(req *http.Request) *http.Request {

    req.Header.Add("User-Agent", userAgent)
    req.Header.Add("Authorization", Concat("Bearer", " ", c.apiToken))
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("X-Version", "1")
    req.Header.Add("Accept", "application/json")

    return req
}

// Send() function allows a user to send a message as defined
// by the Message object passed to it.
func (c *RestClient) Send(in Message) (*SendResponse, error) {

    jsonBody, _ := json.Marshal(in)

    // Dispatch the request to the API
    req, _ := http.NewRequest("POST", Concat(apiEndpoint, "rest/message"), bytes.NewBuffer(jsonBody))
    resp, err := c.client.Do(c.applyHeaders(req))

    result := &SendResponse{}

    if err == nil {

        err = json.NewDecoder(resp.Body).Decode(result)

        if err == nil {
            err = result.Error.GetError()
        }
    }

    return result, err
}

// GetBalance() function returns the users balance as per
// Clickatell account status.
func (c *RestClient) GetBalance() (*GetBalanceResponse, error) {

    // Dispatch the request to the API
    req, _ := http.NewRequest("GET", Concat(apiEndpoint, "rest/account/balance"), nil)
    resp, err := c.client.Do(c.applyHeaders(req))

    result := &GetBalanceResponse{}

    if err == nil {

        err = json.NewDecoder(resp.Body).Decode(result)

        if err == nil {
            err = result.Error.GetError()
        }
    }

    return result, err
}