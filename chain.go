package gormx

import (
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

// QueryBuilder provides a fluent interface for building database queries
type QueryBuilder[T any] struct {
	db       *gorm.DB
	model    *T
	where    *Where
	selects  []string
	orders   *OrderBy
	joins    []JoinClause
	preloads []string
	limit    *int
	offset   *int
	group    []string
	having   *Where
	distinct bool
}

// JoinClause represents a join clause
type JoinClause struct {
	Type      string // "INNER", "LEFT", "RIGHT", "FULL"
	Table     string
	Condition string
	Args      []interface{}
}

// NewQuery creates a new query builder for the given model type
func NewQuery[T any]() *QueryBuilder[T] {
	return &QueryBuilder[T]{
		db:       GetDB(),
		model:    new(T),
		where:    NewWhere(),
		orders:   &OrderBy{},
		joins:    make([]JoinClause, 0),
		preloads: make([]string, 0),
		group:    make([]string, 0),
		having:   NewWhere(),
	}
}

// Where adds a WHERE condition to the query
func (q *QueryBuilder[T]) Where(field string, value interface{}, opts ...*SetWhereOptions) *QueryBuilder[T] {
	q.where.Set(field, value, opts...)
	return q
}

// WhereEqual adds an equality WHERE condition
func (q *QueryBuilder[T]) WhereEqual(field string, value interface{}) *QueryBuilder[T] {
	return q.Where(field, value, &SetWhereOptions{IsEqual: true})
}

// WhereNotEqual adds a not equal WHERE condition
func (q *QueryBuilder[T]) WhereNotEqual(field string, value interface{}) *QueryBuilder[T] {
	return q.Where(field, value, &SetWhereOptions{IsNotEqual: true})
}

// WhereIn adds an IN WHERE condition
func (q *QueryBuilder[T]) WhereIn(field string, values interface{}) *QueryBuilder[T] {
	return q.Where(field, values, &SetWhereOptions{IsIn: true})
}

// WhereNotIn adds a NOT IN WHERE condition
func (q *QueryBuilder[T]) WhereNotIn(field string, values interface{}) *QueryBuilder[T] {
	return q.Where(field, values, &SetWhereOptions{IsNotIn: true})
}

// WhereLike adds a LIKE WHERE condition
func (q *QueryBuilder[T]) WhereLike(field string, value string) *QueryBuilder[T] {
	return q.Where(field, value, &SetWhereOptions{IsFuzzy: true})
}

// WhereBetween adds a BETWEEN WHERE condition
func (q *QueryBuilder[T]) WhereBetween(field string, start, end interface{}) *QueryBuilder[T] {
	// This would need to be implemented in the Where.Build() method
	q.where.Add(field, []interface{}{start, end}, &SetWhereOptions{IsPlain: true})
	return q
}

// WhereRaw adds a raw WHERE condition
func (q *QueryBuilder[T]) WhereRaw(sql string, args ...interface{}) *QueryBuilder[T] {
	q.where.Add("", nil, &SetWhereOptions{IsPlain: true})
	// Store raw SQL separately - this would need special handling in Build()
	return q
}

// Select specifies the columns to be selected
func (q *QueryBuilder[T]) Select(columns ...string) *QueryBuilder[T] {
	q.selects = append(q.selects, columns...)
	return q
}

// OrderBy adds an ORDER BY clause
func (q *QueryBuilder[T]) OrderBy(field string, desc ...bool) *QueryBuilder[T] {
	isDesc := false
	if len(desc) > 0 {
		isDesc = desc[0]
	}
	q.orders.Set(field, isDesc)
	return q
}

// OrderByAsc adds an ascending ORDER BY clause
func (q *QueryBuilder[T]) OrderByAsc(field string) *QueryBuilder[T] {
	return q.OrderBy(field, false)
}

// OrderByDesc adds a descending ORDER BY clause
func (q *QueryBuilder[T]) OrderByDesc(field string) *QueryBuilder[T] {
	return q.OrderBy(field, true)
}

// Limit sets the LIMIT clause
func (q *QueryBuilder[T]) Limit(limit int) *QueryBuilder[T] {
	q.limit = &limit
	return q
}

// Offset sets the OFFSET clause
func (q *QueryBuilder[T]) Offset(offset int) *QueryBuilder[T] {
	q.offset = &offset
	return q
}

// Page sets both LIMIT and OFFSET for pagination
func (q *QueryBuilder[T]) Page(page, pageSize int) *QueryBuilder[T] {
	if pageSize > 0 {
		q.limit = &pageSize
	}
	if page > 0 && pageSize > 0 {
		offset := (page - 1) * pageSize
		q.offset = &offset
	}
	return q
}

// Join adds a JOIN clause
func (q *QueryBuilder[T]) Join(table, condition string, args ...interface{}) *QueryBuilder[T] {
	q.joins = append(q.joins, JoinClause{
		Type:      "INNER",
		Table:     table,
		Condition: condition,
		Args:      args,
	})
	return q
}

// LeftJoin adds a LEFT JOIN clause
func (q *QueryBuilder[T]) LeftJoin(table, condition string, args ...interface{}) *QueryBuilder[T] {
	q.joins = append(q.joins, JoinClause{
		Type:      "LEFT",
		Table:     table,
		Condition: condition,
		Args:      args,
	})
	return q
}

// RightJoin adds a RIGHT JOIN clause
func (q *QueryBuilder[T]) RightJoin(table, condition string, args ...interface{}) *QueryBuilder[T] {
	q.joins = append(q.joins, JoinClause{
		Type:      "RIGHT",
		Table:     table,
		Condition: condition,
		Args:      args,
	})
	return q
}

// Preload adds a PRELOAD clause for eager loading associations
func (q *QueryBuilder[T]) Preload(associations ...string) *QueryBuilder[T] {
	q.preloads = append(q.preloads, associations...)
	return q
}

// GroupBy adds a GROUP BY clause
func (q *QueryBuilder[T]) GroupBy(fields ...string) *QueryBuilder[T] {
	q.group = append(q.group, fields...)
	return q
}

// Having adds a HAVING clause
func (q *QueryBuilder[T]) Having(field string, value interface{}, opts ...*SetWhereOptions) *QueryBuilder[T] {
	q.having.Set(field, value, opts...)
	return q
}

// Distinct adds a DISTINCT clause
func (q *QueryBuilder[T]) Distinct() *QueryBuilder[T] {
	q.distinct = true
	return q
}

// buildQuery builds the final GORM query
func (q *QueryBuilder[T]) buildQuery() *gorm.DB {
	query := q.db.Model(q.model)

	// Apply WHERE conditions
	if q.where != nil && len(q.where.Items) > 0 {
		whereClause, whereValues, err := q.where.Build()
		if err == nil && whereClause != "" {
			query = query.Where(whereClause, whereValues...)
		}
	}

	// Apply SELECT
	if len(q.selects) > 0 {
		query = query.Select(strings.Join(q.selects, ", "))
	}

	// Apply DISTINCT
	if q.distinct {
		query = query.Distinct()
	}

	// Apply JOINs
	for _, join := range q.joins {
		query = query.Joins(fmt.Sprintf("%s JOIN %s ON %s", join.Type, join.Table, join.Condition), join.Args...)
	}

	// Apply ORDER BY
	if q.orders != nil && len(*q.orders) > 0 {
		for _, order := range *q.orders {
			orderMod := "ASC"
			if order.IsDESC {
				orderMod = "DESC"
			}
			query = query.Order(fmt.Sprintf("%s %s", order.Key, orderMod))
		}
	}

	// Apply GROUP BY
	if len(q.group) > 0 {
		query = query.Group(strings.Join(q.group, ", "))
	}

	// Apply HAVING
	if q.having != nil && len(q.having.Items) > 0 {
		havingClause, havingValues, err := q.having.Build()
		if err == nil && havingClause != "" {
			query = query.Having(havingClause, havingValues...)
		}
	}

	// Apply LIMIT and OFFSET
	if q.limit != nil {
		query = query.Limit(*q.limit)
	}
	if q.offset != nil {
		query = query.Offset(*q.offset)
	}

	// Apply PRELOADs
	for _, preload := range q.preloads {
		query = query.Preload(preload)
	}

	return query
}

// Find executes the query and returns all matching records
func (q *QueryBuilder[T]) Find() ([]*T, error) {
	var results []*T
	err := q.buildQuery().Find(&results).Error
	return results, err
}

// First executes the query and returns the first matching record
func (q *QueryBuilder[T]) First() (*T, error) {
	var result T
	err := q.buildQuery().First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Last executes the query and returns the last matching record
func (q *QueryBuilder[T]) Last() (*T, error) {
	var result T
	err := q.buildQuery().Last(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Count executes the query and returns the count of matching records
func (q *QueryBuilder[T]) Count() (int64, error) {
	var count int64
	err := q.buildQuery().Count(&count).Error
	return count, err
}

// Exists checks if any records match the query conditions
func (q *QueryBuilder[T]) Exists() (bool, error) {
	count, err := q.Count()
	return count > 0, err
}

// Delete executes the query and deletes all matching records
func (q *QueryBuilder[T]) Delete() error {
	return q.buildQuery().Delete(q.model).Error
}

// Update executes the query and updates all matching records
func (q *QueryBuilder[T]) Update(updates map[string]interface{}) error {
	return q.buildQuery().Updates(updates).Error
}

// UpdateColumn executes the query and updates specific columns
func (q *QueryBuilder[T]) UpdateColumn(updates map[string]interface{}) error {
	return q.buildQuery().UpdateColumns(updates).Error
}

// Scan executes the query and scans the result into the provided destination
func (q *QueryBuilder[T]) Scan(dest interface{}) error {
	return q.buildQuery().Scan(dest).Error
}

// Raw executes a raw SQL query and returns the result
func (q *QueryBuilder[T]) Raw(sql string, values ...interface{}) *QueryBuilder[T] {
	q.db = q.db.Raw(sql, values...)
	return q
}

// Clone creates a copy of the query builder
func (q *QueryBuilder[T]) Clone() *QueryBuilder[T] {
	clone := &QueryBuilder[T]{
		db:       q.db,
		model:    q.model,
		where:    q.where,
		selects:  make([]string, len(q.selects)),
		orders:   q.orders,
		joins:    make([]JoinClause, len(q.joins)),
		preloads: make([]string, len(q.preloads)),
		limit:    q.limit,
		offset:   q.offset,
		group:    make([]string, len(q.group)),
		having:   q.having,
		distinct: q.distinct,
	}

	copy(clone.selects, q.selects)
	copy(clone.joins, q.joins)
	copy(clone.preloads, q.preloads)
	copy(clone.group, q.group)

	return clone
}

// GetTableName returns the table name for the model
func (q *QueryBuilder[T]) GetTableName() string {
	var model T
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Try to get table name from model if it implements TableName() method
	if tableNamer, ok := interface{}(model).(interface{ TableName() string }); ok {
		return tableNamer.TableName()
	}

	// Fallback to GORM's naming convention
	return strings.ToLower(t.Name())
}

// ToSQL returns the generated SQL query as a string (for debugging)
func (q *QueryBuilder[T]) ToSQL() (string, error) {
	query := q.buildQuery()
	stmt := query.Statement
	return stmt.SQL.String(), nil
}

// Aggregate Methods

// Sum calculates the sum of a field
func (q *QueryBuilder[T]) Sum(field string) (float64, error) {
	var result struct {
		Sum float64 `gorm:"column:sum"`
	}

	query := q.buildQuery().Select(fmt.Sprintf("SUM(%s) as sum", field))
	err := query.Scan(&result).Error
	return result.Sum, err
}

// Avg calculates the average of a field
func (q *QueryBuilder[T]) Avg(field string) (float64, error) {
	var result struct {
		Avg float64 `gorm:"column:avg"`
	}

	query := q.buildQuery().Select(fmt.Sprintf("AVG(%s) as avg", field))
	err := query.Scan(&result).Error
	return result.Avg, err
}

// Min finds the minimum value of a field
func (q *QueryBuilder[T]) Min(field string) (interface{}, error) {
	var result map[string]interface{}

	query := q.buildQuery().Select(fmt.Sprintf("MIN(%s) as min", field))
	err := query.Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result["min"], nil
}

// Max finds the maximum value of a field
func (q *QueryBuilder[T]) Max(field string) (interface{}, error) {
	var result map[string]interface{}

	query := q.buildQuery().Select(fmt.Sprintf("MAX(%s) as max", field))
	err := query.Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result["max"], nil
}

// Pluck retrieves a single column as a slice
func (q *QueryBuilder[T]) Pluck(column string, dest interface{}) error {
	return q.buildQuery().Pluck(column, dest).Error
}

// Transaction Methods

// Transaction executes a function within a database transaction
func (q *QueryBuilder[T]) Transaction(fn func(tx *QueryBuilder[T]) error) error {
	return q.db.Transaction(func(tx *gorm.DB) error {
		txQuery := q.Clone()
		txQuery.db = tx
		return fn(txQuery)
	})
}

// Create inserts a new record
func (q *QueryBuilder[T]) Create(value *T) error {
	return q.db.Create(value).Error
}

// Save saves the record (insert if not exists, update if exists)
func (q *QueryBuilder[T]) Save(value *T) error {
	return q.db.Save(value).Error
}

// CreateInBatches inserts records in batches
func (q *QueryBuilder[T]) CreateInBatches(values []*T, batchSize int) error {
	return q.db.CreateInBatches(values, batchSize).Error
}

// FindInBatches processes records in batches
func (q *QueryBuilder[T]) FindInBatches(batchSize int, fn func(tx *gorm.DB, batch int) error) error {
	var results []*T
	return q.buildQuery().FindInBatches(&results, batchSize, fn).Error
}

// Paginate returns paginated results with total count
func (q *QueryBuilder[T]) Paginate(page, pageSize int) ([]*T, int64, error) {
	// Get total count
	var total int64
	countQuery := q.Clone().buildQuery()
	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	q.Page(page, pageSize)
	results, err := q.Find()
	return results, total, err
}

// Chunk processes records in chunks
func (q *QueryBuilder[T]) Chunk(chunkSize int, callback func([]*T) error) error {
	offset := 0
	for {
		query := q.Clone()
		query.Limit(chunkSize).Offset(offset)

		results, err := query.Find()
		if err != nil {
			return err
		}

		if len(results) == 0 {
			break
		}

		if err := callback(results); err != nil {
			return err
		}

		if len(results) < chunkSize {
			break
		}

		offset += chunkSize
	}
	return nil
}

// GetDB returns the underlying GORM DB instance
func (q *QueryBuilder[T]) GetDB() *gorm.DB {
	return q.buildQuery()
}
