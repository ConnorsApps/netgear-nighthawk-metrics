package utils

import (
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
)

func MetricsRequest(login RouterLogin) io.ReadCloser {
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

	metricsUrl := login.Url + "/RST_stattbl.htm"

	req, err := http.NewRequest("GET", metricsUrl, nil)
	req.SetBasicAuth(login.Username, login.Password)

	if err != nil {
		log.Fatalln("Unable Create Request", err)
	}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln("Unable to preform GET", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalln("Response Code", resp.StatusCode, resp)
	}

	return resp.Body
}
