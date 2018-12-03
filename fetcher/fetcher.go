package fetcher

import (
	"github.com/ddliu/go-httpclient"
)

// Fetcher :
type Fetcher struct {
	uid      string
	email    string
	password string
	client   *httpclient.HttpClient
}

// NewFetcher :
func NewFetcher(email string, password string, cookies string) *Fetcher {
	client := httpclient.Defaults(httpclient.Map{
		"Accept-Encoding": "gzip, deflate",
		"Accept":          "text/html, application/xhtml+xml, */*",
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36",
		"Cache-Control":   "no-cache",
	})

	fetcher := &Fetcher{
		email:    email,
		password: password,
		client:   client,
	}

	fetcher.checkLogin()

	return fetcher
}

func (f *Fetcher) checkLogin() {

}

func (f *Fetcher) login() {

}

// func LoadCookie() []*http.Cookie {
// 	var cookies = []*http.Cookie{}
// 	return cookies
// }
