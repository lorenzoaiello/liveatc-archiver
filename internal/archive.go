package internal

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jlaffaye/ftp"
)

func archiveFile(c *ftp.ServerConn, date, station, timestamp, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	baseDir := os.Getenv("FTP_BASE")

	remoteFile := fmt.Sprintf(baseDir+"/%s/%s/%s.mp3", date, station, timestamp)

	_ = c.MakeDir(baseDir + "/" + date)
	_ = c.MakeDir(baseDir + "/" + date + "/" + station)

	err = c.Stor(remoteFile, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("successfully archived: %s", remoteFile)

	return nil
}
