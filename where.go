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
type Where []WhereOne

// SetWhereOptions is the options for SetWhere.
type SetWhereOptions struct {
	IsFuzzy              bool
	IsNot                bool
	IsIn                 bool
	IsFullTextSearch     bool
	FullTextSearchFields []string
}

// Set sets a where.
func (w *Where) Set(key string, value interface{}, opts ...*SetWhereOptions) error {
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

	if isFullTextSearch {
		if len(fullTextSearchFields) == 0 {
			return fmt.Errorf("FullTextSearchFields is required when IsFullTextSearch is true (key: %s)", key)
		}
	}

	*w = append(*w, WhereOne{
		Key:     key,
		Value:   value,
		IsFuzzy: isFuzzy,
		isNot:   isNot,
		isIn:    isIn,
		//
		IsFullTextSearch:     isFullTextSearch,
		FullTextSearchFields: fullTextSearchFields,
	})

	return nil
}

// Get gets a where.
func (w *Where) Get(key string) (interface{}, bool) {
	for _, v := range *w {
		if v.Key == key {
			return v.Value, true
		}
	}

	return "", false
}

// Del deletes a where.
func (w *Where) Del(key string) {
	for i, v := range *w {
		if v.Key == key {
			*w = append((*w)[:i], (*w)[i+1:]...)
			return
		}
	}
}

// Length returns the length of the wheres.
func (w *Where) Length() int {
	return len(*w)
}

// Build builds the wheres.
func (w *Where) Build() (string, []interface{}) {
	whereClauses := []string{}
	whereValues := []interface{}{}
	for _, w := range *w {
		// @TODO full text search search keyword
		if w.IsFullTextSearch {
			keywordExtract := strings.Replace(w.Key, ":*", "", 1)
			keywordFuzzy := fmt.Sprintf("%%%s%%", keywordExtract)
			qs := []string{}
			args := []interface{}{}

			fields := w.FullTextSearchFields
			for _, field := range fields {
				qs = append(qs, fmt.Sprintf("%s ILike ?", field))
				args = append(args, keywordFuzzy)
			}
			query := strings.Join(qs, " OR ")

			whereClauses = append(whereClauses, fmt.Sprintf("(%s)", query))
			whereValues = append(whereValues, args...)
		} else {
			if w.IsFuzzy {
				whereClauses = append(whereClauses, fmt.Sprintf("%s ILike ?", w.Key))
				whereValues = append(whereValues, fmt.Sprintf("%%%s%%", w.Value))
			} else if w.isNot {
				whereClauses = append(whereClauses, fmt.Sprintf("%s != ?", w.Key))
				whereValues = append(whereValues, w.Value)
			} else if w.isIn {
				whereClauses = append(whereClauses, fmt.Sprintf("%s in (?)", w.Key))
				whereValues = append(whereValues, w.Value)
			} else {
				whereClauses = append(whereClauses, fmt.Sprintf("%s = ?", w.Key))
				whereValues = append(whereValues, w.Value)
			}
		}
	}
	whereClause := strings.Join(whereClauses, " AND ")

	return whereClause, whereValues
}

// Debug prints the wheres.
func (w *Where) Debug() {
	for _, where := range *w {
		var fuzzy string
		if where.IsFuzzy {
			fuzzy = "Fuzzy"
		} else {
			fuzzy = "Extract"
		}

		fmt.Printf("[where] %s %s %s\n", where.Key, where.Value, fuzzy)
	}
}
