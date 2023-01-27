package repositories

import (
	"fmt"
	"math"
	"simple-catalog-v2/models"
)

func GetPaginationLinks(params models.PaginationParams) (models.PaginationLinks, error) {
	var links []models.PageLink

	totalPages := int(math.Ceil(float64(params.TotalRows) / float64(params.PerPage)))

	for i := 1; i <= totalPages; i++ {
		links = append(links, models.PageLink{
			Page:          i,
			Url:           fmt.Sprintf("/%s?page=%s", params.Path, fmt.Sprint(i)),
			IsCurrentPage: i == params.CurrentPage,
		})
	}

	var nextPage, prevPage int

	prevPage = 1
	nextPage = totalPages

	if params.CurrentPage > 2 {
		prevPage = params.CurrentPage - 1
	}

	if params.CurrentPage < totalPages {
		nextPage = params.CurrentPage + 1
	}

	return models.PaginationLinks{
		CurrentPage: fmt.Sprintf("/%s?page=%s", params.Path, fmt.Sprint(params.CurrentPage)),
		NextPage:    fmt.Sprintf("/%s?page=%s", params.Path, fmt.Sprint(nextPage)),
		PrevPage:    fmt.Sprintf("/%s?page=%s", params.Path, fmt.Sprint(prevPage)),
		TotalRows:   params.TotalRows,
		TotalPages:  totalPages,
		Links:       links,
	}, nil
}