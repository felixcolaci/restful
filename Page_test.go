package restful

import "testing"

func TestPage_NewPage(t *testing.T) {

	page := Page{}.NewPage(0, 0)

	if page.Size != 10 {
		t.Errorf("Size was wrong. Got %v expected %v", page.Size, 10)
	}

	if page.Page != 1 {
		t.Errorf("Page was wrong. Got %v expected %v", page.Page, 1)
	}

	page2 := Page{}.NewPage(5, 5)

	if page2.Size != 5 {
		t.Errorf("Size was wrong. Got %v expected %v", page2.Size, 5)
	}

	if page2.Page != 5 {
		t.Errorf("Page was wrong. Got %v expected %v", page2.Page, 5)
	}

}

func TestPage_Skip(t *testing.T) {

	page := Page{}.NewPage(5, 5)

	offset := page.Skip()

	if offset != 20 {
		t.Errorf("Offset was wrong. Got %v expected %v", offset, 20)
	}

}

func TestPage_Next(t *testing.T) {
	page := Page{}.NewPage(5, 5)

	next := page.Next()

	if next.Size != 5 {
		t.Errorf("Size was wrong. Got %v expected %v", next.Size, 5)
	}

	if next.Page != 6 {
		t.Errorf("Page was wrong. Got %v expected %v", next.Page, 6)
	}

}

func TestPage_Validate(t *testing.T) {
	page := Page{}.NewPage(5, 5)

	result := page.Validate()

	if !result {
		t.Error("Validation incorrect.")
	}

}
