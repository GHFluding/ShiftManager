package variables

type Webhook struct {
	clientID     string
	clientSecret string
	domain       string
	auth         string
}

func WebhookInit(id, secret, domain, auth string) Webhook {
	w := Webhook{
		clientID:     id,
		clientSecret: secret,
		domain:       domain,
		auth:         auth,
	}
	return w
}
func (w *Webhook) GetSecret() string {
	return w.clientSecret
}
func (w *Webhook) GetID() string {
	return w.clientID
}
func (w *Webhook) GetDomain() string {
	return w.domain
}
func (w *Webhook) GetAuthToken() string {
	return w.auth
}
