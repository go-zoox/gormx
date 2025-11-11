package gormx

import (
	"testing"
)

func TestWhereRangeModes(t *testing.T) {
	// Test Closed Range [A, B]
	t.Run("Closed Range [A, B]", func(t *testing.T) {
		where := NewWhere()
		where.Set("age", []int{18, 65}, &SetWhereOptions{
			IsRange:   true,
			RangeMode: RangeModeClosed,
		})

		query, args, err := where.Build()
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}

		expectedQuery := "age BETWEEN ? AND ?"
		if query != expectedQuery {
			t.Errorf("Expected query = %v, got %v", expectedQuery, query)
		}

		if len(args) != 2 {
			t.Errorf("Expected 2 args, got %d", len(args))
		}

		if args[0] != 18 || args[1] != 65 {
			t.Errorf("Expected args [18, 65], got %v", args)
		}
	})

	// Test Default (should be Closed)
	t.Run("Default Range (should be Closed)", func(t *testing.T) {
		where := NewWhere()
		where.Set("age", []int{18, 65}, &SetWhereOptions{IsRange: true})

		query, args, err := where.Build()
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}

		expectedQuery := "age BETWEEN ? AND ?"
		if query != expectedQuery {
			t.Errorf("Expected query = %v, got %v", expectedQuery, query)
		}

		if len(args) != 2 {
			t.Errorf("Expected 2 args, got %d", len(args))
		}
	})

	// Test Open Range (A, B)
	t.Run("Open Range (A, B)", func(t *testing.T) {
		where := NewWhere()
		where.Set("age", []int{18, 65}, &SetWhereOptions{
			IsRange:   true,
			RangeMode: RangeModeOpen,
		})

		query, args, err := where.Build()
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}

		expectedQuery := "(age > ? AND age < ?)"
		if query != expectedQuery {
			t.Errorf("Expected query = %v, got %v", expectedQuery, query)
		}

		if len(args) != 2 {
			t.Errorf("Expected 2 args, got %d", len(args))
		}

		if args[0] != 18 || args[1] != 65 {
			t.Errorf("Expected args [18, 65], got %v", args)
		}
	})

	// Test Left Closed Range [A, B)
	t.Run("Left Closed Range [A, B)", func(t *testing.T) {
		where := NewWhere()
		where.Set("age", []int{18, 65}, &SetWhereOptions{
			IsRange:   true,
			RangeMode: RangeModeLeftClosed,
		})

		query, args, err := where.Build()
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}

		expectedQuery := "(age >= ? AND age < ?)"
		if query != expectedQuery {
			t.Errorf("Expected query = %v, got %v", expectedQuery, query)
		}

		if len(args) != 2 {
			t.Errorf("Expected 2 args, got %d", len(args))
		}

		if args[0] != 18 || args[1] != 65 {
			t.Errorf("Expected args [18, 65], got %v", args)
		}
	})

	// Test Right Closed Range (A, B]
	t.Run("Right Closed Range (A, B]", func(t *testing.T) {
		where := NewWhere()
		where.Set("age", []int{18, 65}, &SetWhereOptions{
			IsRange:   true,
			RangeMode: RangeModeRightClosed,
		})

		query, args, err := where.Build()
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}

		expectedQuery := "(age > ? AND age <= ?)"
		if query != expectedQuery {
			t.Errorf("Expected query = %v, got %v", expectedQuery, query)
		}

		if len(args) != 2 {
			t.Errorf("Expected 2 args, got %d", len(args))
		}

		if args[0] != 18 || args[1] != 65 {
			t.Errorf("Expected args [18, 65], got %v", args)
		}
	})

	// Test with float values
	t.Run("Open Range with floats", func(t *testing.T) {
		where := NewWhere()
		where.Set("price", []float64{10.5, 99.9}, &SetWhereOptions{
			IsRange:   true,
			RangeMode: RangeModeOpen,
		})

		query, args, err := where.Build()
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}

		expectedQuery := "(price > ? AND price < ?)"
		if query != expectedQuery {
			t.Errorf("Expected query = %v, got %v", expectedQuery, query)
		}

		if args[0] != 10.5 || args[1] != 99.9 {
			t.Errorf("Expected args [10.5, 99.9], got %v", args)
		}
	})

	// Test with date strings
	t.Run("Left Closed Range with dates", func(t *testing.T) {
		where := NewWhere()
		where.Set("created_at", []string{"2023-01-01", "2023-12-31"}, &SetWhereOptions{
			IsRange:   true,
			RangeMode: RangeModeLeftClosed,
		})

		query, args, err := where.Build()
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}

		expectedQuery := "(created_at >= ? AND created_at < ?)"
		if query != expectedQuery {
			t.Errorf("Expected query = %v, got %v", expectedQuery, query)
		}

		if args[0] != "2023-01-01" || args[1] != "2023-12-31" {
			t.Errorf("Expected args [2023-01-01, 2023-12-31], got %v", args)
		}
	})

	// Test multiple conditions with different range modes
	t.Run("Multiple conditions with different range modes", func(t *testing.T) {
		where := NewWhere()
		where.Add("age", []int{18, 30}, &SetWhereOptions{
			IsRange:   true,
			RangeMode: RangeModeLeftClosed,
		})
		where.Add("score", []int{60, 100}, &SetWhereOptions{
			IsRange:   true,
			RangeMode: RangeModeClosed,
		})
		where.Set("status", "active")

		query, args, err := where.Build()
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}

		expectedQuery := "(age >= ? AND age < ?) AND score BETWEEN ? AND ? AND status = ?"
		if query != expectedQuery {
			t.Errorf("Expected query = %v, got %v", expectedQuery, query)
		}

		if len(args) != 5 {
			t.Errorf("Expected 5 args, got %d", len(args))
		}
	})

	// Test all four modes in one test
	t.Run("All four modes comparison", func(t *testing.T) {
		modes := []struct {
			mode     RangeMode
			expected string
		}{
			{RangeModeClosed, "age BETWEEN ? AND ?"},
			{RangeModeOpen, "(age > ? AND age < ?)"},
			{RangeModeLeftClosed, "(age >= ? AND age < ?)"},
			{RangeModeRightClosed, "(age > ? AND age <= ?)"},
		}

		for _, tc := range modes {
			where := NewWhere()
			where.Set("age", []int{18, 65}, &SetWhereOptions{
				IsRange:   true,
				RangeMode: tc.mode,
			})

			query, _, err := where.Build()
			if err != nil {
				t.Fatalf("Build() error = %v for mode %s", err, tc.mode)
			}

			if query != tc.expected {
				t.Errorf("Mode %s: Expected query = %v, got %v", tc.mode, tc.expected, query)
			}
		}
	})
}
