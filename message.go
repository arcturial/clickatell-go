package clickatell

// The Message struct is a representation of a Clickatell SMS message. The properties
// of the message can be altered before using it to submit to the Send() method.
type Message struct {
    Binary          bool        `url:"binary,omitempty" json:"binary"`
    ClientMsgId     string      `url:"cliMsgId,omitempty" json:"clientMessageId"`
    Concat          int         `url:"concat,omitempty" json:"maxMessageParts"`
    DeliveryTime    int         `url:"deliv_time,omitempty" json:"-"`
    DeliveryQueue   int         `url:"queue,omitempty" json: "userPriorityQueue"`
    Destination     []string    `url:"-" json:"to"`
    Callback        int         `url:"callback,omitempty" json:"callback"`
    Escalate        bool        `url:"escalate,omitempty" json:"escalate"`
    MaxCredits      int         `url:"max_credits,omitempty" json:"maxCredits"`
    Body            string      `url:"text" json:"text"`
    Mo              bool        `url:"mo,omitempty" json:"mo"`
    ScheduledTime   int         `url:"scheduled_time,omitempty" json:"scheduledDeliveryTime"`
    From            string      `url:"from,omitempty" json:"from"`
    Unicode         bool        `url:"unicode,omitempty" json:"unicode"`
    ValidityPeriod  int         `url:"validity,omitempty" json:"validityPeriod"`
}