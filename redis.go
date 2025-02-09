package mailer

const (
	REDIS_EMAIL_CHANNEL = "email-channel"
)

const (
	REDIS_EMAIL_TYPE_VERIFY           = "verify"
	REDIS_EMAIL_TYPE_VERIFIED         = "verified"
	REDIS_EMAIL_TYPE_PASSWORDLESS     = "passwordless"
	REDIS_EMAIL_TYPE_PASSWORD_RESET   = "password-reset"
	REDIS_EMAIL_TYPE_PASSWORD_UPDATED = "password-updated"
	REDIS_EMAIL_TYPE_EMAIL_RESET      = "email-reset"
	REDIS_EMAIL_TYPE_EMAIL_UPDATED    = "email-updated"
	REDIS_EMAIL_TYPE_ACCOUNT_CREATED  = "account-created"
	REDIS_EMAIL_TYPE_ACCOUNT_UPDATED  = "account-updated"
)

type RedisQueueEmail struct {
	Name      string `json:"name"`
	To        string `json:"to"`
	EmailType string `json:"type"`
	Ttl       string `json:"ttl"`
	Token     string `json:"token"`
	//RedirectUrl string `json:"redirectUrl"`
	LinkUrl string `json:"linkUrl"`
}
