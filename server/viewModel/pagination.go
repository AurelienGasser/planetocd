package viewModel

import (
	"net/url"
)

type Paginable interface {
	GetPages() []PaginationPage
}

type PaginationPage interface {
	GetURL() *url.URL
	GetPageNumber() int
}

type Pagination struct {
	CurrentPageNumber int
	Pages             map[int]PaginationPage
	PreviousURL       *url.URL
	NextURL           *url.URL
}

func GetPagination(toPaginate Paginable, pageNumber int) *Pagination {
	pages := toPaginate.GetPages()

	res := &Pagination{
		CurrentPageNumber: pageNumber,
		Pages:             make(map[int]PaginationPage, len(pages)),
	}

	for i, p := range pages {
		pn := i + 1
		res.Pages[p.GetPageNumber()] = p
		if pn == pageNumber+1 {
			res.NextURL = p.GetURL()
		}
		if pn == pageNumber-1 {
			res.PreviousURL = p.GetURL()
		}
	}

	return res
}
