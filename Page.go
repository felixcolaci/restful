package restful

type Page struct {
	Page  int `json:"page,omitempty"`
	Size  int `json:"size,omitempty"`
	Total int `json:"total,omitempty"`
}

func (p Page) NewPage(page int, size int) Page {
	p.Page = page
	p.Size = size

	if page == 0 {
		p.Page = 1
	}
	if size == 0 {
		p.Size = 10
	}
	return p
}

func (p Page) Validate() bool {
	if p.Page < 1 || p.Size < 1 {
		return false
	} else {
		return true
	}
}

func (p Page) Skip() int {
	factor := p.Page - 1
	return int(factor * p.Size)
}

func (p Page) Next() Page {

	page := p

	page.Page++

	return page
}
