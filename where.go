package gormx

import (
	"fmt"
	"strings"
)

// WhereOne is the where one.
type WhereOne struct {
	Key     string
	Value   interface{}
	IsFuzzy bool
	isNot   bool
	isIn    bool
	//
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
	IsFuzzy              bool
	IsNot                bool
	IsIn                 bool
	IsFullTextSearch     bool
	FullTextSearchFields []string
}

// Set sets a where.
func (w *Where) Set(key string, value interface{}, opts ...*SetWhereOptions) {
	var isFuzzy bool
	var isNot bool
	var isIn bool
	var isFullTextSearch bool
	var fullTextSearchFields []string
	if len(opts) > 0 && opts[0] != nil {
		isFuzzy = opts[0].IsFuzzy
		isNot = opts[0].IsNot
		isIn = opts[0].IsIn
		isFullTextSearch = opts[0].IsFullTextSearch
		fullTextSearchFields = opts[0].FullTextSearchFields
	}

	w.Items = append(w.Items, WhereOne{
		Key:     key,
		Value:   value,
		IsFuzzy: isFuzzy,
		isNot:   isNot,
		isIn:    isIn,
		//
		IsFullTextSearch:     isFullTextSearch,
		FullTextSearchFields: fullTextSearchFields,
	})
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
				continue
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
			} else if item.isNot {
				whereClauses = append(whereClauses, fmt.Sprintf("%s != ?", item.Key))
				whereValues = append(whereValues, item.Value)
			} else if item.isIn {
				whereClauses = append(whereClauses, fmt.Sprintf("%s in (?)", item.Key))
				whereValues = append(whereValues, item.Value)
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
