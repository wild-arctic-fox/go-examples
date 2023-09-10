package coinapp

import (
	"fmt"
	"net/http"
	"time"
)

type LoggerRoundTripper struct {
	next http.RoundTripper
}

func (l LoggerRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	println(fmt.Sprintf("[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL))
	r.URL.RawQuery = "limit=1"
	println(fmt.Sprintf("[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL))
	return l.next.RoundTrip(r)
}
