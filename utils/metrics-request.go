package utils

import (
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"time"
)

func createClient(login RouterLogin) *http.Client {
	redirectFunc := func(req *http.Request, via []*http.Request) error {
		req.SetBasicAuth(login.Username, login.Password)
		return nil
	}

	cookieJar, err := cookiejar.New(nil)

	if err != nil {
		log.Fatalln("Unable to create cookie jar", err)
	}

	client := &http.Client{
		Jar:           cookieJar,
		CheckRedirect: redirectFunc,
	}

	return client
}

func getRequest(client *http.Client, login RouterLogin, url string) *http.Response {
	req, reqErr := http.NewRequest("GET", url, nil)

	req.SetBasicAuth(login.Username, login.Password)
	req.Header.Set("Connection", "keep-alive")

	if reqErr != nil {
		log.Fatalln(url, "Unable Create Request", reqErr)
	}

	maxRetries := 7
	retries := 0

	var resp *http.Response
	var err error

	for retries < maxRetries {
		resp, err = client.Do(req)

		if err == nil && resp.StatusCode == http.StatusOK {
			break
		} else {
			time.Sleep(1 * time.Second)
		}
		retries++
	}
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Panicln(url, "Unable to make request", err, resp)
	}

	return resp
}

func MetricsRequest(login RouterLogin) io.ReadCloser {
	client := createClient(login)

	request := func(url string) *http.Response { return getRequest(client, login, url) }

	metricsUrl := login.Url + "/RST_stattbl.htm"

	resp := request(metricsUrl)

	return resp.Body
}
