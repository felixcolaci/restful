package restful

import (
	"net/http"
	"strconv"
)

func ExtractPageFromRequest(r *http.Request) Page {
	var page = 0
	var size = 0

	//Extract PageNumber from Query
	pageNum := extractIntParamFromQuery("page",  r)
	if pageNum > 0 {
		page = pageNum
	}

	//Extract limit from Query
	sizeNum := extractIntParamFromQuery("size", r)
	if sizeNum > 0 {
		size = sizeNum
	}

	return Page{}.NewPage(page, size)

}

func extractIntParamFromQuery(key string, r *http.Request) (int) {

	paramString := r.URL.Query().Get(key)
	if len(paramString) > 0 {
		return stringToInt(paramString)
	} else {
		return 0
	}

}

func stringToInt(string string) (int) {
	number, err := strconv.ParseInt(string, 10, 0)
	if err != nil {
		return 0
	} else {
		return int(number)
	}
}
