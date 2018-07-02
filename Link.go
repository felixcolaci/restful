package restful

import (
	"net/http"
	"os"
	"strings"
)

type Link struct {
	Href      string `json:"href,omitempty"`
	Rel       string `json:"-"`
	Method    string `json:"method,omitempty"`
	Required  bool   `json:"required,omitempty"`
	SchemaRef *Href   `json:"schemaRef,omitempty"`
}

func (l *Link) MakeInteraction(method string, required bool, schema Href) {
	l.Method = method
	l.Required = required
	l.SchemaRef = &schema
}

func (l Link) From(rel string, ref string, r *http.Request) Link {

	l.Rel = rel
	href := ref

	host := os.Getenv("API_BASE")

	protocol := "http"

	host = r.Host

	if strings.Index(r.Proto, "S") > -1 || strings.Index(r.Proto, "s") > 1 {
		protocol = "https"
	}

	if r.Header.Get("x-forwarded-host") != "" {
		host = r.Header.Get("x-forwarded-host")
	}
	if r.Header.Get("X-Forwarded-Host") != "" {
		host = r.Header.Get("X-Forwarded-Host")
	}

	if r.Header.Get("x-forwarded-proto") != "" {
		protocol = r.Header.Get("x-forwarded-proto")
	}
	if r.Header.Get("X-Forwarded-Proto") != "" {
		protocol = r.Header.Get("X-Forwarded-Proto")
	}

	if strings.Index(ref, "/") == -1 || strings.Index(ref, "/") > 0 {
		href = "/" + ref
	}
	l.Href = protocol + "://" + host + href
	return l
}
