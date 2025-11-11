package main

import (
	"fmt"
	"time"

	"github.com/go-zoox/gormx"
)

// User is the model
type User struct {
	ID        uint      `gorm:"primarykey"`
	Name      string    `gorm:"size:255"`
	Age       int       `gorm:"default:0"`
	Score     float64   `gorm:"default:0"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// Product is the model
type Product struct {
	ID       uint    `gorm:"primarykey"`
	Name     string  `gorm:"size:255"`
	Price    float64 `gorm:"default:0"`
	Quantity int     `gorm:"default:0"`
	Category string  `gorm:"size:100"`
}

func main() {
	// This example demonstrates the comparison operators for WHERE conditions
	// For actual database operations, you need to set up a database connection
	// Example: err := gormx.LoadDB("mysql", "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")

	fmt.Println("GORMX Comparison Operators Examples")
	fmt.Println("====================================")
	fmt.Println()

	// Example 1: Greater Than (>)
	fmt.Println("=== Example 1: Greater Than (>) ===")
	where1 := gormx.NewWhere()
	where1.Set("age", 18, &gormx.SetWhereOptions{IsGreaterThan: true})

	query1, args1, _ := where1.Build()
	fmt.Printf("Query: %s\n", query1)
	fmt.Printf("Args: %v\n", args1)
	fmt.Println("SQL: SELECT * FROM users WHERE age > 18")
	fmt.Println()

	// Example 2: Less Than (<)
	fmt.Println("=== Example 2: Less Than (<) ===")
	where2 := gormx.NewWhere()
	where2.Set("age", 65, &gormx.SetWhereOptions{IsLessThan: true})

	query2, args2, _ := where2.Build()
	fmt.Printf("Query: %s\n", query2)
	fmt.Printf("Args: %v\n", args2)
	fmt.Println("SQL: SELECT * FROM users WHERE age < 65")
	fmt.Println()

	// Example 3: Greater Or Equal (>=)
	fmt.Println("=== Example 3: Greater Or Equal (>=) ===")
	where3 := gormx.NewWhere()
	where3.Set("score", 60.0, &gormx.SetWhereOptions{IsGreaterOrEqualThan: true})

	query3, args3, _ := where3.Build()
	fmt.Printf("Query: %s\n", query3)
	fmt.Printf("Args: %v\n", args3)
	fmt.Println("SQL: SELECT * FROM users WHERE score >= 60.0")
	fmt.Println()

	// Example 4: Less Or Equal (<=)
	fmt.Println("=== Example 4: Less Or Equal (<=) ===")
	where4 := gormx.NewWhere()
	where4.Set("price", 999.99, &gormx.SetWhereOptions{IsLessOrEqualThan: true})

	query4, args4, _ := where4.Build()
	fmt.Printf("Query: %s\n", query4)
	fmt.Printf("Args: %v\n", args4)
	fmt.Println("SQL: SELECT * FROM products WHERE price <= 999.99")
	fmt.Println()

	// Example 5: Multiple Comparison Conditions (Age Range using >= and <=)
	fmt.Println("=== Example 5: Multiple Comparison Conditions ===")
	where5 := gormx.NewWhere()
	// Use Add to allow multiple conditions on the same field
	where5.Add("age", 18, &gormx.SetWhereOptions{IsGreaterOrEqualThan: true})
	where5.Add("age", 65, &gormx.SetWhereOptions{IsLessOrEqualThan: true})

	query5, args5, _ := where5.Build()
	fmt.Printf("Query: %s\n", query5)
	fmt.Printf("Args: %v\n", args5)
	fmt.Println("SQL: SELECT * FROM users WHERE age >= 18 AND age <= 65")
	fmt.Println("Note: For simple ranges, consider using IsRange instead")
	fmt.Println()

	// Example 6: Comparison with Date (String)
	fmt.Println("=== Example 6: Comparison with Date (String) ===")
	where6 := gormx.NewWhere()
	where6.Set("created_at", "2023-01-01", &gormx.SetWhereOptions{IsGreaterThan: true})

	query6, args6, _ := where6.Build()
	fmt.Printf("Query: %s\n", query6)
	fmt.Printf("Args: %v\n", args6)
	fmt.Println("SQL: SELECT * FROM users WHERE created_at > '2023-01-01'")
	fmt.Println()

	// Example 7: Comparison with time.Time
	fmt.Println("=== Example 7: Comparison with time.Time ===")
	date := time.Date(2023, 6, 30, 23, 59, 59, 0, time.UTC)
	where7 := gormx.NewWhere()
	where7.Set("created_at", date, &gormx.SetWhereOptions{IsLessThan: true})

	query7, _, _ := where7.Build()
	fmt.Printf("Query: %s\n", query7)
	fmt.Printf("Args: [%s]\n", date.Format("2006-01-02"))
	fmt.Println()

	// Example 8: Mixing Comparison with Other Conditions
	fmt.Println("=== Example 8: Mixing Comparison with Other Conditions ===")
	where8 := gormx.NewWhere()
	where8.Set("price", 100.0, &gormx.SetWhereOptions{IsGreaterThan: true})
	where8.Set("name", "Phone", &gormx.SetWhereOptions{IsFuzzy: true})
	where8.Set("category", []string{"Electronics", "Gadgets"}, &gormx.SetWhereOptions{IsIn: true})

	query8, args8, _ := where8.Build()
	fmt.Printf("Query: %s\n", query8)
	fmt.Printf("Args: %v\n", args8)
	fmt.Println("SQL: SELECT * FROM products WHERE price > 100.0 AND name ILike '%Phone%' AND category IN ('Electronics', 'Gadgets')")
	fmt.Println()

	// Example 9: Stock Management Query
	fmt.Println("=== Example 9: Stock Management Query ===")
	where9 := gormx.NewWhere()
	where9.Set("quantity", 10, &gormx.SetWhereOptions{IsLessThan: true})
	where9.Set("price", 0, &gormx.SetWhereOptions{IsGreaterThan: true})

	query9, args9, _ := where9.Build()
	fmt.Printf("Query: %s\n", query9)
	fmt.Printf("Args: %v\n", args9)
	fmt.Println("SQL: SELECT * FROM products WHERE quantity < 10 AND price > 0")
	fmt.Println("Use case: Find low-stock products with valid prices")
	fmt.Println()

	// Example 10: URL Query Parameter Syntax
	fmt.Println("=== Example 10: URL Query Parameter Syntax ===")
	fmt.Println("In HTTP requests, you can use the following syntax:")
	fmt.Println("  ?age=18:>          (age > 18)")
	fmt.Println("  ?age=65:<          (age < 65)")
	fmt.Println("  ?score=60:>=       (score >= 60)")
	fmt.Println("  ?price=100:<=      (price <= 100)")
	fmt.Println()
	fmt.Println("Example in HTTP handler:")
	fmt.Println("  params := gormx.NewParams(ctx)")
	fmt.Println("  where := params.Where()")
	fmt.Println("  // URL: /users?age=18:>&score=60:>=&status=active")
	fmt.Println("  // Will generate: WHERE age > 18 AND score >= 60 AND status = 'active'")
	fmt.Println()

	// Example 11: Real-world Use Cases
	fmt.Println("=== Example 11: Real-world Use Cases ===")
	fmt.Println()

	// Use case 1: Adult users
	fmt.Println("Use Case 1: Find adult users (age >= 18)")
	whereAdult := gormx.NewWhere()
	whereAdult.Set("age", 18, &gormx.SetWhereOptions{IsGreaterOrEqualThan: true})
	queryAdult, _, _ := whereAdult.Build()
	fmt.Printf("  Query: %s\n", queryAdult)
	fmt.Println()

	// Use case 2: Recent orders
	fmt.Println("Use Case 2: Find recent orders (created after 2023-01-01)")
	whereRecent := gormx.NewWhere()
	whereRecent.Set("created_at", "2023-01-01", &gormx.SetWhereOptions{IsGreaterThan: true})
	queryRecent, _, _ := whereRecent.Build()
	fmt.Printf("  Query: %s\n", queryRecent)
	fmt.Println()

	// Use case 3: Budget products
	fmt.Println("Use Case 3: Find budget products (price <= 50)")
	whereBudget := gormx.NewWhere()
	whereBudget.Set("price", 50.0, &gormx.SetWhereOptions{IsLessOrEqualThan: true})
	queryBudget, _, _ := whereBudget.Build()
	fmt.Printf("  Query: %s\n", queryBudget)
	fmt.Println()

	// Use case 4: High performers
	fmt.Println("Use Case 4: Find high performers (score > 90)")
	whereHighScore := gormx.NewWhere()
	whereHighScore.Set("score", 90.0, &gormx.SetWhereOptions{IsGreaterThan: true})
	queryHighScore, _, _ := whereHighScore.Build()
	fmt.Printf("  Query: %s\n", queryHighScore)
	fmt.Println()

	fmt.Println("\nNote: Uncomment the query examples and set up your database connection to execute actual queries.")
}

// ExampleWithDB demonstrates usage with a database connection
func ExampleWithDB() {
	// This is an example of how you might use comparison queries with an actual database

	// Find adult users (age >= 18)
	// where := gormx.NewWhere()
	// where.Set("age", 18, &gormx.SetWhereOptions{IsGreaterOrEqualThan: true})
	// users, total, err := gormx.Find[User](1, 10, where, nil)

	// Find high-value products (price > 1000)
	// productWhere := gormx.NewWhere()
	// productWhere.Set("price", 1000.0, &gormx.SetWhereOptions{IsGreaterThan: true})
	// products, _, err := gormx.Find[Product](1, 20, productWhere, nil)

	// Find recent users (created after a specific date)
	// recentWhere := gormx.NewWhere()
	// recentWhere.Set("created_at", "2023-01-01", &gormx.SetWhereOptions{IsGreaterThan: true})
	// recentUsers, _, err := gormx.Find[User](1, 50, recentWhere, nil)

	// Complex query: Find affordable electronics with good ratings
	// complexWhere := gormx.NewWhere()
	// complexWhere.Set("price", 500.0, &gormx.SetWhereOptions{IsLessOrEqualThan: true})
	// complexWhere.Set("score", 4.0, &gormx.SetWhereOptions{IsGreaterOrEqualThan: true})
	// complexWhere.Set("category", "Electronics")
	// result, _, err := gormx.Find[Product](1, 20, complexWhere, nil)
}

