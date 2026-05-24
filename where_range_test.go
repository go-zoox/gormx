package gormx

import (
	"testing"
	"time"
)

func TestWhereRange(t *testing.T) {
	// Test with string dates
	t.Run("String dates", func(t *testing.T) {
		where := NewWhere()
		where.Set("created_at", []string{"2023-01-01", "2023-12-31"}, &SetWhereOptions{IsRange: true})

		query, args, err := where.Build()
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}

		expectedQuery := "created_at BETWEEN ? AND ?"
		if query != expectedQuery {
			t.Errorf("Expected query = %v, got %v", expectedQuery, query)
		}

		if len(args) != 2 {
			t.Errorf("Expected 2 args, got %d", len(args))
		}

		if args[0] != "2023-01-01" || args[1] != "2023-12-31" {
			t.Errorf("Expected args [2023-01-01, 2023-12-31], got %v", args)
		}
	})

	// Test with integers
	t.Run("Integers", func(t *testing.T) {
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

		if args[0] != 18 || args[1] != 65 {
			t.Errorf("Expected args [18, 65], got %v", args)
		}
	})

	// Test with float64
	t.Run("Float64", func(t *testing.T) {
		where := NewWhere()
		where.Set("price", []float64{10.5, 99.9}, &SetWhereOptions{IsRange: true})

		query, args, err := where.Build()
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}

		expectedQuery := "price BETWEEN ? AND ?"
		if query != expectedQuery {
			t.Errorf("Expected query = %v, got %v", expectedQuery, query)
		}

		if len(args) != 2 {
			t.Errorf("Expected 2 args, got %d", len(args))
		}

		if args[0] != 10.5 || args[1] != 99.9 {
			t.Errorf("Expected args [10.5, 99.9], got %v", args)
		}
	})

	// Test with time.Time
	t.Run("Time.Time", func(t *testing.T) {
		where := NewWhere()
		start := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		end := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)
		where.Set("created_at", []interface{}{start, end}, &SetWhereOptions{IsRange: true})

		query, args, err := where.Build()
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}

		expectedQuery := "created_at BETWEEN ? AND ?"
		if query != expectedQuery {
			t.Errorf("Expected query = %v, got %v", expectedQuery, query)
		}

		if len(args) != 2 {
			t.Errorf("Expected 2 args, got %d", len(args))
		}

		if args[0] != start || args[1] != end {
			t.Errorf("Expected args [%v, %v], got %v", start, end, args)
		}
	})

	// Test with multiple conditions
	t.Run("Multiple conditions", func(t *testing.T) {
		where := NewWhere()
		where.Set("age", []int{18, 65}, &SetWhereOptions{IsRange: true})
		where.Set("name", "John", &SetWhereOptions{IsFuzzy: true})
		where.Set("status", "active")

		query, args, err := where.Build()
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}

		expectedQuery := "age BETWEEN ? AND ? AND name ILike ? AND status = ?"
		if query != expectedQuery {
			t.Errorf("Expected query = %v, got %v", expectedQuery, query)
		}

		if len(args) != 4 {
			t.Errorf("Expected 4 args, got %d", len(args))
		}
	})

	// Test with interface{} array
	t.Run("Interface array", func(t *testing.T) {
		where := NewWhere()
		where.Set("score", []interface{}{0, 100}, &SetWhereOptions{IsRange: true})

		query, args, err := where.Build()
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}

		expectedQuery := "score BETWEEN ? AND ?"
		if query != expectedQuery {
			t.Errorf("Expected query = %v, got %v", expectedQuery, query)
		}

		if len(args) != 2 {
			t.Errorf("Expected 2 args, got %d", len(args))
		}

		if args[0] != 0 || args[1] != 100 {
			t.Errorf("Expected args [0, 100], got %v", args)
		}
	})
}
