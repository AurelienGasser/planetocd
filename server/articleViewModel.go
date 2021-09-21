package server

type articleViewModel struct {
	Article          *article
	CurrentPageIndex int
	Pagination       *pagination
}
