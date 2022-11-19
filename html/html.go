package html

import (
	"io"
	"net/http"
)

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
