package config

import (
	"bsm/internal/config/structures/variables"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env     string
	Webhook variables.Webhook
}

func MustLoad() *Config {
	err := godotenv.Load("../../configs/env/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	env := getEnv("ENV")

	webhook := variables.WebhookInit(
		getEnv("WEBHOOK_ID"),
		getEnv("WEBHOOK_SECRET"),
		getEnv("WEBHOOK_DOMAIN"),
		getEnv("WEBHOOK_TOKEN"))
	cfg := &Config{
		Env:     env,
		Webhook: webhook,
	}

	return cfg
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Fatalf("No value in Env")
	return ""
}
