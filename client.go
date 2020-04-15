package harbor

import (
	"time"
)

// Ep contains a harbor server credential info
type Ep struct {
	Domain       string
	User         string
	Pwd          string
	cache        time.Duration
	cachedCookie func() (string, error)
}

// NewEndpoint return a new Ep struct
func NewEndpoint(url, u, p string, d time.Duration) Ep {
	e := Ep{
		Domain: url,
		User:   u,
		Pwd:    p,
		cache:  d,
	}
	e.cachedCookie = e.getCookie()
	return e
}
