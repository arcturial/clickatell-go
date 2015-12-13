# Clickatell GO Library

This package provides integration with the [Clickatell SMS gateway](https://www.clickatell.com).

## Installation

To acquire the library in your project, simply run:

```
    go get github.com/arcturial/clickatell-go
```

## Basic Usage

Sending messages relies on the `Message` object discussed further down. You can integrate either with the *HTTP* or *REST* API Clickatell provides. The API objects both provide the same
methods and can be used interchangeably.

```
package main

import (
    "log"
    "github.com/arcturial/clickatell-go"
)

func main() {

    http := clickatell.Http(apiId, "user", "password", nil)

    m := clickatell.Message{
        Destination:    []string{"2"},
        Body:           "Testing",
    }

    httpResp, err := http.Send(m)

    log.Println(err)
    log.Println(httpResp)

    rest := clickatell.Rest("apiToken", nil)

    restResp, err := rest.Send(m)

    log.Println(err)
    log.Println(restResp)
}
```

## Send Response

The `Send()` function will return a `SendResponse` object. The response objects are defined as `response_[method].go` files.

```
type SendResponse struct {
    Error ErrorResponse `json:"error"`

    Data struct {
        Message []SendResponseMessage `json:"message"`
    } `json:"data"`
}

type SendResponseMessage struct {
    To          string          `json:"to"`
    MessageId   string          `json:"apiMessageId"`
    Error       ErrorResponse   `json:"error"`
}
```

Interfacing with these response objects can be done as follows:

```
func main() {

    http := clickatell.Http(apiId, "user", "password", nil)

    m := clickatell.Message{
        Destination:    []string{"2"},
        Body:           "Testing",
    }

    httpResp, err := http.Send(m)

    for _, entry := range httpResp.Data.Message {
        log.Println(entry.Error)
        log.Println(entry.To)
        log.Println(entry.MessageId)
    }
}
```

## Message Object

The message struct represents a Clickatell message and maps to the appropriate API parameters.

```
type Message struct {
    Binary          bool        // Send the message as binary (optional)
    ClientMsgId     string      // Specify the clientMessageId (optional)
    Concat          int         // Max concatenation value (optional)
    DeliveryQueue   int         // The user queue priority. (optional)
    Destination     []string    // The destination value. (required)
    Callback        int         // The callback type. (optional)
    Escalate        bool        // Specify the escalate value. (optional)
    MaxCredits      int         // The maximum amount of credits to use. (optional)
    Body            string      // The message body (required)
    Mo              bool        // Wether to use MO or not. (optional)
    ScheduledTime   int         // Scheduled delivery time. (optional)
    From            string      // The sender ID or "from" address. (optional)
    Unicode         bool        // Should the message be sent with unicode. (optional)
    ValidityPeriod  int         // How long the message is valid if queued. (optional)
}
```

## Available Methods

Current available methods are:

``Send(m Message)``

``GetBalance()``

Note: These will be extended with future versions

## Contributing

The library is still in it's initial release and is missing some methods/parameters. Contributions to the available methods or the way the source code is structured or optimized are welcome.