package utils

import (
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
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
	req, err := http.NewRequest("GET", url, nil)

	req.SetBasicAuth(login.Username, login.Password)
	req.Header.Set("Connection", "keep-alive")

	if err != nil {
		log.Fatalln(url, "Unable Create Request", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(url, "Unable to preform GET", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalln(url, "Response Code", resp.StatusCode, resp)
	}

	return resp
}

func MetricsRequest(login RouterLogin) io.ReadCloser {
	client := createClient(login)

	request := func(url string) *http.Response { return getRequest(client, login, url) }

	metricsUrl := login.Url + "/RST_stattbl.htm"

	baseUrlRequest := request(login.Url)
	defer baseUrlRequest.Body.Close()

	resp := request(metricsUrl)

	return resp.Body
}
