package utils

import (
	"flag"
	"log"
	"os"
	"strings"
)

const DEFAULT_USERNAME = "admin"
const DEFAULT_PORT = "8080"
const DEFAULT_URL = "http://www.routerlogin.com"

type AppArgs struct {
	Url      string
	Username string
	Password string
	Port     string
}

func ParseArgs() AppArgs {
	// Environment Variables
	env_url := os.Getenv("ROUTER_URL")
	env_password := os.Getenv("ROUTER_PASSWORD")
	env_username := os.Getenv("ROUTER_USERNAME")
	env_port := os.Getenv("PORT")

	// CMD flags
	url := *flag.String("url", env_url, "Router Url")
	username := *flag.String("username", env_username, "Router Username")
	port := *flag.String("port", env_port, "App Port")

	flag.Parse()

	if env_password == "" {
		log.Panicln("Router Password is required")
	}

	// Set defaults
	if username == "" {
		username = DEFAULT_USERNAME
	}

	if url == "" {
		url = DEFAULT_URL
	} else {
		url = strings.TrimSuffix(url, "/")
	}

	if port == "" {
		port = DEFAULT_PORT
	} else {
		port = strings.TrimSpace(port)
	}

	return AppArgs{
		Url:      url,
		Username: username,
		Password: env_password,
		Port:     port,
	}
}
