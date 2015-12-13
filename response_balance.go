package clickatell

// GetBalanceResponse struct represents the response of a "getbalance"
// API call.
type GetBalanceResponse struct {
    Error ErrorResponse `json:"error"`

    Data struct {
        Balance float64 `json:"balance,string"`
    } `json:"data"`
}