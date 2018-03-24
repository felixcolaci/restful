package restful

import (
	"testing"
	"net/http"
)

func TestExtractPageFromRequest(t *testing.T) {

	{
		request, err := http.NewRequest("GET", "/v1/users/1", nil)
		if err != nil {
			t.Error("Request could not be constructed")
		}

		page := ExtractPageFromRequest(request)

		if page.Size != 10 {
			t.Errorf("Size was wrong. Got %v expected %v", page.Size, 10)
		}

		if page.Page != 1 {
			t.Errorf("Page was wrong. Got %v expected %v", page.Page, 1)
		}
	}

	{
		request2, err := http.NewRequest("GET", "/v1/users?page=5", nil)
		if err != nil {
			t.Error("Request could not be constructed")
		}

		page := ExtractPageFromRequest(request2)

		if page.Size != 10 {
			t.Errorf("Size was wrong. Got %v expected %v", page.Size, 10)
		}

		if page.Page != 5 {
			t.Errorf("Page was wrong. Got %v expected %v", page.Page, 5)
		}
	}

	{
		request, err := http.NewRequest("GET", "/v1/users?size=5", nil)
		if err != nil {
			t.Error("Request could not be constructed")
		}

		page := ExtractPageFromRequest(request)

		if page.Size != 5 {
			t.Errorf("Size was wrong. Got %v expected %v", page.Size, 5)
		}

		if page.Page != 1 {
			t.Errorf("Page was wrong. Got %v expected %v", page.Page, 1)
		}
	}

	{
		request, err := http.NewRequest("GET", "/v1/users?size=5&page=5", nil)
		if err != nil {
			t.Error("Request could not be constructed")
		}

		page := ExtractPageFromRequest(request)

		if page.Size != 5 {
			t.Errorf("Size was wrong. Got %v expected %v", page.Size, 5)
		}

		if page.Page != 5 {
			t.Errorf("Page was wrong. Got %v expected %v", page.Page, 5)
		}
	}

	{
		request, err := http.NewRequest("GET", "/v1/users/1", nil)
		request.URL.Query().Add("size", "abc")
		if err != nil {
			t.Error("Request could not be constructed")
		}

		page := ExtractPageFromRequest(request)

		if page.Size != 10 {
			t.Errorf("Size was wrong. Got %v expected %v", page.Size, 10)
		}

	}

}