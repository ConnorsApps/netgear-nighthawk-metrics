package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func parseDoc(response io.ReadCloser) (string, *goquery.Document) {
	defer response.Close()

	bodyBytes, err := io.ReadAll(response)
	if err != nil {
		log.Panicln(err)
	}

	bodyString := string(bodyBytes)
	if bodyString == "" {
		log.Panicln("No body found")
	}
	parsedHtmlReader := strings.NewReader(bodyString)

	doc, err := goquery.NewDocumentFromReader(parsedHtmlReader)
	if err != nil {
		log.Panicln(err)
	}
	return bodyString, doc
}

func createClient(appArgs AppArgs) *http.Client {
	redirectFunc := func(req *http.Request, _ []*http.Request) error {
		req.SetBasicAuth(appArgs.Username, appArgs.Password)
		return nil
	}

	cookieJar, err := cookiejar.New(nil)

	if err != nil {
		log.Panicln("Unable to create cookie jar", err)
	}

	client := &http.Client{
		Jar:           cookieJar,
		CheckRedirect: redirectFunc,
	}

	return client
}

func getRequest(client *http.Client, appArgs AppArgs, url string) (string, *goquery.Document) {
	req, reqErr := http.NewRequest("GET", url, nil)

	req.SetBasicAuth(appArgs.Username, appArgs.Password)
	req.Header.Set("Connection", "keep-alive")

	if reqErr != nil {
		log.Panicln(url, "Unable Create Request", reqErr)
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

	return parseDoc(resp.Body)
}

func postRequest(client *http.Client, appArgs AppArgs, url string, body string) {
	bodyReader := strings.NewReader(body)
	req, reqErr := http.NewRequest("POST", url, bodyReader)

	req.SetBasicAuth(appArgs.Username, appArgs.Password)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if reqErr != nil {
		log.Panicln(url, "Unable Create Request", reqErr)
	}

	var resp *http.Response
	var err error

	resp, err = client.Do(req)

	if err != nil || resp.StatusCode != http.StatusOK {
		log.Panicln(url, "Unable to make request", err, resp)
	}
}

func RouterRequest(appArgs AppArgs) (string, *goquery.Document) {
	client := createClient(appArgs)

	metricsUrl := appArgs.Url + "/RST_stattbl.htm"

	body, doc := getRequest(client, appArgs, metricsUrl)

	// Often the R8000P will ask if to end other sessions that are still active
	// End session on other device and retry.

	if strings.Contains(body, `top.location.href = "MNU_access_multiLogin2.htm";`) {
		fmt.Println("Another Session Is Active, Proceed Anyways")

		multipleLoginsUrl := appArgs.Url + "/MNU_access_multiLogin2.htm"
		_, doc = getRequest(client, appArgs, multipleLoginsUrl)

		action, _ := doc.Find("body > form").Attr("action")

		logoutUrl := appArgs.Url + "/" + action

		postRequest(client, appArgs, logoutUrl, "yes=&act=yes")

		body, doc = getRequest(client, appArgs, metricsUrl)
	}

	return body, doc
}
