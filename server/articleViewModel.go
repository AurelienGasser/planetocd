package server

import (
	"github.com/aureliengasser/planetocd/server/cache"
	"github.com/aureliengasser/planetocd/server/viewModel"
)

type articleViewModel struct {
	Article          *cache.Article
	CurrentPageIndex int
	Pagination       *viewModel.Pagination
	Suggestions      []*cache.Article
}
