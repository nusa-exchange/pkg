package filters

import (
	"gorm.io/gorm"

	"github.com/nusa-exchange/pkg/gpa"
)

func WithLimit(limit int) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Limit(limit)
	}
}

func WithOrder(order string) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Order(order)
	}
}

func WithOffset(offset int) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		return query.Offset(offset)
	}
}

func WithPageable(page, size int, order string) gpa.Filter {
	return func(query *gorm.DB) *gorm.DB {
		query = query.Limit(size)
		if page > 0 {
			offset := (page - 1) * size
			query = query.Offset(offset)
		}
		if order != "" {
			query = query.Order(order)
		}
		return query
	}
}
