package config

type WebhookB24 struct {
	clientID     string
	clientSecret string
	domain       string
	auth         string
	webhookURL   string
}

func WebhookB24Init(id, secret, domain, auth, url string) WebhookB24 {
	w := WebhookB24{
		clientID:     id,
		clientSecret: secret,
		domain:       domain,
		auth:         auth,
		webhookURL:   url,
	}
	return w
}
func (w *WebhookB24) GetSecret() string {
	return w.clientSecret
}
func (w *WebhookB24) GetID() string {
	return w.clientID
}
func (w *WebhookB24) GetDomain() string {
	return w.domain
}
func (w *WebhookB24) GetAuthToken() string {
	return w.auth
}
func (w *WebhookB24) GetURL() string {
	return w.webhookURL
}

type WebhookList struct {
	createShift    string
	shiftList      string
	createTask     string
	addShiftWorker string
}

func (w *WebhookList) GetCreateShiftUrl() string {
	return w.createShift
}
func (w *WebhookList) GetShiftListUrl() string {
	return w.shiftList
}
func (w *WebhookList) GetCreateTaskUrl() string {
	return w.createTask
}
func (w *WebhookList) GETAddShiftWorkerURL() string {
	return w.addShiftWorker
}
