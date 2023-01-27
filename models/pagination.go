package models

type PageLink struct {
	Page          int
	Url           string
	IsCurrentPage bool
}

type PaginationLinks struct {
	CurrentPage string
	NextPage    string
	PrevPage    string
	TotalRows   int
	TotalPages  int
	Links       []PageLink
}

type PaginationParams struct {
	Path        string
	TotalRows   int
	PerPage     int
	CurrentPage int
}