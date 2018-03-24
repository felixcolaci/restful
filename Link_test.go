package restful

import (
	"net/http"
	"testing"
)

func TestLink_From_WithoutHeaders(t *testing.T) {

	request, err := http.NewRequest("GET", "/v1/users/1", nil)
	request.Host = "localhost"
	if err != nil {
		t.Error("Request could not be constructed")
	}

	link := Link{}.From("self", "/v1/users/1", request)

	expectedRel := "self"
	if link.Rel != expectedRel {
		t.Errorf("Rel is not matching. Expected %v got %v", expectedRel, link.Rel)
	}

	expectedHref := "http://localhost/v1/users/1"

	if link.Href != expectedHref {
		t.Errorf("Href is not matching. Expected %v got %v", expectedHref, link.Href)
	}

}
func TestLink_From_WithoutLeadingSlash(t *testing.T) {

	request, err := http.NewRequest("GET", "/v1/users/1", nil)
	request.Host = "localhost"
	if err != nil {
		t.Error("Request could not be constructed")
	}

	link := Link{}.From("self", "v1/users/1", request)

	expectedRel := "self"
	if link.Rel != expectedRel {
		t.Errorf("Rel is not matching. Expected %v got %v", expectedRel, link.Rel)
	}

	expectedHref := "http://localhost/v1/users/1"

	if link.Href != expectedHref {
		t.Errorf("Href is not matching. Expected %v got %v", expectedHref, link.Href)
	}

}

func TestLink_From_WithXForwardedHost(t *testing.T) {

	request, err := http.NewRequest("GET", "/v1/users/1", nil)
	request.Host = "localhost"
	request.Header.Set("x-forwarded-host", "example.com")

	if err != nil {
		t.Error("Request could not be constructed")
	}

	link := Link{}.From("self", "/v1/users/1", request)

	expectedRel := "self"
	if link.Rel != expectedRel {
		t.Errorf("Rel is not matching. Expected %v got %v", expectedRel, link.Rel)
	}

	expectedHref := "http://example.com/v1/users/1"

	if link.Href != expectedHref {
		t.Errorf("Href is not matching. Expected %v got %v", expectedHref, link.Href)
	}
}

func TestLink_From_WithXForwardedHost2(t *testing.T) {

	request, err := http.NewRequest("GET", "/v1/users/1", nil)
	request.Host = "localhost"
	request.Header.Set("X-Forwarded-Host", "example.com")

	if err != nil {
		t.Error("Request could not be constructed")
	}

	link := Link{}.From("self", "/v1/users/1", request)

	expectedRel := "self"
	if link.Rel != expectedRel {
		t.Errorf("Rel is not matching. Expected %v got %v", expectedRel, link.Rel)
	}

	expectedHref := "http://example.com/v1/users/1"

	if link.Href != expectedHref {
		t.Errorf("Href is not matching. Expected %v got %v", expectedHref, link.Href)
	}
}

func TestLink_From_XForwardedProto(t *testing.T) {

	request, err := http.NewRequest("GET", "/v1/users/1", nil)
	request.Host = "localhost"
	request.Header.Set("x-forwarded-proto", "https")

	if err != nil {
		t.Error("Request could not be constructed")
	}

	link := Link{}.From("self", "/v1/users/1", request)

	expectedRel := "self"
	if link.Rel != expectedRel {
		t.Errorf("Rel is not matching. Expected %v got %v", expectedRel, link.Rel)
	}

	expectedHref := "https://localhost/v1/users/1"

	if link.Href != expectedHref {
		t.Errorf("Href is not matching. Expected %v got %v", expectedHref, link.Href)
	}

}

func TestLink_From_XForwardedProto2(t *testing.T) {

	request, err := http.NewRequest("GET", "/v1/users/1", nil)
	request.Host = "localhost"
	request.Header.Set("X-Forwarded-Proto", "https")

	if err != nil {
		t.Error("Request could not be constructed")
	}

	link := Link{}.From("self", "/v1/users/1", request)

	expectedRel := "self"
	if link.Rel != expectedRel {
		t.Errorf("Rel is not matching. Expected %v got %v", expectedRel, link.Rel)
	}

	expectedHref := "https://localhost/v1/users/1"

	if link.Href != expectedHref {
		t.Errorf("Href is not matching. Expected %v got %v", expectedHref, link.Href)
	}

}

func TestLink_From_Proto(t *testing.T) {

	request, err := http.NewRequest("GET", "/v1/users/1", nil)
	request.Host = "localhost"
	request.Proto = "https"

	if err != nil {
		t.Error("Request could not be constructed")
	}

	link := Link{}.From("self", "/v1/users/1", request)

	expectedRel := "self"
	if link.Rel != expectedRel {
		t.Errorf("Rel is not matching. Expected %v got %v", expectedRel, link.Rel)
	}

	expectedHref := "https://localhost/v1/users/1"

	if link.Href != expectedHref {
		t.Errorf("Href is not matching. Expected %v got %v", expectedHref, link.Href)
	}

}

func TestLink_From_Proto2(t *testing.T) {

	request, err := http.NewRequest("GET", "/v1/users/1", nil)
	request.Host = "localhost"
	request.Proto = "HTTPS"

	if err != nil {
		t.Error("Request could not be constructed")
	}

	link := Link{}.From("self", "/v1/users/1", request)

	expectedRel := "self"
	if link.Rel != expectedRel {
		t.Errorf("Rel is not matching. Expected %v got %v", expectedRel, link.Rel)
	}

	expectedHref := "https://localhost/v1/users/1"

	if link.Href != expectedHref {
		t.Errorf("Href is not matching. Expected %v got %v", expectedHref, link.Href)
	}

}
