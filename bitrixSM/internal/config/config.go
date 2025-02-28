package config

type Config struct {
	Port            string `env:"PORT" envDefault:"8080"`
	BitrixWebhook   string `env:"BITRIX_WEBHOOK_TOKEN"`
	BitrixPortalURL string `env:"BITRIX_PORTAL_URL"`
	MonolithURL     string `env:"MONOLITH_URL"`
	APIKey          string `env:"API_KEY"`
}
