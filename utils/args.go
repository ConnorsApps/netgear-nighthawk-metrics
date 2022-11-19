package utils

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

const DEFAULT_USERNAME = "admin"
const DEFAULT_URL = "http://www.routerlogin.com"

type RouterLogin struct {
	Url                 string
	Username            string
	Password            string
	LoginTimeoutSeconds int
}

func ParseArgs() RouterLogin {
	// Environment Variables
	env_url := os.Getenv("NETGEAR_URL")
	env_password := os.Getenv("NETGEAR_PASSWORD")
	env_username := os.Getenv("NETGEAR_USERNAME")
	login_timeout := os.Getenv("LOGIN_TIMEOUT_SECONDS")

	// CMD flags
	url := flag.String("url", env_url, "Router Url")
	username := flag.String("username", env_username, "Router Username")
	timeout := flag.String("timeout", login_timeout, "Login Timeout Seconds")

	flag.Parse()

	if env_password == "" {
		log.Panicln("Router Password is required")
	}

	var router_username string
	var router_url string
	var router_login_timeout int

	// Set defaults
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

	if timeout == nil || *timeout == "" {
		router_login_timeout = 2
	} else {
		timeout_parsed, err := strconv.Atoi(*timeout)
		if err != nil {
			log.Fatalln("Unable to parse timeout to int", *timeout)
		}
		router_login_timeout = timeout_parsed
	}

	return RouterLogin{
		Url:                 router_url,
		Username:            router_username,
		Password:            env_password,
		LoginTimeoutSeconds: router_login_timeout,
	}
}
