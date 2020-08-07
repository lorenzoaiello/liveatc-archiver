package internal

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func getDownloadUrl(date, facility, time string) (string, error) {
	resp, err := http.PostForm("https://www.liveatc.net/listen.php", url.Values{
		"date":     []string{date},
		"facility": []string{facility},
		"time":     []string{time},
	})
	if err != nil {
		return "", err
	}

	buf := new(strings.Builder)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return "", err
	}

	s := buf.String()

	var re = regexp.MustCompile(`(?m)<source src="http:\/\/archive-server.liveatc.net\/(.*)" `)
	for _, match := range re.FindAllString(s, -1) {
		cleanMatch := strings.Replace(strings.Replace(match, "<source src=\"", "", -1), "\" ", "", -1)
		return cleanMatch, nil
	}

	return "", fmt.Errorf("no match found")
}
