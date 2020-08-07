package internal

import (
	"fmt"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

func Run() {
	today := time.Now().Add(-24 * time.Hour)
	formattedDate := today.Format("20060102")

	stations, times, err := getStations()
	if err != nil {
		panic(err)
	}

	c, err := ftp.Dial(os.Getenv("FTP_HOST"), ftp.DialWithTimeout(10*time.Second))
	if err != nil {
		panic(err)
	}

	err = c.Login(os.Getenv("FTP_USER"), os.Getenv("FTP_PASS"))
	if err != nil {
		panic(err)
	}

	for _, station := range stations {
		for _, selectedTime := range times {
			url, err := getDownloadUrl(formattedDate, station, selectedTime)
			if err != nil {
				fmt.Printf("error retrieving download url for %s, %s, %s: %s", formattedDate, station, selectedTime, err.Error())
				continue
			}

			err = archiveFile(c, formattedDate, station, selectedTime, url)
			if err != nil {
				fmt.Printf("error archiving file %s: %s", url, err.Error())
				continue
			}
		}
	}

	if err := c.Quit(); err != nil {
		panic(err)
	}
}
