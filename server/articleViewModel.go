package server

import "github.com/aureliengasser/planetocd/server/viewModel"

type articleViewModel struct {
	Article          *article
	CurrentPageIndex int
	Pagination       *viewModel.Pagination
	Suggestions      []*article
}
