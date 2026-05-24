package main

import (
	"fmt"

	"github.com/go-zoox/gormx"
)

func main() {
	fmt.Println("GORMX Range Modes Examples")
	fmt.Println("===========================")
	fmt.Println()

	// Example 1: Closed Range [A, B] (Default)
	fmt.Println("=== Example 1: Closed Range [A, B] (Default) ===")
	where1 := gormx.NewWhere()
	where1.Set("age", []int{18, 65}, &gormx.SetWhereOptions{
		IsRange:   true,
		RangeMode: gormx.RangeModeClosed,
	})

	query1, args1, _ := where1.Build()
	fmt.Printf("Mode: [A, B] (Closed - both inclusive)\n")
	fmt.Printf("Query: %s\n", query1)
	fmt.Printf("Args: %v\n", args1)
	fmt.Println("SQL: SELECT * FROM users WHERE age BETWEEN 18 AND 65")
	fmt.Println("Meaning: 18 <= age <= 65")
	fmt.Println()

	// Example 2: Default (no mode specified, should be Closed)
	fmt.Println("=== Example 2: Default Mode (No RangeMode specified) ===")
	where2 := gormx.NewWhere()
	where2.Set("age", []int{18, 65}, &gormx.SetWhereOptions{IsRange: true})

	query2, args2, _ := where2.Build()
	fmt.Printf("Query: %s\n", query2)
	fmt.Printf("Args: %v\n", args2)
	fmt.Println("Note: Default mode is Closed [A, B]")
	fmt.Println()

	// Example 3: Open Range (A, B)
	fmt.Println("=== Example 3: Open Range (A, B) ===")
	where3 := gormx.NewWhere()
	where3.Set("age", []int{18, 65}, &gormx.SetWhereOptions{
		IsRange:   true,
		RangeMode: gormx.RangeModeOpen,
	})

	query3, args3, _ := where3.Build()
	fmt.Printf("Mode: (A, B) (Open - both exclusive)\n")
	fmt.Printf("Query: %s\n", query3)
	fmt.Printf("Args: %v\n", args3)
	fmt.Println("SQL: SELECT * FROM users WHERE (age > 18 AND age < 65)")
	fmt.Println("Meaning: 18 < age < 65 (excludes 18 and 65)")
	fmt.Println()

	// Example 4: Left Closed Range [A, B)
	fmt.Println("=== Example 4: Left Closed Range [A, B) ===")
	where4 := gormx.NewWhere()
	where4.Set("age", []int{18, 65}, &gormx.SetWhereOptions{
		IsRange:   true,
		RangeMode: gormx.RangeModeLeftClosed,
	})

	query4, args4, _ := where4.Build()
	fmt.Printf("Mode: [A, B) (Left closed, right open)\n")
	fmt.Printf("Query: %s\n", query4)
	fmt.Printf("Args: %v\n", args4)
	fmt.Println("SQL: SELECT * FROM users WHERE (age >= 18 AND age < 65)")
	fmt.Println("Meaning: 18 <= age < 65 (includes 18, excludes 65)")
	fmt.Println()

	// Example 5: Right Closed Range (A, B]
	fmt.Println("=== Example 5: Right Closed Range (A, B] ===")
	where5 := gormx.NewWhere()
	where5.Set("age", []int{18, 65}, &gormx.SetWhereOptions{
		IsRange:   true,
		RangeMode: gormx.RangeModeRightClosed,
	})

	query5, args5, _ := where5.Build()
	fmt.Printf("Mode: (A, B] (Left open, right closed)\n")
	fmt.Printf("Query: %s\n", query5)
	fmt.Printf("Args: %v\n", args5)
	fmt.Println("SQL: SELECT * FROM users WHERE (age > 18 AND age <= 65)")
	fmt.Println("Meaning: 18 < age <= 65 (excludes 18, includes 65)")
	fmt.Println()

	// Example 6: Price Range with Open mode
	fmt.Println("=== Example 6: Price Range with Open Mode ===")
	where6 := gormx.NewWhere()
	where6.Set("price", []float64{100.0, 999.99}, &gormx.SetWhereOptions{
		IsRange:   true,
		RangeMode: gormx.RangeModeOpen,
	})

	query6, args6, _ := where6.Build()
	fmt.Printf("Query: %s\n", query6)
	fmt.Printf("Args: %v\n", args6)
	fmt.Println("SQL: SELECT * FROM products WHERE (price > 100.0 AND price < 999.99)")
	fmt.Println()

	// Example 7: Date Range with Left Closed mode
	fmt.Println("=== Example 7: Date Range with Left Closed Mode ===")
	where7 := gormx.NewWhere()
	where7.Set("created_at", []string{"2023-01-01", "2023-12-31"}, &gormx.SetWhereOptions{
		IsRange:   true,
		RangeMode: gormx.RangeModeLeftClosed,
	})

	query7, args7, _ := where7.Build()
	fmt.Printf("Query: %s\n", query7)
	fmt.Printf("Args: %v\n", args7)
	fmt.Println("SQL: SELECT * FROM orders WHERE (created_at >= '2023-01-01' AND created_at < '2023-12-31')")
	fmt.Println("Use case: Get all data from Jan 1 to Dec 30 (useful for yearly reports)")
	fmt.Println()

	// Example 8: Multiple conditions with different range modes
	fmt.Println("=== Example 8: Multiple Conditions with Different Modes ===")
	where8 := gormx.NewWhere()
	where8.Add("age", []int{18, 30}, &gormx.SetWhereOptions{
		IsRange:   true,
		RangeMode: gormx.RangeModeLeftClosed,
	})
	where8.Add("score", []int{60, 100}, &gormx.SetWhereOptions{
		IsRange:   true,
		RangeMode: gormx.RangeModeClosed,
	})
	where8.Set("status", "active")

	query8, args8, _ := where8.Build()
	fmt.Printf("Query: %s\n", query8)
	fmt.Printf("Args: %v\n", args8)
	fmt.Println("SQL: SELECT * FROM students")
	fmt.Println("     WHERE (age >= 18 AND age < 30)")
	fmt.Println("     AND score BETWEEN 60 AND 100")
	fmt.Println("     AND status = 'active'")
	fmt.Println()

	// Example 9: Comparison of all four modes
	fmt.Println("=== Example 9: Comparison of All Four Modes ===")
	fmt.Println("For range [18, 65]:")
	fmt.Println()

	modes := []struct {
		name     string
		mode     gormx.RangeMode
		notation string
		meaning  string
	}{
		{"Closed", gormx.RangeModeClosed, "[18, 65]", "18 <= age <= 65"},
		{"Open", gormx.RangeModeOpen, "(18, 65)", "18 < age < 65"},
		{"Left Closed", gormx.RangeModeLeftClosed, "[18, 65)", "18 <= age < 65"},
		{"Right Closed", gormx.RangeModeRightClosed, "(18, 65]", "18 < age <= 65"},
	}

	for i, m := range modes {
		where := gormx.NewWhere()
		where.Set("age", []int{18, 65}, &gormx.SetWhereOptions{
			IsRange:   true,
			RangeMode: m.mode,
		})
		query, _, _ := where.Build()
		fmt.Printf("%d. %s %s\n", i+1, m.name, m.notation)
		fmt.Printf("   Meaning: %s\n", m.meaning)
		fmt.Printf("   Query: %s\n", query)
		fmt.Println()
	}

	// Example 10: Real-world Use Cases
	fmt.Println("=== Example 10: Real-world Use Cases ===")
	fmt.Println()

	fmt.Println("Use Case 1: Adult age range (18 <= age <= 65)")
	fmt.Println("  Use RangeModeClosed: [18, 65]")
	fmt.Println()

	fmt.Println("Use Case 2: Exclude boundary values (age between but not equal to 18 and 65)")
	fmt.Println("  Use RangeModeOpen: (18, 65)")
	fmt.Println()

	fmt.Println("Use Case 3: Time range (from start of day, before end of day)")
	fmt.Println("  Use RangeModeLeftClosed: [2023-01-01 00:00:00, 2023-01-02 00:00:00)")
	fmt.Println("  This includes all records on 2023-01-01")
	fmt.Println()

	fmt.Println("Use Case 4: Pagination offset (skip first N, include up to M)")
	fmt.Println("  Use RangeModeRightClosed: (0, 100]")
	fmt.Println("  This includes records 1-100, excluding 0")
	fmt.Println()

	fmt.Println("\n=== Example 11: URL Query Parameter Syntax ===")
	fmt.Println()
	fmt.Println("In HTTP requests, you can use the following simplified syntax:")
	fmt.Println()
	fmt.Println("1. Default Range (Closed) - :range or :range[]")
	fmt.Println("   ?age=18,65:range")
	fmt.Println("   ?age=18,65:range[]")
	fmt.Println("   Generates: WHERE age BETWEEN 18 AND 65")
	fmt.Println("   Meaning: [18, 65] - includes both 18 and 65")
	fmt.Println()
	fmt.Println("2. Open Range - :range()")
	fmt.Println("   ?age=18,65:range()")
	fmt.Println("   Generates: WHERE (age > 18 AND age < 65)")
	fmt.Println("   Meaning: (18, 65) - excludes both 18 and 65")
	fmt.Println()
	fmt.Println("3. Left Closed Range - :range[)")
	fmt.Println("   ?created_at=2023-01-01,2023-12-31:range[)")
	fmt.Println("   Generates: WHERE (created_at >= '2023-01-01' AND created_at < '2023-12-31')")
	fmt.Println("   Meaning: [2023-01-01, 2023-12-31) - includes start, excludes end")
	fmt.Println()
	fmt.Println("4. Right Closed Range - :range(]")
	fmt.Println("   ?score=0,100:range(]")
	fmt.Println("   Generates: WHERE (score > 0 AND score <= 100)")
	fmt.Println("   Meaning: (0, 100] - excludes start, includes end")
	fmt.Println()
	fmt.Println("Summary:")
	fmt.Println("  :range or :range[]  →  [A, B]  (both inclusive)")
	fmt.Println("  :range()            →  (A, B)  (both exclusive)")
	fmt.Println("  :range[)            →  [A, B)  (left inclusive, right exclusive)")
	fmt.Println("  :range(]            →  (A, B]  (left exclusive, right inclusive)")
	fmt.Println()
	fmt.Println("Complete example in HTTP handler:")
	fmt.Println("  params := gormx.NewParams(ctx)")
	fmt.Println("  where := params.Where()")
	fmt.Println("  // URL: /users?age=18,65:range[)&status=active")
	fmt.Println("  // Generates: WHERE (age >= 18 AND age < 65) AND status = 'active'")
	fmt.Println()
}
