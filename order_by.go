package gormx

import (
	"fmt"
	"strings"
)

// OrderByOne is a single order by.
type OrderByOne struct {
	Key    string
	IsDESC bool
}

// OrderBy is a list of order bys.
type OrderBy []OrderByOne

// Set sets a order by.
func (w *OrderBy) Set(key string, IsDESC bool) {
	*w = append(*w, OrderByOne{
		Key:    key,
		IsDESC: IsDESC,
	})
}

// Get gets a order by.
func (w *OrderBy) Get(key string) (bool, bool) {
	for _, v := range *w {
		if v.Key == key {
			return v.IsDESC, true
		}
	}

	return false, false
}

// Del deletes a order by.
func (w *OrderBy) Del(key string) {
	for i, v := range *w {
		if v.Key == key {
			*w = append((*w)[:i], (*w)[i+1:]...)
			break
		}
	}
}

// Debug prints the order bys.
func (w *OrderBy) Debug() {
	for _, v := range *w {
		var desc string
		if v.IsDESC {
			desc = "DESC"
		} else {
			desc = "ASC"
		}

		fmt.Printf("[order_by] %s %s\n", v.Key, desc)
	}
}

// Length returns the length of the order bys.
func (w *OrderBy) Length() int {
	return len(*w)
}

// Build builds the order bys.
func (w *OrderBy) Build() string {
	orders := []string{}
	for _, order := range *w {
		orderMod := "ASC"
		if order.IsDESC {
			orderMod = "DESC"
		}

		orders = append(orders, fmt.Sprintf("%s %s", order.Key, orderMod))
	}

	return strings.Join(orders, ",")
}
