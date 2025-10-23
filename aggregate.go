package gormx

import (
	"fmt"
	"reflect"
	"strings"
)

// AggregateResult represents the result of an aggregate query
type AggregateResult struct {
	Value interface{}
	Error error
}

// Sum calculates the sum of a numeric field
func Sum[T any](field string, where *Where) (float64, error) {
	var result struct {
		Sum float64 `gorm:"column:sum"`
	}

	query := GetDB().Model(new(T))

	if where != nil {
		whereClause, whereValues, err := where.Build()
		if err != nil {
			return 0, err
		}
		if whereClause != "" {
			query = query.Where(whereClause, whereValues...)
		}
	}

	err := query.Select(fmt.Sprintf("SUM(%s) as sum", field)).Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result.Sum, nil
}

// Avg calculates the average of a numeric field
func Avg[T any](field string, where *Where) (float64, error) {
	var result struct {
		Avg float64 `gorm:"column:avg"`
	}

	query := GetDB().Model(new(T))

	if where != nil {
		whereClause, whereValues, err := where.Build()
		if err != nil {
			return 0, err
		}
		if whereClause != "" {
			query = query.Where(whereClause, whereValues...)
		}
	}

	err := query.Select(fmt.Sprintf("AVG(%s) as avg", field)).Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result.Avg, nil
}

// Min finds the minimum value of a field
func Min[T any](field string, where *Where) (interface{}, error) {
	var result map[string]interface{}

	query := GetDB().Model(new(T))

	if where != nil {
		whereClause, whereValues, err := where.Build()
		if err != nil {
			return nil, err
		}
		if whereClause != "" {
			query = query.Where(whereClause, whereValues...)
		}
	}

	err := query.Select(fmt.Sprintf("MIN(%s) as min", field)).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result["min"], nil
}

// Max finds the maximum value of a field
func Max[T any](field string, where *Where) (interface{}, error) {
	var result map[string]interface{}

	query := GetDB().Model(new(T))

	if where != nil {
		whereClause, whereValues, err := where.Build()
		if err != nil {
			return nil, err
		}
		if whereClause != "" {
			query = query.Where(whereClause, whereValues...)
		}
	}

	err := query.Select(fmt.Sprintf("MAX(%s) as max", field)).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result["max"], nil
}

// CountDistinct counts distinct values of a field
func CountDistinct[T any](field string, where *Where) (int64, error) {
	var result struct {
		Count int64 `gorm:"column:count"`
	}

	query := GetDB().Model(new(T))

	if where != nil {
		whereClause, whereValues, err := where.Build()
		if err != nil {
			return 0, err
		}
		if whereClause != "" {
			query = query.Where(whereClause, whereValues...)
		}
	}

	err := query.Select(fmt.Sprintf("COUNT(DISTINCT %s) as count", field)).Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result.Count, nil
}

// GroupByResult represents the result of a group by query
type GroupByResult struct {
	Group map[string]interface{}
	Count int64
	Sum   float64
	Avg   float64
	Min   interface{}
	Max   interface{}
}

// GroupBy performs group by operations with optional aggregations
func GroupBy[T any](fields []string, where *Where, aggregates []string) ([]GroupByResult, error) {
	if len(fields) == 0 {
		return nil, fmt.Errorf("group by fields cannot be empty")
	}

	// Build the SELECT clause
	selectClause := strings.Join(fields, ", ")

	// Add aggregate functions if specified
	if len(aggregates) > 0 {
		selectClause += ", " + strings.Join(aggregates, ", ")
	}

	query := GetDB().Model(new(T))

	if where != nil {
		whereClause, whereValues, err := where.Build()
		if err != nil {
			return nil, err
		}
		if whereClause != "" {
			query = query.Where(whereClause, whereValues...)
		}
	}

	// Execute the query
	rows, err := query.Select(selectClause).Group(strings.Join(fields, ", ")).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []GroupByResult

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		// Create a slice to hold the values
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		// Scan the row
		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		// Create the result
		result := GroupByResult{
			Group: make(map[string]interface{}),
		}

		// Map the values to the result
		for i, column := range columns {
			if i < len(fields) {
				// This is a group by field
				result.Group[column] = values[i]
			} else {
				// This is an aggregate value
				val := values[i]
				if val != nil {
					switch column {
					case "count":
						if count, ok := val.(int64); ok {
							result.Count = count
						}
					case "sum":
						if sum, ok := val.(float64); ok {
							result.Sum = sum
						}
					case "avg":
						if avg, ok := val.(float64); ok {
							result.Avg = avg
						}
					case "min":
						result.Min = val
					case "max":
						result.Max = val
					}
				}
			}
		}

		results = append(results, result)
	}

	return results, nil
}

// Aggregate performs multiple aggregate operations in a single query
func Aggregate[T any](field string, where *Where, operations []string) (map[string]interface{}, error) {
	if len(operations) == 0 {
		return nil, fmt.Errorf("aggregate operations cannot be empty")
	}

	// Build the SELECT clause with multiple aggregates
	var selectParts []string
	for _, op := range operations {
		switch strings.ToLower(op) {
		case "sum":
			selectParts = append(selectParts, fmt.Sprintf("SUM(%s) as sum", field))
		case "avg":
			selectParts = append(selectParts, fmt.Sprintf("AVG(%s) as avg", field))
		case "min":
			selectParts = append(selectParts, fmt.Sprintf("MIN(%s) as min", field))
		case "max":
			selectParts = append(selectParts, fmt.Sprintf("MAX(%s) as max", field))
		case "count":
			selectParts = append(selectParts, fmt.Sprintf("COUNT(%s) as count", field))
		case "count_distinct":
			selectParts = append(selectParts, fmt.Sprintf("COUNT(DISTINCT %s) as count_distinct", field))
		default:
			return nil, fmt.Errorf("unsupported aggregate operation: %s", op)
		}
	}

	query := GetDB().Model(new(T))

	if where != nil {
		whereClause, whereValues, err := where.Build()
		if err != nil {
			return nil, err
		}
		if whereClause != "" {
			query = query.Where(whereClause, whereValues...)
		}
	}

	var result map[string]interface{}
	err := query.Select(strings.Join(selectParts, ", ")).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetFieldType returns the type of a field in a struct
func GetFieldType[T any](field string) (reflect.Type, error) {
	var model T
	t := reflect.TypeOf(model)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Name == field {
			return f.Type, nil
		}
	}

	return nil, fmt.Errorf("field %s not found in struct", field)
}
