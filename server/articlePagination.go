package server

import (
	"net/url"
)

type pagination struct {
	CurrentPageNumber int
	Pages             map[int]articlePage
	PreviousURL       *url.URL
	NextURL           *url.URL
}

func getPagination(article *article, pageNumber int) *pagination {
	res := &pagination{
		CurrentPageNumber: pageNumber,
		Pages:             make(map[int]articlePage, len(article.Pages)),
	}

	for i, p := range article.Pages {
		pn := i + 1
		res.Pages[p.PageNumber] = p
		if pn == pageNumber+1 {
			res.NextURL = article.Pages[i].URL
		}
		if pn == pageNumber-1 {
			res.PreviousURL = article.Pages[i].URL
		}
	}

	return res
}
