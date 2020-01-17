package main

import (
	"net/url"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

type check struct {
	Token       string     `json:"token"`
	URL         string     `json:"url"`
	Alias       string     `json:"alias"`
	LastStatus  int        `json:"last_status"`
	Uptime      float64    `json:"uptime"`
	Down        bool       `json:"down"`
	DownSince   *time.Time `json:"down_since"`
	Error       *time.Time `json:"error"`
	Period      int        `json:"period"`
	ApdexT      float64    `json:"apdex_t"`
	StringMatch string     `json:"string_match"`
	Enabled     bool       `json:"enabled"`
	Published   bool       `json:"published"`
	LastCheckAt time.Time  `json:"last_check_at"`
	NextCheckAt time.Time  `json:"next_check_at"`
	MuteUntil   *time.Time `json:"mute_until"`
	FaviconURL  string     `json:"favicon_url"`
	HTTPVerb    string     `json:"http_verb"`
	HTTPBody    *string    `json:"http_body"`
	TLS         struct {
		TestedAt time.Time `json:"tested_at"`
		Valid    bool      `json:"valid"`
		Error    *string   `json:"error"`
	} `json:"ssl,omitempty"`
}

type status struct {
	Service   string
	LastCheck string
	OK        bool
	Uptime    float64
	Token     string
}

func (c check) status() status {
	service := c.Alias
	if service == "" {
		service = c.URL
	}
	return status{
		Service:   service,
		OK:        !c.Down,
		Uptime:    c.Uptime,
		LastCheck: c.LastCheckAt.UTC().Format("2006-01-02 15:04Z"),
		Token:     c.Token,
	}
}

func (s status) State() string {
	if s.OK {
		return "OK"
	}
	return "Down"
}

func (s status) StateClass() string {
	if s.OK {
		return "up"
	}
	return "down"
}

func getChecks(apiKey string) ([]check, error) {
	var res []check
	client := resty.New()
	client.SetHostURL("https://updown.io/api")
	_, err := client.R().
		SetQueryParam("api-key", apiKey).
		SetResult(&res).
		Get("/checks")
	return res, err
}

func filterPublicDomain(cs []check, domain string) []check {
	var res []check
	for _, c := range cs {
		if !c.Enabled || !c.Published {
			continue
		}
		u, err := url.Parse(c.URL)
		if err != nil {
			continue
		}
		if strings.HasSuffix(u.Hostname(), domain) {
			res = append(res, c)
		}
	}
	return res
}
