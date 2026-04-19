package dto

import (
	"github.com/BramAristyo/saas-pos-core/server/pkg/filter"
)

type ShiftExpenseResponsePagination struct {
	Data []ShiftExpenseResponse `json:"data"`
	Meta filter.Meta            `json:"meta"`
}

func ToShiftExpenseResponsePagination(e []ShiftExpenseResponse, p filter.PaginationWithInputFilter, totalRows int64) ShiftExpenseResponsePagination {
	return ShiftExpenseResponsePagination{
		Data: e,
		Meta: p.ToMeta(totalRows),
	}
}
