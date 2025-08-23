package main

import "github.com/sirupsen/logrus"

func main() {
	log:= logrus.New()

	// set log level
	log.SetLevel(logrus.InfoLevel)

	// set log formatter
	log.SetFormatter(&logrus.JSONFormatter{})

	// logging examples
	// chaining them
	log.Info("This is an INFO. message ✅")
	log.Warn("This is a WARNING message ⚠️")
	log.Error("This is an ERROR message 🔴")

	log.WithFields(logrus.Fields{
		"username":"Skyy",
		"method":"GET", // can be others, just for example
	}).Info("User logged In ☑️")
}