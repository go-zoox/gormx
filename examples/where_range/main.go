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
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func main() {
	// This example demonstrates the IsRange feature for WHERE conditions
	// For actual database operations, you need to set up a database connection
	// Example: err := gormx.LoadDB("mysql", "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")

	fmt.Println("GORMX Range Query Examples")
	fmt.Println("===========================")
	fmt.Println()

	// Example 1: Age Range (Integer)
	fmt.Println("=== Example 1: Age Range (Integer) ===")
	where1 := gormx.NewWhere()
	where1.Set("age", []int{28, 38}, &gormx.SetWhereOptions{IsRange: true})

	query1, args1, _ := where1.Build()
	fmt.Printf("Query: %s\n", query1)
	fmt.Printf("Args: %v\n", args1)
	fmt.Println("SQL: SELECT * FROM users WHERE age BETWEEN 28 AND 38")
	fmt.Println()

	// Example 2: Date Range (String)
	fmt.Println("=== Example 2: Date Range (String) ===")
	where2 := gormx.NewWhere()
	where2.Set("created_at", []string{"2023-03-01", "2023-09-30"}, &gormx.SetWhereOptions{IsRange: true})

	query2, args2, _ := where2.Build()
	fmt.Printf("Query: %s\n", query2)
	fmt.Printf("Args: %v\n", args2)
	fmt.Println("SQL: SELECT * FROM users WHERE created_at BETWEEN '2023-03-01' AND '2023-09-30'")
	fmt.Println()

	// Example 3: Date Range (time.Time)
	fmt.Println("=== Example 3: Date Range (time.Time) ===")
	start := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2023, 6, 30, 23, 59, 59, 0, time.UTC)
	where3 := gormx.NewWhere()
	where3.Set("created_at", []interface{}{start, end}, &gormx.SetWhereOptions{IsRange: true})

	query3, _, _ := where3.Build()
	fmt.Printf("Query: %s\n", query3)
	fmt.Printf("Args: [%v, %v]\n", start.Format("2006-01-02"), end.Format("2006-01-02"))
	fmt.Println()

	// Example 4: Multiple Conditions with Range
	fmt.Println("=== Example 4: Multiple Conditions with Range ===")
	where4 := gormx.NewWhere()
	where4.Set("age", []int{25, 40}, &gormx.SetWhereOptions{IsRange: true})
	where4.Set("created_at", []string{"2023-01-01", "2023-12-31"}, &gormx.SetWhereOptions{IsRange: true})

	query4, args4, _ := where4.Build()
	fmt.Printf("Query: %s\n", query4)
	fmt.Printf("Args: %v\n", args4)
	fmt.Println("SQL: SELECT * FROM users WHERE age BETWEEN 25 AND 40 AND created_at BETWEEN '2023-01-01' AND '2023-12-31'")
	fmt.Println()

	// Example 5: Range with Other Conditions
	fmt.Println("=== Example 5: Range with Other Conditions ===")
	where5 := gormx.NewWhere()
	where5.Set("age", []int{30, 45}, &gormx.SetWhereOptions{IsRange: true})
	where5.Set("name", "e", &gormx.SetWhereOptions{IsFuzzy: true})

	query5, args5, _ := where5.Build()
	fmt.Printf("Query: %s\n", query5)
	fmt.Printf("Args: %v\n", args5)
	fmt.Printf("SQL: SELECT * FROM users WHERE age BETWEEN 30 AND 45 AND name ILike '%%e%%'\n")
	fmt.Println()

	// Example 6: Price Range (Float)
	fmt.Println("=== Example 6: Price Range (Float) ===")
	where6 := gormx.NewWhere()
	where6.Set("price", []float64{10.5, 99.99}, &gormx.SetWhereOptions{IsRange: true})

	query6, args6, _ := where6.Build()
	fmt.Printf("Query: %s\n", query6)
	fmt.Printf("Args: %v\n", args6)
	fmt.Println("SQL: SELECT * FROM products WHERE price BETWEEN 10.5 AND 99.99")
	fmt.Println()

	// Example 7: Usage in actual queries (commented, requires DB connection)
	fmt.Println("=== Example 7: Usage in Actual Queries ===")
	fmt.Println("// Find users with age between 18 and 65")
	fmt.Println("// where := gormx.NewWhere()")
	fmt.Println("// where.Set(\"age\", []int{18, 65}, &gormx.SetWhereOptions{IsRange: true})")
	fmt.Println("// users, total, err := gormx.Find[User](1, 10, where, nil)")
	fmt.Println()

	// Example 8: URL Query Parameter Syntax
	fmt.Println("=== Example 8: URL Query Parameter Syntax ===")
	fmt.Println("In HTTP requests, you can use the following syntax:")
	fmt.Println("  ?age=18,65:range")
	fmt.Println("  ?created_at=2023-01-01,2023-12-31:range")
	fmt.Println("  ?price=10.5,99.99:range")
	fmt.Println()
	fmt.Println("Example in HTTP handler:")
	fmt.Println("  params := gormx.NewParams(ctx)")
	fmt.Println("  where := params.Where()")
	fmt.Println("  // URL: /users?age=18,65:range&status=active")
	fmt.Println("  // Will generate: WHERE age BETWEEN 18 AND 65 AND status = 'active'")
	fmt.Println()

	fmt.Println("\nNote: Uncomment the query examples and set up your database connection to execute actual queries.")
}

// ExampleWithDB demonstrates usage with a database connection
func ExampleWithDB() {
	// This is an example of how you might use range queries with an actual database

	// Find users within age range
	// where := gormx.NewWhere()
	// where.Set("age", []int{18, 65}, &gormx.SetWhereOptions{IsRange: true})
	// users, total, err := gormx.Find[User](1, 10, where, nil)
	// if err != nil {
	//     log.Printf("Failed to find users: %v", err)
	//     return
	// }
	// fmt.Printf("Found %d users (total: %d)\n", len(users), total)

	// Find orders within date range
	// orderWhere := gormx.NewWhere()
	// orderWhere.Set("order_date", []string{"2023-01-01", "2023-12-31"}, &gormx.SetWhereOptions{IsRange: true})
	// orderWhere.Set("status", "completed")
	// orders, _, err := gormx.Find[Order](1, 50, orderWhere, nil)

	// Find products within price range
	// productWhere := gormx.NewWhere()
	// productWhere.Set("price", []float64{10.0, 100.0}, &gormx.SetWhereOptions{IsRange: true})
	// products, _, err := gormx.Find[Product](1, 20, productWhere, nil)
}
