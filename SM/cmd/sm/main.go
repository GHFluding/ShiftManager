package main

import (
	_ "github.com/GHFluding/ShiftManager/SM/docs"
	shiftManager "github.com/GHFluding/ShiftManager/SM/internal/app/SM"
)

// @title           Shift manager api
// @version         1.0
// @description     API for managing shifts and adding tasks
// @termsOfService  http://example.com/terms/

// @contact.name   API Support
// @contact.url    http://81.177.220.96/
// @contact.email  example@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api

// @schemes         http https
// @accept          json
// @produce         json
func main() {
	shiftManager.Run()
}
