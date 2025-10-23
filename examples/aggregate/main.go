package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-zoox/gormx"
	"gorm.io/gorm"
)

// Product represents a product model
type Product struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"column:name" json:"name"`
	Category  string         `gorm:"column:category" json:"category"`
	Price     float64        `gorm:"column:price" json:"price"`
	Quantity  int            `gorm:"column:quantity" json:"quantity"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (Product) TableName() string {
	return "products"
}

func main() {
	// Initialize database connection (you need to set up your database connection)
	// Example: err := gormx.LoadDB("mysql", "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
	// if err != nil {
	//     log.Fatal(err)
	// }

	// For demonstration purposes, we'll show the API usage
	fmt.Println("GORMX Aggregate Query Examples")
	fmt.Println("==============================")

	// Example 1: Basic Sum query
	fmt.Println("\n1. Sum all product prices:")
	// total, err := gormx.Sum[Product]("price", nil)
	// if err != nil {
	//     log.Printf("Sum error: %v", err)
	// } else {
	//     fmt.Printf("Total price: %.2f\n", total)
	// }

	// Example 2: Average price
	fmt.Println("\n2. Average product price:")
	// avgPrice, err := gormx.Avg[Product]("price", nil)
	// if err != nil {
	//     log.Printf("Avg error: %v", err)
	// } else {
	//     fmt.Printf("Average price: %.2f\n", avgPrice)
	// }

	// Example 3: Min and Max prices
	fmt.Println("\n3. Min and Max prices:")
	// minPrice, err := gormx.Min[Product]("price", nil)
	// if err != nil {
	//     log.Printf("Min error: %v", err)
	// } else {
	//     fmt.Printf("Minimum price: %v\n", minPrice)
	// }

	// maxPrice, err := gormx.Max[Product]("price", nil)
	// if err != nil {
	//     log.Printf("Max error: %v", err)
	// } else {
	//     fmt.Printf("Maximum price: %v\n", maxPrice)
	// }

	// Example 4: Count distinct categories
	fmt.Println("\n4. Count distinct categories:")
	// categoryCount, err := gormx.CountDistinct[Product]("category", nil)
	// if err != nil {
	//     log.Printf("CountDistinct error: %v", err)
	// } else {
	//     fmt.Printf("Number of distinct categories: %d\n", categoryCount)
	// }

	// Example 5: Sum with conditions
	fmt.Println("\n5. Sum prices for Electronics category:")
	where := gormx.NewWhere()
	where.Set("category", "Electronics")

	// electronicsTotal, err := gormx.Sum[Product]("price", where)
	// if err != nil {
	//     log.Printf("Sum with where error: %v", err)
	// } else {
	//     fmt.Printf("Electronics total price: %.2f\n", electronicsTotal)
	// }

	// Example 6: Group by category with aggregations
	fmt.Println("\n6. Group by category with count and sum:")
	// results, err := gormx.GroupBy[Product](
	//     []string{"category"},
	//     nil,
	//     []string{"COUNT(*) as count", "SUM(price) as sum", "AVG(price) as avg"},
	// )
	// if err != nil {
	//     log.Printf("GroupBy error: %v", err)
	// } else {
	//     for _, result := range results {
	//         fmt.Printf("Category: %v, Count: %d, Sum: %.2f, Avg: %.2f\n",
	//             result.Group["category"], result.Count, result.Sum, result.Avg)
	//     }
	// }

	// Example 7: Multiple aggregations in one query
	fmt.Println("\n7. Multiple aggregations:")
	// aggregates, err := gormx.Aggregate[Product](
	//     "price",
	//     nil,
	//     []string{"sum", "avg", "min", "max", "count"},
	// )
	// if err != nil {
	//     log.Printf("Aggregate error: %v", err)
	// } else {
	//     fmt.Printf("Price Statistics:\n")
	//     fmt.Printf("  Sum: %v\n", aggregates["sum"])
	//     fmt.Printf("  Average: %v\n", aggregates["avg"])
	//     fmt.Printf("  Min: %v\n", aggregates["min"])
	//     fmt.Printf("  Max: %v\n", aggregates["max"])
	//     fmt.Printf("  Count: %v\n", aggregates["count"])
	// }

	// Example 8: Complex where conditions
	fmt.Println("\n8. Complex where conditions:")
	complexWhere := gormx.NewWhere()
	complexWhere.Set("category", "Electronics")
	complexWhere.Set("price", 100, &gormx.SetWhereOptions{IsFuzzy: true}) // price > 100

	// expensiveElectronics, err := gormx.Sum[Product]("price", complexWhere)
	// if err != nil {
	//     log.Printf("Complex where error: %v", err)
	// } else {
	//     fmt.Printf("Expensive electronics total: %.2f\n", expensiveElectronics)
	// }

	fmt.Println("\nNote: Uncomment the code above and set up your database connection to run these examples.")
}

// Example usage in a web handler
func handleProductStats() {
	// This is an example of how you might use aggregate queries in a web handler

	// Get total revenue
	totalRevenue, err := gormx.Sum[Product]("price", nil)
	if err != nil {
		log.Printf("Failed to get total revenue: %v", err)
		return
	}

	// Get category breakdown
	categoryStats, err := gormx.GroupBy[Product](
		[]string{"category"},
		nil,
		[]string{"COUNT(*) as count", "SUM(price) as total", "AVG(price) as average"},
	)
	if err != nil {
		log.Printf("Failed to get category stats: %v", err)
		return
	}

	// Process results
	fmt.Printf("Total Revenue: %.2f\n", totalRevenue)
	fmt.Println("Category Breakdown:")
	for _, stat := range categoryStats {
		category := stat.Group["category"]
		fmt.Printf("  %s: %d products, total: %.2f, average: %.2f\n",
			category, stat.Count, stat.Sum, stat.Avg)
	}
}
