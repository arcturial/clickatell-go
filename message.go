package clickatell

// The Message struct is a representation of a Clickatell SMS message. The properties
// of the message can be altered before using it to submit to the Send() method.
type Message struct {
    Binary          bool        `url:"binary,omitempty" json:"binary,omitempty"`
    ClientMsgId     string      `url:"cliMsgId,omitempty" json:"clientMessageId,omitempty"`
    Concat          int         `url:"concat,omitempty" json:"maxMessageParts,omitempty"`
    DeliveryTime    int         `url:"deliv_time,omitempty" json:"-,omitempty"`
    DeliveryQueue   int         `url:"queue,omitempty" json: "userPriorityQueue,omitempty"`
    Destination     []string    `url:"-" json:"to"`
    Callback        int         `url:"callback,omitempty" json:"callback,omitempty"`
    Escalate        bool        `url:"escalate,omitempty" json:"escalate,omitempty"`
    MaxCredits      int         `url:"max_credits,omitempty" json:"maxCredits,omitempty"`
    Body            string      `url:"text" json:"text"`
    Mo              bool        `url:"mo,omitempty" json:"mo,omitempty"`
    ScheduledTime   int         `url:"scheduled_time,omitempty" json:"scheduledDeliveryTime,omitempty"`
    From            string      `url:"from,omitempty" json:"from,omitempty"`
    Unicode         bool        `url:"unicode,omitempty" json:"unicode,omitempty"`
    ValidityPeriod  int         `url:"validity,omitempty" json:"validityPeriod,omitempty"`
}