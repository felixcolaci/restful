package restful

import (
	"errors"
	"math"
	"net/http"
	"strconv"
)

type CollectionLinks struct {
	firstRel string
	nextRel  string
	prevRel  string
	lastRel  string
	selfRel  string
	ref      string
	Links    map[string]Href
}

func NewCollectionLinks(collectionRef string) *CollectionLinks {
	return &CollectionLinks{"first", "next", "previous", "last", "self", collectionRef, nil}
}

func (c CollectionLinks) GetLinks(page Page, r *http.Request) (map[string]Href, error) {

	if page.Page < 1 || page.Size < 0 || page.Total < 0 {
		return nil, errors.New("Page invalid")
	}

	c.Links = c.add(Link{}.From(c.selfRel, c.ref+"?page="+strconv.Itoa(page.Page)+"&size="+strconv.Itoa(page.Size), r))

	//next link & last link
	if page.Page*page.Size < page.Total {
		//next
		c.Links = c.add(Link{}.From(c.nextRel, c.ref+"?page="+strconv.Itoa(page.Page+1)+"&size="+strconv.Itoa(page.Size), r))
		//last
		c.Links = c.add(Link{}.From(c.lastRel, c.ref+"?page="+strconv.Itoa(int(math.Ceil(float64(page.Total)/float64(page.Size))))+"&size="+strconv.Itoa(page.Size), r))
	}

	//prev link & first link
	if page.Page > 1 {
		//prev
		c.Links = c.add(Link{}.From(c.prevRel, c.ref+"?page="+strconv.Itoa(page.Page-1)+"&size="+strconv.Itoa(page.Size), r))
		//first
		c.Links = c.add(Link{}.From(c.firstRel, c.ref+"?page=1&size="+strconv.Itoa(page.Size), r))

	}
	return c.Links, nil
}

func (c CollectionLinks) add(link Link) map[string]Href {
	if len(c.Links) < 1 || c.Links == nil {
		c.Links = map[string]Href{link.Rel: Href{link.Href}}
	} else {
		c.Links[link.Rel] = Href{link.Href}
	}
	return c.Links
}
