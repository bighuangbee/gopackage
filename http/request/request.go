package request

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type RequestParams map[string]interface{}
type RequestHeaders map[string]string

func Request(method , addr, url string, params RequestParams, headers RequestHeaders) ([]byte, error) {
	paramsBytes, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest(method, addr + url, bytes.NewReader(paramsBytes))

	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	for key, val := range headers{
		req.Header.Add(key, val)
	}

	response, err := (&http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}).Do(req)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK{
		return nil, errors.New("StatusCode " + fmt.Sprintf("%d", response.StatusCode))
	}

	var body []byte
	switch response.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(response.Body)
			for {
				buf := make([]byte, 1024)
				n, err := reader.Read(buf)

				if err != nil && err != io.EOF {
					return nil, err
				}

				if n == 0 {
					break
				}
				body = append(body, buf...);
			}
		default:
			body, _ = ioutil.ReadAll(response.Body)

	}
	return body, nil
}