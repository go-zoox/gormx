package gormx

import (
	"strings"
	"testing"
)

// TestParamsRangeModesParsingLogic tests the range mode parsing logic
// without requiring zoox.Context setup
func TestParamsRangeModesParsingLogic(t *testing.T) {
	// Test parsing logic for different range formats
	t.Run("Parse default range", func(t *testing.T) {
		// Simulate: age=18,65:range
		where := parseRangeQuery("age", "18,65:range")
		query, _, _ := where.Build()
		expected := "age BETWEEN ? AND ?"
		if query != expected {
			t.Errorf("Expected %v, got %v", expected, query)
		}
	})

	t.Run("Parse open range", func(t *testing.T) {
		// Simulate: age=18,65:range()
		where := parseRangeQuery("age", "18,65:range()")
		query, _, _ := where.Build()
		expected := "(age > ? AND age < ?)"
		if query != expected {
			t.Errorf("Expected %v, got %v", expected, query)
		}
	})

	t.Run("Parse left_closed range", func(t *testing.T) {
		// Simulate: age=18,65:range[)
		where := parseRangeQuery("age", "18,65:range[)")
		query, _, _ := where.Build()
		expected := "(age >= ? AND age < ?)"
		if query != expected {
			t.Errorf("Expected %v, got %v", expected, query)
		}
	})

	t.Run("Parse right_closed range", func(t *testing.T) {
		// Simulate: age=18,65:range(]
		where := parseRangeQuery("age", "18,65:range(]")
		query, _, _ := where.Build()
		expected := "(age > ? AND age <= ?)"
		if query != expected {
			t.Errorf("Expected %v, got %v", expected, query)
		}
	})

	t.Run("Parse explicit closed range", func(t *testing.T) {
		// Simulate: age=18,65:range[]
		where := parseRangeQuery("age", "18,65:range[]")
		query, _, _ := where.Build()
		expected := "age BETWEEN ? AND ?"
		if query != expected {
			t.Errorf("Expected %v, got %v", expected, query)
		}
	})

	t.Run("Parse date range with left_closed", func(t *testing.T) {
		// Simulate: created_at=2023-01-01,2023-12-31:range[)
		where := parseRangeQuery("created_at", "2023-01-01,2023-12-31:range[)")
		query, _, _ := where.Build()
		expected := "(created_at >= ? AND created_at < ?)"
		if query != expected {
			t.Errorf("Expected %v, got %v", expected, query)
		}
	})

	t.Run("All modes with new syntax", func(t *testing.T) {
		modes := []struct {
			urlValue    string
			expectedSQL string
		}{
			{"10,20:range", "field BETWEEN ? AND ?"},
			{"10,20:range[]", "field BETWEEN ? AND ?"},
			{"10,20:range()", "(field > ? AND field < ?)"},
			{"10,20:range[)", "(field >= ? AND field < ?)"},
			{"10,20:range(]", "(field > ? AND field <= ?)"},
		}

		for _, tc := range modes {
			where := parseRangeQuery("field", tc.urlValue)
			query, _, _ := where.Build()
			if query != tc.expectedSQL {
				t.Errorf("Value %s: Expected SQL = %v, got %v", tc.urlValue, tc.expectedSQL, query)
			}
		}
	})
}

// parseRangeQuery simulates the parsing logic from params.go
// This is a helper function for testing
func parseRangeQuery(key, value string) *Where {
	where := NewWhere()

	if strings.Contains(value, ":") {
		parts := strings.Split(value, ":")
		pattern := parts[len(parts)-1]

		if pattern == "range" || pattern == "range[]" || pattern == "range()" || pattern == "range[)" || pattern == "range(]" {
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
		} else {
			where.Set(key, value)
		}
	} else {
		where.Set(key, value)
	}

	return where
}
