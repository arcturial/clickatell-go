package clickatell

// SendResponseMessage struct represents the response of a message contained
// withing a "send" API call.
type SendResponseMessage struct {
    To          string          `json:"to"`
    MessageId   string          `json:"apiMessageId"`
    Error       ErrorResponse   `json:"error"`
}

// SendReponse struct represents the response of a "send"
// API call.
type SendResponse struct {
    Error ErrorResponse `json:"error"`

    Data struct {
        Message []SendResponseMessage `json:"message"`
    } `json:"data"`
}