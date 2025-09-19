package mailserver

type Payload struct {
	DataType string `json:"dataType"`
	Data     string `json:"data"`
}

type MailItem struct {
	Name      string   `json:"name"`
	To        string   `json:"to"`
	EmailType string   `json:"type"`
	TTL       string   `json:"ttl"`
	Payload   *Payload `json:"payload,omitempty"`
	//RedirectUrl string `json:"redirectUrl"`
	LinkUrl string `json:"linkUrl"`
	// Used to indicate if emails should be for web users or
	// geared towards API users who just want the tokens without
	// any links
	Mode string `json:"mode"`
}

type MailQueue interface {
	SendMail(mail *MailItem) error
}
