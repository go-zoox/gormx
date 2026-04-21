package gormx

import (
	"fmt"
	"strings"
)

// ParamsValue defines the minimum value abilities used by Params.
type ParamsValue interface {
	String() string
	Int64() int64
}

// ParamsGetter defines a key-value getter for query/param objects.
type ParamsGetter interface {
	Get(key string) ParamsValue
}

// ParamsQueries defines query collection operations used by Params.
type ParamsQueries interface {
	Del(key string)
	Iterator() map[string]any
}

// ParamsContext defines the minimum context abilities required by Params.
type ParamsContext interface {
	BindQuery(obj interface{}) error
	Param() ParamsGetter
	Query() ParamsGetter
	Queries() ParamsQueries
}

// Params is the interface that wraps the basic methods.
type Params struct {
	ctx ParamsContext
	//
	page *Page
}

// NewParams returns the params.
func NewParams(ctx ParamsContext) *Params {
	return &Params{
		ctx: ctx,
	}
}

func (c *Params) parsePage() error {
	if c.page != nil {
		return nil
	}

	c.page = &Page{}
	if err := c.ctx.BindQuery(c.page); err != nil {
		return err
	}

	if c.page.PageSize == 0 {
		c.page.PageSize = 10
	}

	if c.page.Page == 0 {
		c.page.Page = 1
	}

	if c.page.Page > 1000 {
		c.page.Page = 1000
	}

	if c.page.PageSize > 100 {
		c.page.PageSize = 100
	}

	return nil
}

// Page is the struct that wraps the basic fields.
func (c *Params) Page() (uint, error) {
	if err := c.parsePage(); err != nil {
		return 0, err
	}

	return uint(c.page.Page), nil
}

// PageSize is the struct that wraps the basic fields.
func (c *Params) PageSize() (uint, error) {
	if err := c.parsePage(); err != nil {
		return 0, err
	}

	return uint(c.page.PageSize), nil
}

// ID is the struct that wraps the basic fields.
func (c *Params) ID() (uint, error) {
	id := c.ctx.Param().Get("id").Int64()

	if id == 0 {
		return 0, fmt.Errorf("invalid id: %s", c.ctx.Param().Get("id").String())
	}

	return uint(id), nil
}

// Where is the struct that wraps the basic fields.
func (c *Params) Where() *Where {
	var where Where

	whereObject := c.ctx.Queries()

	whereObject.Del("page")
	whereObject.Del("pageSize")
	whereObject.Del("orderBy")

	whereObject.Del("page-size")
	whereObject.Del("order-by")

	for key, value := range whereObject.Iterator() {
		if vs, ok := value.(string); ok {
			if strings.Contains(vs, ":") {
				parts := strings.Split(vs, ":")
				pattern := parts[len(parts)-1]

				if pattern == "*" {
					value := strings.Join(parts[0:len(parts)-1], ":")
					where.Set(key, value, &SetWhereOptions{IsFuzzy: true})
				} else if pattern == "!" {
					value := strings.Join(parts[0:len(parts)-1], ":")
					where.Set(key, value, &SetWhereOptions{IsNotEqual: true})
				} else if pattern == "in" {
					value := strings.Join(parts[0:len(parts)-1], ":")
					where.Set(key, strings.Split(value, ","), &SetWhereOptions{IsIn: true})
				} else if pattern == "!in" {
					value := strings.Join(parts[0:len(parts)-1], ":")
					where.Set(key, strings.Split(value, ","), &SetWhereOptions{IsNotIn: true})
				} else if pattern == "range" || pattern == "range[]" || pattern == "range()" || pattern == "range[)" || pattern == "range(]" {
					// Format: startValue,endValue:range_mode
					// Examples:
					//   18,65:range       -> [A, B] (default, closed)
					//   18,65:range[]     -> [A, B] (closed)
					//   18,65:range()     -> (A, B) (open)
					//   18,65:range[)     -> [A, B) (left closed)
					//   18,65:range(]     -> (A, B] (right closed)

					var rangeMode RangeMode = RangeModeClosed // default

					// Determine range mode based on pattern
					switch pattern {
					case "range", "range[]":
						rangeMode = RangeModeClosed
					case "range()":
						rangeMode = RangeModeOpen
					case "range[)":
						rangeMode = RangeModeLeftClosed
					case "range(]":
						rangeMode = RangeModeRightClosed
					}

					valueStr := strings.Join(parts[0:len(parts)-1], ":")
					rangeValues := strings.Split(valueStr, ",")
					if len(rangeValues) >= 2 {
						where.Set(key, rangeValues, &SetWhereOptions{
							IsRange:   true,
							RangeMode: rangeMode,
						})
					}
				} else if pattern == ">" {
					value := strings.Join(parts[0:len(parts)-1], ":")
					where.Set(key, value, &SetWhereOptions{IsGreaterThan: true})
				} else if pattern == "<" {
					value := strings.Join(parts[0:len(parts)-1], ":")
					where.Set(key, value, &SetWhereOptions{IsLessThan: true})
				} else if pattern == ">=" {
					value := strings.Join(parts[0:len(parts)-1], ":")
					where.Set(key, value, &SetWhereOptions{IsGreaterOrEqualThan: true})
				} else if pattern == "<=" {
					value := strings.Join(parts[0:len(parts)-1], ":")
					where.Set(key, value, &SetWhereOptions{IsLessOrEqualThan: true})
				} else {
					where.Set(key, vs)
				}
			} else {
				where.Set(key, vs)
			}
		} else {
			where.Set(key, value)
		}
	}

	return &where
}

// OrderBy is the struct that wraps the basic fields.
func (c *Params) OrderBy() *OrderBy {
	var orderBy OrderBy

	orderByRaw := c.ctx.Query().Get("orderBy").String()
	if orderByRaw == "" {
		orderByRaw = c.ctx.Query().Get("order-by").String()
	}

	if orderByRaw != "" {
		orderByRaws := strings.Split(orderByRaw, ",")
		for _, one := range orderByRaws {
			one = strings.TrimSpace(one)
			if one == "" {
				continue
			}

			orderByRaws := strings.Split(one, ":")
			if len(orderByRaws) != 2 {
				continue
			}

			key := orderByRaws[0]
			order := strings.ToLower(orderByRaws[1])
			isDESC := false
			if order == "desc" {
				isDESC = true
			} else if order == "DESC" {
				isDESC = true
			}

			orderBy.Set(key, isDESC)
		}
	}

	return &orderBy
}

// ListParams is the struct that wraps the basic fields.
type ListParams struct {
	Page     uint
	PageSize uint
	Where    *Where
	OrderBy  *OrderBy
}

// ListParamsDefault is the struct that wraps the basic fields.
type ListParamsDefault struct {
	Page     uint
	PageSize uint
}

// GetList is the struct that wraps the basic fields.
func (c *Params) GetList(defaults ...*ListParamsDefault) (*ListParams, error) {
	var defaultsX *ListParamsDefault
	if len(defaults) > 0 && defaults[0] != nil {
		defaultsX = defaults[0]
	}

	var listParams ListParams
	var err error

	listParams.Page, err = c.Page()
	if err != nil {
		return nil, fmt.Errorf("parse page error: %s", err)
	} else if defaultsX != nil && listParams.Page == 0 && defaultsX.Page != 0 {
		listParams.Page = defaultsX.Page
	}

	listParams.PageSize, err = c.PageSize()
	if err != nil {
		return nil, fmt.Errorf("parse page size error: %s", err)
	} else if defaultsX != nil && listParams.PageSize == 0 && defaultsX.PageSize != 0 {
		listParams.PageSize = defaultsX.PageSize
	}

	listParams.Where = c.Where()
	listParams.OrderBy = c.OrderBy()

	return &listParams, nil
}
