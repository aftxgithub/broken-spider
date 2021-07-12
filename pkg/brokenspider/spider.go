package brokenspider

import (
	"net/http"
	"net/url"

	"github.com/anaskhan96/soup"
)

type LinkStatus struct {
	URL    string `json:"url"`
	Broken bool   `json:"broken"`
}

type Spider struct{}

func New() *Spider {
	return &Spider{}
}

func (s *Spider) Walk(URL string, out chan LinkStatus) {
	walk(URL, out)
	close(out)
}

func walk(URL string, out chan LinkStatus) {
	resp, _ := soup.Get(URL)
	doc := soup.HTMLParse(resp)

	for _, a := range doc.FindAll("a") {
		href := a.Attrs()["href"]
		if href == "" {
			continue
		}
		// make href absolute
		href = absolutize(URL, href)
		linkStatus := LinkStatus{URL: href, Broken: isBroken(href)}
		out <- linkStatus
		// if link is not broken and is of the same domain, walk it
		if !linkStatus.Broken && isSameDomain(URL, href) {
			walk(href, out)
		}
	}
}

// isBroken reports true if the link is invalid or unreachable
func isBroken(URL string) bool {
	resp, err := http.Get(URL)
	return err != nil || resp.StatusCode != 200
}

// isAbsolute returns true if URL is absolute
func isAbsolute(URL string) bool {
	u, err := url.Parse(URL)
	if err != nil {
		return false
	}
	return u.Scheme != ""
}

func absolutize(baseURL string, URL string) string {
	if isAbsolute(URL) {
		return URL
	}
	u, err := url.Parse(baseURL)
	if err != nil {
		return URL
	}
	ref, err := url.Parse(URL)
	if err != nil {
		return URL
	}
	return u.ResolveReference(ref).String()
}

func isSameDomain(baseURL string, URL string) bool {
	base, err := url.Parse(baseURL)
	if err != nil {
		return false
	}
	ref, err := url.Parse(URL)
	if err != nil {
		return false
	}
	return base.Host == ref.Host
}
