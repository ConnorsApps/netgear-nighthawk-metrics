package utils

import (
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
		log.Fatalln(err)
	}

	bodyString := string(bodyBytes)
	if bodyString == "" {
		log.Fatalln("No body found")
	}
	parsedHtmlReader := strings.NewReader(bodyString)

	doc, err := goquery.NewDocumentFromReader(parsedHtmlReader)
	if err != nil {
		log.Fatalln(err)
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
		log.Fatalln("Unable to create cookie jar", err)
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

	return parseDoc(resp.Body)
}

func postRequest(client *http.Client, appArgs AppArgs, url string) (string, *goquery.Document) {
	req, reqErr := http.NewRequest("POST", url, nil)

	req.SetBasicAuth(appArgs.Username, appArgs.Password)
	req.Header.Set("Connection", "keep-alive")

	if reqErr != nil {
		log.Fatalln(url, "Unable Create Request", reqErr)
	}

	var resp *http.Response
	var err error

	resp, err = client.Do(req)

	if err != nil || resp.StatusCode != http.StatusOK {
		log.Panicln(url, "Unable to make request", err, resp)
	}

	return parseDoc(resp.Body)
}

func RouterRequest(appArgs AppArgs) (string, *goquery.Document) {
	client := createClient(appArgs)

	metricsUrl := appArgs.Url + "/RST_stattbl.htm"

	body, doc := getRequest(client, appArgs, metricsUrl)

	// Often the R8000P will ask if to end other sessions that are still active
	// End session on other device and retry.

	if strings.Contains(body, `top.location.href = "MNU_access_multiLogin2.htm";`) {
		multipleLoginsUrl := appArgs.Url + "/MNU_access_multiLogin2.htm"
		body, doc = getRequest(client, appArgs, multipleLoginsUrl)

		action, exists := doc.Find("body > form").Attr("action")
		log.Println("action", action, "exists", exists)
		log.Println("body", body)
		// document.forms[0].act.value="yes";
		// document.forms[0].submit();
		// use proxy

		logoutUrl := appArgs.Url + "/" + action
		body, _ = postRequest(client, appArgs, logoutUrl)

		log.Println("body", body)

		body, doc = getRequest(client, appArgs, metricsUrl)
	}

	return body, doc
}
