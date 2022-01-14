package env

import (
	"log"
	"os"
)

type envVars struct {
	Port string
}

// Var envVars
var Var envVars

// Setup func
// load env vars from system
func Setup() {

	var ok bool

	if Var.Port, ok = os.LookupEnv("PORT"); !ok {
		log.Fatalln("Env PORT needs to be specified")
	}
}
