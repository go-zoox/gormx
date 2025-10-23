package main

import (
	"fmt"
	"time"

	_ "github.com/go-zoox/gormx" // imported for examples
	"gorm.io/gorm"
)

// ChainProduct represents a product model for chain examples
type ChainProduct struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"column:name" json:"name"`
	Category  string         `gorm:"column:category" json:"category"`
	Price     float64        `gorm:"column:price" json:"price"`
	Quantity  int            `gorm:"column:quantity" json:"quantity"`
	InStock   bool           `gorm:"column:in_stock" json:"in_stock"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (ChainProduct) TableName() string {
	return "products"
}

func main() {
	fmt.Println("GORMX Chain Query Examples")
	fmt.Println("==========================")
	fmt.Println()

	// Note: Initialize your database connection before using these examples
	// Example: err := gormx.LoadDB("mysql", "user:password@tcp(localhost:3306)/dbname")

	// Example 1: Simple Query
	fmt.Println("1. Simple Query with Where:")
	example1()

	// Example 2: Multiple Where Conditions
	fmt.Println("\n2. Multiple Where Conditions:")
	example2()

	// Example 3: Ordering and Limiting
	fmt.Println("\n3. Ordering and Limiting:")
	example3()

	// Example 4: Pagination
	fmt.Println("\n4. Pagination:")
	example4()

	// Example 5: Aggregate Functions
	fmt.Println("\n5. Aggregate Functions:")
	example5()

	// Example 6: Complex Chained Query
	fmt.Println("\n6. Complex Chained Query:")
	example6()

	// Example 7: Group By
	fmt.Println("\n7. Group By Query:")
	example7()

	// Example 8: Chunk Processing
	fmt.Println("\n8. Chunk Processing:")
	example8()

	// Example 9: Transaction
	fmt.Println("\n9. Transaction Example:")
	example9()

	// Example 10: Joins
	fmt.Println("\n10. Join Example:")
	example10()
}

func example1() {
	// Find all electronics products
	// products, err := gormx.NewQuery[ChainChainProduct]().
	// 	Where("category", "Electronics").
	// 	Find()
	// if err != nil {
	// 	log.Printf("Error: %v", err)
	// 	return
	// }
	// for _, p := range products {
	// 	fmt.Printf("  - %s: $%.2f\n", p.Name, p.Price)
	// }

	fmt.Println("  products, err := gormx.NewQuery[ChainChainProduct]().")
	fmt.Println("      Where(\"category\", \"Electronics\").")
	fmt.Println("      Find()")
}

func example2() {
	// Find in-stock electronics under $1000
	// products, err := gormx.NewQuery[ChainChainProduct]().
	// 	Where("category", "Electronics").
	// 	Where("in_stock", true).
	// 	Where("price", 1000).  // Less than
	// 	Find()
	// if err != nil {
	// 	log.Printf("Error: %v", err)
	// 	return
	// }

	fmt.Println("  products, err := gormx.NewQuery[ChainProduct]().")
	fmt.Println("      Where(\"category\", \"Electronics\").")
	fmt.Println("      Where(\"in_stock\", true).")
	fmt.Println("      Where(\"price\", 1000).")
	fmt.Println("      Find()")
}

func example3() {
	// Get top 5 most expensive products
	// products, err := gormx.NewQuery[ChainProduct]().
	// 	OrderByDesc("price").
	// 	Limit(5).
	// 	Find()
	// if err != nil {
	// 	log.Printf("Error: %v", err)
	// 	return
	// }

	fmt.Println("  products, err := gormx.NewQuery[ChainProduct]().")
	fmt.Println("      OrderByDesc(\"price\").")
	fmt.Println("      Limit(5).")
	fmt.Println("      Find()")
}

func example4() {
	// Get page 2 with 10 items per page
	// products, total, err := gormx.NewQuery[ChainProduct]().
	// 	Where("in_stock", true).
	// 	OrderByAsc("name").
	// 	Paginate(2, 10)
	// if err != nil {
	// 	log.Printf("Error: %v", err)
	// 	return
	// }
	// fmt.Printf("  Page 2: %d products (Total: %d)\n", len(products), total)

	fmt.Println("  products, total, err := gormx.NewQuery[ChainProduct]().")
	fmt.Println("      Where(\"in_stock\", true).")
	fmt.Println("      OrderByAsc(\"name\").")
	fmt.Println("      Paginate(2, 10)")
}

func example5() {
	// Calculate statistics
	// totalValue, _ := gormx.NewQuery[ChainProduct]().
	// 	Where("in_stock", true).
	// 	Sum("price")
	//
	// avgPrice, _ := gormx.NewQuery[ChainProduct]().
	// 	Where("in_stock", true).
	// 	Avg("price")
	//
	// count, _ := gormx.NewQuery[ChainProduct]().
	// 	Where("in_stock", true).
	// 	Count()
	//
	// fmt.Printf("  Total Value: $%.2f\n", totalValue)
	// fmt.Printf("  Average Price: $%.2f\n", avgPrice)
	// fmt.Printf("  ChainProduct Count: %d\n", count)

	fmt.Println("  totalValue, _ := gormx.NewQuery[ChainProduct]().")
	fmt.Println("      Where(\"in_stock\", true).")
	fmt.Println("      Sum(\"price\")")
	fmt.Println()
	fmt.Println("  avgPrice, _ := gormx.NewQuery[ChainProduct]().")
	fmt.Println("      Where(\"in_stock\", true).")
	fmt.Println("      Avg(\"price\")")
}

func example6() {
	// Complex query with multiple conditions and operations
	// products, err := gormx.NewQuery[ChainProduct]().
	// 	WhereIn("category", []string{"Electronics", "Books"}).
	// 	Where("in_stock", true).
	// 	OrderByDesc("price").
	// 	OrderByAsc("name").
	// 	Limit(10).
	// 	Select("name", "price", "category").
	// 	Find()
	// if err != nil {
	// 	log.Printf("Error: %v", err)
	// 	return
	// }

	fmt.Println("  products, err := gormx.NewQuery[ChainProduct]().")
	fmt.Println("      WhereIn(\"category\", []string{\"Electronics\", \"Books\"}).")
	fmt.Println("      Where(\"in_stock\", true).")
	fmt.Println("      OrderByDesc(\"price\").")
	fmt.Println("      OrderByAsc(\"name\").")
	fmt.Println("      Limit(10).")
	fmt.Println("      Select(\"name\", \"price\", \"category\").")
	fmt.Println("      Find()")
}

func example7() {
	// Group by category with aggregations
	// var results []struct {
	// 	Category string
	// 	Count    int64
	// 	Total    float64
	// 	Avg      float64
	// }
	//
	// err := gormx.NewQuery[ChainProduct]().
	// 	Select("category", "COUNT(*) as count", "SUM(price) as total", "AVG(price) as avg").
	// 	GroupBy("category").
	// 	Scan(&results)
	// if err != nil {
	// 	log.Printf("Error: %v", err)
	// 	return
	// }
	//
	// for _, r := range results {
	// 	fmt.Printf("  %s: %d products, Total: $%.2f, Avg: $%.2f\n",
	// 		r.Category, r.Count, r.Total, r.Avg)
	// }

	fmt.Println("  var results []struct {")
	fmt.Println("      Category string")
	fmt.Println("      Count    int64")
	fmt.Println("      Total    float64")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("  err := gormx.NewQuery[ChainProduct]().")
	fmt.Println("      Select(\"category\", \"COUNT(*) as count\", \"SUM(price) as total\").")
	fmt.Println("      GroupBy(\"category\").")
	fmt.Println("      Scan(&results)")
}

func example8() {
	// Process records in chunks
	// err := gormx.NewQuery[ChainProduct]().
	// 	Where("in_stock", true).
	// 	Chunk(100, func(products []*ChainProduct) error {
	// 		// Process each chunk
	// 		for _, p := range products {
	// 			// Do something with each product
	// 			fmt.Printf("  Processing: %s\n", p.Name)
	// 		}
	// 		return nil
	// 	})
	// if err != nil {
	// 	log.Printf("Error: %v", err)
	// }

	fmt.Println("  err := gormx.NewQuery[ChainProduct]().")
	fmt.Println("      Where(\"in_stock\", true).")
	fmt.Println("      Chunk(100, func(products []*ChainProduct) error {")
	fmt.Println("          // Process each chunk")
	fmt.Println("          return nil")
	fmt.Println("      })")
}

func example9() {
	// Execute operations in a transaction
	// err := gormx.NewQuery[ChainProduct]().Transaction(func(tx *gormx.QueryBuilder[ChainProduct]) error {
	// 	// Create a new product
	// 	product := &ChainProduct{
	// 		Name:     "New ChainProduct",
	// 		Category: "Electronics",
	// 		Price:    299.99,
	// 		Quantity: 10,
	// 		InStock:  true,
	// 	}
	// 	if err := tx.Create(product); err != nil {
	// 		return err
	// 	}
	//
	// 	// Update other products
	// 	updates := map[string]interface{}{
	// 		"in_stock": false,
	// 	}
	// 	if err := tx.Where("quantity", 0).Update(updates); err != nil {
	// 		return err
	// 	}
	//
	// 	return nil
	// })
	//
	// if err != nil {
	// 	log.Printf("Transaction failed: %v", err)
	// }

	fmt.Println("  err := gormx.NewQuery[ChainProduct]().Transaction(func(tx *gormx.QueryBuilder[ChainProduct]) error {")
	fmt.Println("      // Create a new product")
	fmt.Println("      if err := tx.Create(product); err != nil {")
	fmt.Println("          return err")
	fmt.Println("      }")
	fmt.Println()
	fmt.Println("      // Update other products")
	fmt.Println("      return tx.Where(\"quantity\", 0).Update(updates)")
	fmt.Println("  })")
}

func example10() {
	// Join with another table
	// type ChainProductWithSupplier struct {
	// 	ChainProduct
	// 	SupplierName string `gorm:"column:supplier_name"`
	// }
	//
	// var results []*ChainProductWithSupplier
	// err := gormx.NewQuery[ChainProduct]().
	// 	Select("products.*, suppliers.name as supplier_name").
	// 	LeftJoin("suppliers", "suppliers.id = products.supplier_id").
	// 	Where("products.in_stock", true).
	// 	Scan(&results)
	// if err != nil {
	// 	log.Printf("Error: %v", err)
	// 	return
	// }

	fmt.Println("  err := gormx.NewQuery[ChainProduct]().")
	fmt.Println("      Select(\"products.*, suppliers.name as supplier_name\").")
	fmt.Println("      LeftJoin(\"suppliers\", \"suppliers.id = products.supplier_id\").")
	fmt.Println("      Where(\"products.in_stock\", true).")
	fmt.Println("      Scan(&results)")
}

// Real-world use case examples (see CHAIN.md for implementation details)

/*
Example implementations:

func GetProductsByCategory(category string, page, pageSize int) ([]*ChainProduct, int64, error) {
	return gormx.NewQuery[ChainProduct]().
		Where("category", category).
		Where("in_stock", true).
		OrderByDesc("created_at").
		Paginate(page, pageSize)
}

func SearchChainProducts(keyword string, minPrice, maxPrice float64) ([]*ChainProduct, error) {
	return gormx.NewQuery[ChainProduct]().
		WhereLike("name", keyword).
		Where("price", minPrice). // >= minPrice (would need custom implementation)
		Where("price", maxPrice). // <= maxPrice (would need custom implementation)
		OrderByAsc("price").
		Find()
}

func GetLowStockChainProducts(threshold int) ([]*ChainProduct, error) {
	return gormx.NewQuery[ChainProduct]().
		Where("quantity", threshold). // < threshold
		Where("in_stock", true).
		OrderByAsc("quantity").
		Find()
}

func CalculateCategoryStats(category string) (count int64, total float64, avg float64, err error) {
	query := gormx.NewQuery[ChainProduct]().Where("category", category)

	count, err = query.Count()
	if err != nil {
		return
	}

	total, err = query.Sum("price")
	if err != nil {
		return
	}

	avg, err = query.Avg("price")
	return
}
*/
