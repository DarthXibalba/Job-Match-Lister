package main

import (
	"fmt"
	"os"

	"github.com/darthxibalba/webscraper/html"
)

func main() {
	urls := []string{
		"https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/776/",
		"https://leetcode.com/explore/interview/card/top-interview-questions-medium/107/linked-list/785/",
		"https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/648/",
	}

	for _, url := range urls {
		fmt.Printf("url := %s\n", url)
		body, err := html.GetHtmlPage(url)
		if err != nil {
			fmt.Printf("Error occured %s", err)
			os.Exit(1)
		}
		fmt.Printf("body := %s\n\n", body)
	}
}
