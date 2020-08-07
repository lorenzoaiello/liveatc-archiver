package internal

import (
	"fmt"
	"net/http"

	"github.com/jlaffaye/ftp"
)

func archiveFile(c *ftp.ServerConn, date, station, timestamp, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	remoteFile := fmt.Sprintf("/30-39 Datasets/30 Aviation/30.04 ATC Recordings/%s/%s/%s.mp3", station, date, timestamp)

	_ = c.MakeDir("/30-39 Datasets/30 Aviation/30.04 ATC Recordings/" + station)
	_ = c.MakeDir("/30-39 Datasets/30 Aviation/30.04 ATC Recordings/" + station + "/" + date)

	err = c.Stor(remoteFile, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("successfully archived: %s", remoteFile)

	return nil
}
