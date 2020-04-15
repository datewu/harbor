package harbor

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func newHTTPClient() *http.Client {
	ts := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{
		Transport: ts,
		Timeout:   5 * time.Minute,
	}
	return c
}

func plainGetJSON(url string, res interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(res)
	return nil
}

func (e Ep) reqHTTPWithCookie(method string, url string, r io.Reader, cHeaders map[string]string) ([]byte, error) {
	client := newHTTPClient()
	req, err := http.NewRequest(method, url, r)
	if err != nil {
		return nil, err
	}
	for k, v := range cHeaders {
		req.Header.Set(k, v)
	}
	c, err := e.cachedCookie()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", c)
	response, err := client.Do(req)
	if nil != err {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode > 400 {
		return nil, errors.New("respond code greate than 400: " + response.Status)
	}
	respBytes, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return nil, err
	}
	return respBytes, nil
}
