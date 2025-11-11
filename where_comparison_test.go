package gormx

import (
	"testing"
	"time"
)

func TestWhereComparison(t *testing.T) {
	// Test IsGreaterThan
	t.Run("IsGreaterThan", func(t *testing.T) {
		where := NewWhere()
		where.Set("age", 18, &SetWhereOptions{IsGreaterThan: true})

		query, args, err := where.Build()
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}

		expectedQuery := "age > ?"
		if query != expectedQuery {
			t.Errorf("Expected query = %v, got %v", expectedQuery, query)
		}

		if len(args) != 1 {
			t.Errorf("Expected 1 arg, got %d", len(args))
		}

		if args[0] != 18 {
			t.Errorf("Expected arg 18, got %v", args[0])
		}
	})

	// Test IsLessThan
	t.Run("IsLessThan", func(t *testing.T) {
		where := NewWhere()
		where.Set("age", 65, &SetWhereOptions{IsLessThan: true})

		query, args, err := where.Build()
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}

		expectedQuery := "age < ?"
		if query != expectedQuery {
			t.Errorf("Expected query = %v, got %v", expectedQuery, query)
		}

		if len(args) != 1 {
			t.Errorf("Expected 1 arg, got %d", len(args))
		}

		if args[0] != 65 {
			t.Errorf("Expected arg 65, got %v", args[0])
		}
	})

	// Test IsGreaterOrEqualThan
	t.Run("IsGreaterOrEqualThan", func(t *testing.T) {
		where := NewWhere()
		where.Set("price", 100.0, &SetWhereOptions{IsGreaterOrEqualThan: true})

		query, args, err := where.Build()
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}

		expectedQuery := "price >= ?"
		if query != expectedQuery {
			t.Errorf("Expected query = %v, got %v", expectedQuery, query)
		}

		if len(args) != 1 {
			t.Errorf("Expected 1 arg, got %d", len(args))
		}

		if args[0] != 100.0 {
			t.Errorf("Expected arg 100.0, got %v", args[0])
		}
	})

	// Test IsLessOrEqualThan
	t.Run("IsLessOrEqualThan", func(t *testing.T) {
		where := NewWhere()
		where.Set("score", 90, &SetWhereOptions{IsLessOrEqualThan: true})

		query, args, err := where.Build()
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}

		expectedQuery := "score <= ?"
		if query != expectedQuery {
			t.Errorf("Expected query = %v, got %v", expectedQuery, query)
		}

		if len(args) != 1 {
			t.Errorf("Expected 1 arg, got %d", len(args))
		}

		if args[0] != 90 {
			t.Errorf("Expected arg 90, got %v", args[0])
		}
	})

	// Test with string date
	t.Run("GreaterThan with date string", func(t *testing.T) {
		where := NewWhere()
		where.Set("created_at", "2023-01-01", &SetWhereOptions{IsGreaterThan: true})

		query, args, err := where.Build()
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}

		expectedQuery := "created_at > ?"
		if query != expectedQuery {
			t.Errorf("Expected query = %v, got %v", expectedQuery, query)
		}

		if len(args) != 1 {
			t.Errorf("Expected 1 arg, got %d", len(args))
		}

		if args[0] != "2023-01-01" {
			t.Errorf("Expected arg 2023-01-01, got %v", args[0])
		}
	})

	// Test with time.Time
	t.Run("LessThan with time.Time", func(t *testing.T) {
		where := NewWhere()
		date := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)
		where.Set("created_at", date, &SetWhereOptions{IsLessThan: true})

		query, args, err := where.Build()
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}

		expectedQuery := "created_at < ?"
		if query != expectedQuery {
			t.Errorf("Expected query = %v, got %v", expectedQuery, query)
		}

		if len(args) != 1 {
			t.Errorf("Expected 1 arg, got %d", len(args))
		}

		if args[0] != date {
			t.Errorf("Expected arg %v, got %v", date, args[0])
		}
	})

	// Test multiple conditions with comparisons
	t.Run("Multiple comparison conditions", func(t *testing.T) {
		where := NewWhere()
		// Use Add instead of Set to allow multiple conditions on the same field
		where.Add("age", 18, &SetWhereOptions{IsGreaterOrEqualThan: true})
		where.Add("age", 65, &SetWhereOptions{IsLessOrEqualThan: true})
		where.Set("status", "active")

		query, args, err := where.Build()
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}

		expectedQuery := "age >= ? AND age <= ? AND status = ?"
		if query != expectedQuery {
			t.Errorf("Expected query = %v, got %v", expectedQuery, query)
		}

		if len(args) != 3 {
			t.Errorf("Expected 3 args, got %d", len(args))
		}

		if args[0] != 18 || args[1] != 65 || args[2] != "active" {
			t.Errorf("Expected args [18, 65, active], got %v", args)
		}
	})

	// Test mixing comparison with other conditions
	t.Run("Mix comparison with other conditions", func(t *testing.T) {
		where := NewWhere()
		where.Set("price", 100.0, &SetWhereOptions{IsGreaterThan: true})
		where.Set("name", "Product", &SetWhereOptions{IsFuzzy: true})
		where.Set("category", []string{"Electronics", "Books"}, &SetWhereOptions{IsIn: true})

		query, args, err := where.Build()
		if err != nil {
			t.Fatalf("Build() error = %v", err)
		}

		expectedQuery := "price > ? AND name ILike ? AND category in (?)"
		if query != expectedQuery {
			t.Errorf("Expected query = %v, got %v", expectedQuery, query)
		}

		if len(args) != 3 {
			t.Errorf("Expected 3 args, got %d", len(args))
		}
	})

	// Test all comparison operators together
	t.Run("All comparison operators", func(t *testing.T) {
		where1 := NewWhere()
		where1.Set("a", 1, &SetWhereOptions{IsGreaterThan: true})
		query1, _, _ := where1.Build()
		if query1 != "a > ?" {
			t.Errorf("IsGreaterThan failed: %v", query1)
		}

		where2 := NewWhere()
		where2.Set("b", 2, &SetWhereOptions{IsLessThan: true})
		query2, _, _ := where2.Build()
		if query2 != "b < ?" {
			t.Errorf("IsLessThan failed: %v", query2)
		}

		where3 := NewWhere()
		where3.Set("c", 3, &SetWhereOptions{IsGreaterOrEqualThan: true})
		query3, _, _ := where3.Build()
		if query3 != "c >= ?" {
			t.Errorf("IsGreaterOrEqualThan failed: %v", query3)
		}

		where4 := NewWhere()
		where4.Set("d", 4, &SetWhereOptions{IsLessOrEqualThan: true})
		query4, _, _ := where4.Build()
		if query4 != "d <= ?" {
			t.Errorf("IsLessOrEqualThan failed: %v", query4)
		}
	})
}
