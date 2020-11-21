package common

import (
	"io"
	"net/http"
)

// DownloadFileUrl downloads an arbitrary file from the provided url and returns a reader
func DownloadFileUrl(url string) (r io.Reader, err error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	r = res.Body

	return res.Body, nil
}
