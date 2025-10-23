package gormx

import (
	"fmt"
	"strings"
)

// WhereCondition is a type constraint for where conditions.
// It can be either map[any]any or *Where.
type WhereCondition interface {
	map[any]any | *Where
}

// ToWhere converts a where condition to *Where.
// If the input is already *Where, return it directly.
// If the input is map[any]any, convert it to *Where.
func ToWhere[T WhereCondition](condition T) *Where {
	var result any = condition

	// If it's already *Where, return directly
	if w, ok := result.(*Where); ok {
		return w
	}

	// If it's map[any]any, convert to *Where
	if m, ok := result.(map[any]any); ok {
		where := NewWhere()
		for k, v := range m {
			if key, ok := k.(string); ok {
				where.Set(key, v)
			}
		}
		return where
	}

	// Should not reach here due to type constraint
	return nil
}

// WhereOne is the where one.
type WhereOne struct {
	Key   string
	Value interface{}

	// IsEqual => =
	IsEqual bool
	// IsNotEqual => !=
	IsNotEqual bool

	// IsFuzzy => ILike
	IsFuzzy bool

	// IsIn => in (?)
	IsIn bool
	// IsNotIn => not in (?)
	IsNotIn bool

	// IsPlain => plain
	IsPlain bool

	// IsFullTextSearch => ILike (field1) OR ILike (field2) OR ...
	IsFullTextSearch     bool
	FullTextSearchFields []string
}

// Where is the where.
type Where struct {
	Items []WhereOne
	//
	FullTextSearchFields []string
}

// SetWhereOptions is the options for SetWhere.
type SetWhereOptions struct {
	IsEqual              bool
	IsNotEqual           bool
	IsFuzzy              bool
	IsIn                 bool
	IsNotIn              bool
	IsPlain              bool
	IsFullTextSearch     bool
	FullTextSearchFields []string
}

// NewWhere returns a new where.
func NewWhere() *Where {
	return &Where{}
}

// Set sets a where, if exists, update.
func (w *Where) Set(key string, value interface{}, opts ...*SetWhereOptions) {
	// @TODO cannot real update

	_, ok := w.Get(key)
	if ok {
		w.Del(key)
	}

	w.Add(key, value, opts...)
}

// Add adds a where, if exists, append.
func (w *Where) Add(key string, value interface{}, opts ...*SetWhereOptions) {
	item := WhereOne{
		Key:   key,
		Value: value,
	}

	for _, opt := range opts {
		if opt == nil {
			continue
		}

		item.IsFuzzy = opt.IsFuzzy
		item.IsEqual = opt.IsEqual
		item.IsNotEqual = opt.IsNotEqual
		item.IsIn = opt.IsIn
		item.IsNotIn = opt.IsNotIn
		item.IsPlain = opt.IsPlain
		item.IsFullTextSearch = opt.IsFullTextSearch
		item.FullTextSearchFields = opt.FullTextSearchFields
	}

	w.Items = append(w.Items, item)
}

// Get gets a where.
func (w *Where) Get(key string) (interface{}, bool) {
	for _, v := range w.Items {
		if v.Key == key {
			return v.Value, true
		}
	}

	return "", false
}

// Del deletes a where.
func (w *Where) Del(key string) {
	for i, v := range w.Items {
		if v.Key == key {
			w.Items = append(w.Items[:i], w.Items[i+1:]...)
			return
		}
	}
}

// Length returns the length of the wheres.
func (w *Where) Length() int {
	return len(w.Items)
}

// Build builds the wheres.
func (w *Where) Build() (query string, args []interface{}, err error) {
	whereClauses := []string{}
	whereValues := []interface{}{}
	for _, item := range w.Items {
		// @TODO built-in keywords
		if item.Key == "q" {
			item.IsFullTextSearch = true
			item.FullTextSearchFields = w.FullTextSearchFields
		}

		// @TODO full text search search keyword
		if item.IsFullTextSearch {
			// ignore if no fields
			if len(item.FullTextSearchFields) == 0 {
				// return "", nil, fmt.Errorf("FullTextSearchFields is required when IsFullTextSearch is true (key: %s)", item.Key)
				// continue
				if len(w.FullTextSearchFields) == 0 {
					// return "", nil, fmt.Errorf("FullTextSearchFields is required when IsFullTextSearch is true (key: %s)", item.Key)
					continue
				}

				item.FullTextSearchFields = w.FullTextSearchFields
			}

			keyword, v := item.Value.(string)
			if !v {
				return "", nil, fmt.Errorf("value must be string when IsFullTextSearch is true (key: %s)", item.Key)
			}

			// @TODO
			keywordExtract := strings.Replace(keyword, ":*", "", 1)

			//
			keywordFuzzy := fmt.Sprintf("%%%s%%", keywordExtract)
			qs := []string{}
			args := []interface{}{}

			fields := item.FullTextSearchFields
			for _, field := range fields {
				qs = append(qs, fmt.Sprintf("%s ILike ?", field))
				args = append(args, keywordFuzzy)
			}
			query := strings.Join(qs, " OR ")

			whereClauses = append(whereClauses, fmt.Sprintf("(%s)", query))
			whereValues = append(whereValues, args...)
		} else {
			if item.IsFuzzy {
				whereClauses = append(whereClauses, fmt.Sprintf("%s ILike ?", item.Key))
				whereValues = append(whereValues, fmt.Sprintf("%%%s%%", item.Value))
			} else if item.IsEqual {
				whereClauses = append(whereClauses, fmt.Sprintf("%s = ?", item.Key))
				whereValues = append(whereValues, item.Value)
			} else if item.IsNotEqual {
				whereClauses = append(whereClauses, fmt.Sprintf("%s != ?", item.Key))
				whereValues = append(whereValues, item.Value)
			} else if item.IsIn {
				whereClauses = append(whereClauses, fmt.Sprintf("%s in (?)", item.Key))
				whereValues = append(whereValues, item.Value)
			} else if item.IsNotIn {
				whereClauses = append(whereClauses, fmt.Sprintf("%s not in (?)", item.Key))
				whereValues = append(whereValues, item.Value)
			} else if item.IsPlain {
				whereClauses = append(whereClauses, fmt.Sprintf("(%s)", item.Key))
				if v, ok := item.Value.([]any); ok {
					whereValues = append(whereValues, v...)
				} else {
					whereValues = append(whereValues, item.Value)
				}
			} else {
				whereClauses = append(whereClauses, fmt.Sprintf("%s = ?", item.Key))
				whereValues = append(whereValues, item.Value)
			}
		}
	}
	whereClause := strings.Join(whereClauses, " AND ")

	return whereClause, whereValues, nil
}

// Debug prints the wheres.
func (w *Where) Debug() {
	for _, item := range w.Items {
		var fuzzy string
		if item.IsFuzzy {
			fuzzy = "Fuzzy"
		} else {
			fuzzy = "Extract"
		}

		fmt.Printf("[where] %s %s %s\n", item.Key, item.Value, fuzzy)
	}
}

// Reset resets the wheres.
func (w *Where) Reset() {
	w.Items = []WhereOne{}
}
