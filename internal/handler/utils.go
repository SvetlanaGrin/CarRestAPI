package handler

import (
	"carRestAPI/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewParams(c *gin.Context) models.Params {
	limit, _ := strconv.Atoi(c.Query("limit"))
	if limit == 0 {
		limit = 10
	}
	offset, _ := strconv.Atoi(c.Query("offset"))
	if offset == 0 {
		offset = 0
	}

	return models.Params{
		Pagination: models.Pagination{
			Limit:  limit,
			Offset: offset,
		},
		Filtration: models.Filtration{
			Mark: c.Query("mark"),
		},
	}
}
