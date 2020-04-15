package client

import (
	"net/http"
	"strings"
	"time"
)

func (e Ep) buildCredential() string {
	crd := "principal=" + e.User + "&password=" + e.Pwd
	return crd

}

func (e Ep) currentUser() error {
	const path = "/api/users/current"
	_, err := e.reqHTTPWithCookie("GET", e.Domain+path, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (e Ep) getCookie() func() (string, error) {
	const (
		key  = "Set-Cookie"
		path = "/api/systeminfo"
	)
	t := time.Now()
	c := ""
	req := func() (string, error) {
		if time.Since(t) > e.cache {
			t = time.Now()
			c = ""
		}
		if c != "" {
			return c, nil
		}
		resp, err := http.Get(e.Domain + path)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		setCookie := resp.Header.Get(key)
		cks := strings.Split(setCookie, "; ")
		c = cks[0]
		return c, nil
	}
	return req

}
