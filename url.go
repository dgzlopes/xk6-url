package url

import (
	"net"
	"net/url"

	"github.com/goware/urlx"
	"github.com/loadimpact/k6/js/modules"
)

func init() {
	modules.Register("k6/x/url", new(URL))
}

// URL is the k6 url extension.
type URL struct{}

// SPLIT is used by SplitHostPort
type SPLIT struct {
	Host string
	Port string
}

// Parse parses raw URL string into the net/url URL struct.
func (*URL) Parse(url string) *url.URL {
	parsed, err := urlx.Parse(url)
	if err != nil {
		ReportError(err, "Failed to parse the url")
	}
	return parsed
}

// Normalize returns normalized URL.
func (*URL) Normalize(url string) string {
	parsed, err := urlx.Parse(url)
	if err != nil {
		ReportError(err, "Failed to parse the url")
	}
	normalized, err := urlx.Normalize(parsed)
	if err != nil {
		ReportError(err, "Failed to normalize the url")
	}
	return normalized
}

// SplitHostPort splits network address of the form "host:port" into host and port.
func (*URL) SplitHostPort(url string) *SPLIT {
	parsed, err := urlx.Parse(url)
	if err != nil {
		ReportError(err, "Failed to parse the url")
	}
	host, port, err := urlx.SplitHostPort(parsed)
	if err != nil {
		ReportError(err, "Failed to split the url")
	}
	return &SPLIT{Host: host, Port: port}
}

// Resolve resolves the URL host to its IP address.
func (*URL) Resolve(url string) *net.IPAddr {
	parsed, err := urlx.Parse(url)
	ip, err := urlx.Resolve(parsed)
	if err != nil {
		ReportError(err, "Failed to resolve the url")
	}
	return ip
}
