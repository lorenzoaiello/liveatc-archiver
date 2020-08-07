package internal

import (
	"io"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func getStations() ([]string, []string, error) {
	resp, err := http.Get("https://www.liveatc.net/archive.php")
	if err != nil {
		return nil, nil, err
	}

	buf := new(strings.Builder)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return nil, nil, err
	}

	s := buf.String()

	var re = regexp.MustCompile(`(?m)\d\d\d\dZ`)
	var stations []string
	var times []string
	d := html.NewTokenizer(strings.NewReader(s))
	for {
		tokenType := d.Next()
		if tokenType == html.ErrorToken {
			break
		}

		token := d.Token()
		switch tokenType {
		case html.StartTagToken:
			if token.Data == "option" {
				for _, a := range token.Attr {
					if a.Key == "value" {
						isTime := false
						for range re.FindAllString(a.Val, -1) {
							isTime = true
						}

						if !isTime {
							stations = append(stations, a.Val)
						} else {
							times = append(times, a.Val)
						}
					}
				}
			}
		}
	}

	return stations, times, nil
}
