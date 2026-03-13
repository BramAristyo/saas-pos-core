package filter

import "math"

type PaginationInput struct {
	PageSize   int `json:"pageSize" form:"pageSize,default=10"`
	PageNumber int `json:"pageNumber" form:"pageNumber,default=1"`
}

type Meta struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"pageSize"`
	TotalRows  int64 `json:"totalRows"`
	TotalPages int   `json:"totalPages"`
	HasNext    bool  `json:"hasNext"`
	HasPrev    bool  `json:"hasPrev"`
}

type PaginationWithInputFilter struct {
	PaginationInput
}

func (p *PaginationWithInputFilter) Offset() int {
	return (p.PageNumber - 1) * p.PageSize
}

func (p *PaginationWithInputFilter) ToMeta(totalRows int64) Meta {
	return Meta{
		Page:       p.PageNumber,
		PageSize:   p.PageSize,
		TotalRows:  totalRows,
		TotalPages: int(math.Ceil(float64(totalRows) / float64(p.PageSize))),
		HasNext:    p.PageNumber < p.PageSize,
		HasPrev:    p.PageNumber > 1,
	}
}
