package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {
	logger,err:=zap.NewProduction()
	if err !=nil{
		log.Println("🔴Error initializing Zap-logger:",err)
	}

	// flush out buffer before anything (Zap might contain buffer)
	defer logger.Sync()

	// Zap automatically logs in JSON, unlike Logrus
	logger.Info("This is an INFO. message ☑️")

	// Customization
	logger.Info("User Logged In ☑️", zap.String("username", "Skyy"), zap.String("method" ,"GET"))



}