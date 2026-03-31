package database

import (
	"strings"
	"time"

	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"gorm.io/gorm"
)

type PreloadEntity struct {
	Entity string
}

func BuildQuery(
	db *gorm.DB,
	req filter.DynamicFilter,
	searchable []string,
	allowedFields map[string]string,
) *gorm.DB {
	db = applySearch(db, req.Search, searchable, allowedFields)
	db = applySearch(db, req.Search, searchable, allowedFields)
	return db
}

func applySearch(
	db *gorm.DB,
	search string,
	searchable []string,
	allowedFields map[string]string,
) *gorm.DB {
	if search == "" {
		return db
	}

	conditions := make([]string, len(searchable))
	args := make([]any, len(searchable))

	for _, field := range searchable {
		if col, ok := allowedFields[field]; ok {
			conditions = append(conditions, col+" ILIKE ?")
			args = append(args, "%"+search+"%")
		}
	}

	if len(conditions) > 0 {
		db = db.Where("("+strings.Join(conditions, " OR ")+")", args...)
	}

	return db
}

func applySort(
	db *gorm.DB,
	sorts []filter.Sort,
	allowedFields map[string]string,
) *gorm.DB {
	for _, s := range sorts {
		col, ok := allowedFields[s.Column]
		if !ok {
			continue
		}

		order := strings.ToLower(s.Order)
		if order != "asc" && order != "desc" {
			continue
		}

		db = db.Order(col + " " + order)
	}

	return db
}

func applyFilter(
	db *gorm.DB,
	filters map[string]filter.Filter,
	allowedFields map[string]string,
) *gorm.DB {

	// TODO: Handle Al Greater etc...

	for field, f := range filters {
		col, ok := allowedFields[field]
		if !ok {
			continue
		}

		if f.FilterType == filter.DataTypeDate && f.Type == filter.OpInRange {

			if f.From == "" || f.To == "" {
				continue
			}

			fromDate, err := time.Parse("2006-01-02", f.From)
			if err != nil {
				continue
			}

			toDate, err := time.Parse("2006-01-02", f.To)
			if err != nil {
				continue
			}

			toDate = toDate.Add(24*time.Hour - time.Nanosecond)

			db = db.Where(col+" BETWEEN ? AND ?", fromDate, toDate)
		}
	}

	return db
}
