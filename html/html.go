package html

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

/**
 * Get the base url from an http.Response
 * Returns: baseUrl as a string
 */
func GetBaseUrlFromResp(resp *http.Response) string {
	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	return baseUrl.String() // := "https://domain.com"
}

/**
 * Sends a GET request to the given webPageUrl parameter
 * Returns: The body of the response (string)
 */
func GetHtmlPage(webPageUrl string) (string, error) {
	resp, err := http.Get(webPageUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// Read body and return as a string
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/**
 * Reads HTML from a file
 * Returns: The HTML file as a string
 */
func ReadHtmlFromFile(fileName string) (string, error) {
	bs, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
