package utils

import (
	"flag"
	"log"
	"os"
	"strings"
)

const DEFAULT_USERNAME = "admin"
const DEFAULT_URL = "http://www.routerlogin.com"

type RouterLogin struct {
	Url      string
	Username string
	Password string
}

func ParseArgs() RouterLogin {
	// Environment Variables
	env_url := os.Getenv("NETGEAR_URL")
	env_password := os.Getenv("NETGEAR_PASSWORD")
	env_username := os.Getenv("NETGEAR_USERNAME")

	// CMD flags
	url := flag.String("url", env_url, "Router Url")
	username := flag.String("username", env_username, "Router Username")

	flag.Parse()

	if env_password == "" {
		log.Panicln("Router Password is required")
	}

	var router_username string
	var router_url string

	if username == nil || *username == "" {
		router_username = DEFAULT_USERNAME
	} else {
		router_username = *username
	}
	if url == nil || *url == "" {
		router_url = DEFAULT_URL
	} else {
		router_url = strings.TrimSuffix(*url, "/")
	}
	return RouterLogin{
		Url:      router_url,
		Username: router_username,
		Password: env_password,
	}
}
