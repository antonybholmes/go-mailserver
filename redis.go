package mailer

const (
	REDIS_EMAIL_CHANNEL = "email-channel"
)

const (
	REDIS_EMAIL_TYPE_PASSWORDLESS = "passwordless"
)

type RedisQueueEmail struct {
	Name        string `json:"name"`
	To          string `json:"to"`
	EmailType   string `json:"type"`
	Ttl         string `json:"ttl"`
	Token       string `json:"token"`
	CallBackUrl string `json:"callbackurl"`
	VisitUrl    string `json:"visiturl"`
}
