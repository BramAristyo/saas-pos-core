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

// Payload Example for PaginationFilterRequest
// {
//   "pageSize": 10,
//   "pageNumber": 1,
// 	 "search": "coffe",
//   "sort": [
//     {
//       "column": "created_at",
//       "orderBy": "desc"
//     },
//     {
//       "column": "price",
//       "orderBy": "asc"
//     }
//   ],
//   "filter": {
//     "product_name": {
//       "type": "contains",
//       "from": "coffee",
//       "to": "",
//       "filterType": "text"
//     },
//     "created_at": {
//       "type": "inRange",
//       "from": "2026-03-01",
//       "to": "2026-03-24",
//       "filterType": "date"
//     }
//   }
// }

type PaginationWithInputFilter struct {
	PaginationInput
	DynamicFilter
}

func (p *PaginationWithInputFilter) Offset() int {
	return (p.PageNumber - 1) * p.PageSize
}

func (p *PaginationWithInputFilter) ToMeta(totalRows int64) Meta {
	totalPages := int(math.Ceil(float64(totalRows) / float64(p.PageSize)))
	return Meta{
		Page:       p.PageNumber,
		PageSize:   p.PageSize,
		TotalRows:  totalRows,
		TotalPages: int(math.Ceil(float64(totalRows) / float64(p.PageSize))),
		HasNext:    p.PageNumber < totalPages,
		HasPrev:    p.PageNumber > 1,
	}
}
