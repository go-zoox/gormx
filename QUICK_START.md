# GORMX Quick Start Guide

A quick guide to get you started with GORMX - a powerful GORM extension with aggregate queries and chain query builder.

## Installation

```bash
go get -u github.com/go-zoox/gormx
```

## Setup Database Connection

```go
package main

import (
    "github.com/go-zoox/gormx"
)

func main() {
    // Connect to MySQL
    err := gormx.LoadDB("mysql", "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
    if err != nil {
        panic(err)
    }
    
    // Or PostgreSQL
    // err := gormx.LoadDB("postgres", "host=localhost user=user password=pass dbname=test port=5432 sslmode=disable")
    
    // Or SQLite
    // err := gormx.LoadDB("sqlite", "test.db")
}
```

## Define Your Model

```go
type Product struct {
    ID        uint      `gorm:"primarykey"`
    Name      string    `gorm:"column:name"`
    Category  string    `gorm:"column:category"`
    Price     float64   `gorm:"column:price"`
    Quantity  int       `gorm:"column:quantity"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Product) TableName() string {
    return "products"
}
```

## Basic CRUD Operations

### Create
```go
product := &Product{
    Name:     "Laptop",
    Category: "Electronics",
    Price:    999.99,
    Quantity: 10,
}
_, err := gormx.Create(product)
```

### Read
```go
// Find by ID
product, err := gormx.FindByID[Product](1)

// Find one
product, err := gormx.FindOne[Product](map[any]any{"name": "Laptop"})

// Find all with conditions
where := gormx.NewWhere()
where.Set("category", "Electronics")
products, err := gormx.FindAll[Product](where, nil)
```

### Update
```go
err := gormx.Update[Product](1, func(p *Product) {
    p.Price = 899.99
    p.Quantity = 15
})
```

### Delete
```go
err := gormx.DeleteOneByID[Product](1)
```

## Chain Query Builder

The chain query builder provides a fluent interface for building complex queries:

### Simple Query
```go
products, err := gormx.NewQuery[Product]().
    Where("category", "Electronics").
    OrderByDesc("price").
    Limit(10).
    Find()
```

### Pagination
```go
products, total, err := gormx.NewQuery[Product]().
    Where("category", "Electronics").
    OrderByAsc("name").
    Paginate(1, 20)  // Page 1, 20 items per page

fmt.Printf("Found %d products (Total: %d)\n", len(products), total)
```

### Multiple Conditions
```go
products, err := gormx.NewQuery[Product]().
    Where("category", "Electronics").
    Where("price", 1000).  // price > 1000 (needs custom implementation)
    Where("in_stock", true).
    OrderByDesc("price").
    Find()
```

### Search with LIKE
```go
products, err := gormx.NewQuery[Product]().
    WhereLike("name", "Laptop").
    Find()
```

### Using IN clause
```go
products, err := gormx.NewQuery[Product]().
    WhereIn("category", []string{"Electronics", "Books", "Toys"}).
    Find()
```

## Aggregate Queries

### Sum
```go
totalValue, err := gormx.Sum[Product]("price", nil)

// With conditions
where := gormx.NewWhere()
where.Set("category", "Electronics")
electronicsTotal, err := gormx.Sum[Product]("price", where)
```

### Average
```go
avgPrice, err := gormx.Avg[Product]("price", nil)
```

### Min/Max
```go
minPrice, err := gormx.Min[Product]("price", nil)
maxPrice, err := gormx.Max[Product]("price", nil)
```

### Count
```go
count, err := gormx.Count[Product](nil)

// Count with conditions
where := gormx.NewWhere()
where.Set("in_stock", true)
inStockCount, err := gormx.Count[Product](where)
```

### Count Distinct
```go
uniqueCategories, err := gormx.CountDistinct[Product]("category", nil)
```

### Using Chain Query for Aggregates
```go
// Sum with chain query
totalValue, err := gormx.NewQuery[Product]().
    Where("category", "Electronics").
    Sum("price")

// Average with chain query
avgPrice, err := gormx.NewQuery[Product]().
    Where("in_stock", true).
    Avg("price")
```

## Group By Queries

### Using GroupBy Function
```go
results, err := gormx.GroupBy[Product](
    []string{"category"}, 
    nil, 
    []string{"COUNT(*) as count", "SUM(price) as total", "AVG(price) as avg"},
)

for _, result := range results {
    fmt.Printf("Category: %v\n", result.Group["category"])
    fmt.Printf("  Count: %d\n", result.Count)
    fmt.Printf("  Total: %.2f\n", result.Sum)
    fmt.Printf("  Average: %.2f\n", result.Avg)
}
```

### Using Chain Query
```go
var results []struct {
    Category string
    Count    int64
    Total    float64
    Avg      float64
}

err := gormx.NewQuery[Product]().
    Select("category", "COUNT(*) as count", "SUM(price) as total", "AVG(price) as avg").
    GroupBy("category").
    Scan(&results)
```

## Transactions

```go
err := gormx.NewQuery[Product]().Transaction(func(tx *gormx.QueryBuilder[Product]) error {
    // Create a new product
    product := &Product{
        Name:     "New Product",
        Category: "Electronics",
        Price:    299.99,
    }
    if err := tx.Create(product); err != nil {
        return err
    }
    
    // Update other products
    updates := map[string]interface{}{
        "in_stock": false,
    }
    if err := tx.Where("quantity", 0).Update(updates); err != nil {
        return err
    }
    
    // Transaction will auto-commit if no error, or rollback if error is returned
    return nil
})
```

## Chunk Processing

Process large datasets in chunks to avoid memory issues:

```go
err := gormx.NewQuery[Product]().
    Where("in_stock", true).
    Chunk(100, func(products []*Product) error {
        // Process each chunk of 100 products
        for _, p := range products {
            // Do something with each product
            fmt.Printf("Processing: %s\n", p.Name)
        }
        return nil
    })
```

## Where Conditions

GORMX provides flexible where condition building:

```go
where := gormx.NewWhere()

// Equal
where.Set("category", "Electronics")

// Not Equal
where.Set("status", "deleted", &gormx.SetWhereOptions{IsNotEqual: true})

// Fuzzy Search (LIKE)
where.Set("name", "Laptop", &gormx.SetWhereOptions{IsFuzzy: true})

// IN
where.Set("category", []string{"Electronics", "Books"}, &gormx.SetWhereOptions{IsIn: true})

// NOT IN
where.Set("id", []int{1, 2, 3}, &gormx.SetWhereOptions{IsNotIn: true})

// Use the where conditions
products, err := gormx.FindAll[Product](where, nil)
```

## Ordering

```go
orderBy := &gormx.OrderBy{}
orderBy.Set("price", true)   // DESC
orderBy.Set("name", false)   // ASC

products, err := gormx.FindAll[Product](nil, orderBy)

// Or using chain query
products, err := gormx.NewQuery[Product]().
    OrderByDesc("price").
    OrderByAsc("name").
    Find()
```

## List with Pagination

```go
page := uint(1)
pageSize := uint(20)

where := gormx.NewWhere()
where.Set("category", "Electronics")

orderBy := &gormx.OrderBy{}
orderBy.Set("created_at", true)

products, total, err := gormx.List[Product](page, pageSize, where, orderBy)

fmt.Printf("Page %d: %d products (Total: %d)\n", page, len(products), total)
```

## Real-World Example

```go
package main

import (
    "fmt"
    "github.com/go-zoox/gormx"
)

func main() {
    // Initialize database
    err := gormx.LoadDB("mysql", "user:pass@tcp(localhost:3306)/shop")
    if err != nil {
        panic(err)
    }
    
    // Get electronics products under $1000, sorted by price
    products, err := gormx.NewQuery[Product]().
        Where("category", "Electronics").
        Where("price", 1000).  // Needs custom comparison
        Where("in_stock", true).
        OrderByAsc("price").
        Limit(20).
        Find()
    
    if err != nil {
        panic(err)
    }
    
    for _, p := range products {
        fmt.Printf("%s: $%.2f\n", p.Name, p.Price)
    }
    
    // Calculate category statistics
    totalValue, _ := gormx.NewQuery[Product]().
        Where("category", "Electronics").
        Sum("price")
    
    avgPrice, _ := gormx.NewQuery[Product]().
        Where("category", "Electronics").
        Avg("price")
    
    count, _ := gormx.NewQuery[Product]().
        Where("category", "Electronics").
        Count()
    
    fmt.Printf("\nElectronics Statistics:\n")
    fmt.Printf("Total Products: %d\n", count)
    fmt.Printf("Total Value: $%.2f\n", totalValue)
    fmt.Printf("Average Price: $%.2f\n", avgPrice)
}
```

## Next Steps

- See [CHAIN.md](CHAIN.md) for complete chain query builder documentation
- See [AGGREGATE.md](AGGREGATE.md) for complete aggregate query documentation
- Check out the [examples](examples/) directory for more usage examples

## Tips

1. **Use Chain Queries**: They provide better readability and type safety
2. **Use Pagination**: Don't load all records at once in production
3. **Use Transactions**: For operations that need to be atomic
4. **Use Chunk Processing**: For processing large datasets
5. **Use Indexes**: Create database indexes on frequently queried fields
