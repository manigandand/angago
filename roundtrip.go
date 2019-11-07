package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type customTransport struct {
	http.RoundTripper
}

// RoundTrip custom RoundTrip method which implements the RoundTrip interface
// add custom headers
func (t *customTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	resp, err = t.RoundTripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	resp.Header.Set("Host", req.Host)
	if strings.TrimSpace(resp.Header.Get("Content-Type")) == "" {
		resp.Header.Set("Content-Type", "application/json")
	}
	return resp, nil
}

func modifyProxyResponse(resp *http.Response) (err error) {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = resp.Body.Close()
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	defer gz.Close()
	_, err = gz.Write(b)
	if err != nil {
		return err
	}
	if err = gz.Flush(); err != nil {
		return err
	}
	compressedData := buf.Bytes()
	resp.Body = ioutil.NopCloser(bytes.NewReader(compressedData))
	resp.ContentLength = int64(len(b))
	resp.Header.Set("Content-Length", strconv.Itoa(len(b)))
	if strings.TrimSpace(resp.Header.Get("Content-Type")) == "" {
		resp.Header.Set("Content-Type", "application/json")
	}
	resp.Header.Set("Content-Encoding", "gzip")

	return nil
}
