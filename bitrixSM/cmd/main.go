package main

import (
	apilogic "bsm/internal/services/apiLogic"
	"log"

	"github.com/gin-gonic/gin"
	goBX24 "github.com/whatcrm/go-bitrix24"
)

func main() {
	// New API client
	const (
		clientID     = "example-client-id"
		clientSecret = "example-client-secret"
		domain       = "example-bitrix24-domain"
		auth         = "example-auth-token"
	)

	b24 := goBX24.NewAPI(clientID, clientSecret)
	if err := b24.SetOptions(domain, auth, true); err != nil {
		log.Fatalf("Setting API error: %v", err)
	}

	r := gin.Default()
	// Handling incoming message
	r.POST("/webhook", apilogic.HandleMessage)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Run server error: %v", err)
	}
}
