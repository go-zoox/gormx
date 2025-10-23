# GORMX Chain Query Guide

This guide explains how to use the chain query builder in GORMX for building complex database queries with a fluent interface.

## Overview

The chain query builder provides a fluent, intuitive API for constructing database queries. It allows you to build complex queries step by step while maintaining readability.

## Basic Usage

### Creating a Query Builder

```go
query := gormx.NewQuery[Product]()
```

## Query Methods

### Where Conditions

#### Simple Where
```go
products, err := gormx.NewQuery[Product]().
    Where("category", "Electronics").
    Find()
```

#### WhereEqual
```go
products, err := gormx.NewQuery[Product]().
    WhereEqual("status", "active").
    Find()
```

#### WhereNotEqual
```go
products, err := gormx.NewQuery[Product]().
    WhereNotEqual("status", "deleted").
    Find()
```

#### WhereIn
```go
products, err := gormx.NewQuery[Product]().
    WhereIn("category", []string{"Electronics", "Books"}).
    Find()
```

#### WhereNotIn
```go
products, err := gormx.NewQuery[Product]().
    WhereNotIn("id", []int{1, 2, 3}).
    Find()
```

#### WhereLike (Fuzzy Search)
```go
products, err := gormx.NewQuery[Product]().
    WhereLike("name", "Laptop").
    Find()
```

### Ordering

```go
// Ascending order
products, err := gormx.NewQuery[Product]().
    OrderByAsc("price").
    Find()

// Descending order
products, err := gormx.NewQuery[Product]().
    OrderByDesc("created_at").
    Find()

// Multiple orders
products, err := gormx.NewQuery[Product]().
    OrderByDesc("price").
    OrderByAsc("name").
    Find()
```

### Limiting Results

```go
// Limit
products, err := gormx.NewQuery[Product]().
    Limit(10).
    Find()

// Limit with Offset
products, err := gormx.NewQuery[Product]().
    Limit(10).
    Offset(20).
    Find()

// Pagination helper
products, err := gormx.NewQuery[Product]().
    Page(2, 10).  // Page 2, 10 items per page
    Find()
```

### Selecting Columns

```go
products, err := gormx.NewQuery[Product]().
    Select("name", "price", "category").
    Find()
```

### Joins

```go
// Inner Join
results, err := gormx.NewQuery[Product]().
    Join("suppliers", "suppliers.id = products.supplier_id").
    Find()

// Left Join
results, err := gormx.NewQuery[Product]().
    LeftJoin("categories", "categories.id = products.category_id").
    Find()

// Right Join
results, err := gormx.NewQuery[Product]().
    RightJoin("orders", "orders.product_id = products.id").
    Find()
```

### Preload (Eager Loading)

```go
products, err := gormx.NewQuery[Product]().
    Preload("Category").
    Preload("Supplier").
    Find()
```

### Group By and Having

```go
var results []struct {
    Category string
    Count    int64
    Total    float64
}

err := gormx.NewQuery[Product]().
    Select("category", "COUNT(*) as count", "SUM(price) as total").
    GroupBy("category").
    Having("count", 10).  // count > 10
    Scan(&results)
```

### Distinct

```go
var categories []string
err := gormx.NewQuery[Product]().
    Distinct().
    Pluck("category", &categories)
```

## Execution Methods

### Find (Get All)
```go
products, err := gormx.NewQuery[Product]().
    Where("in_stock", true).
    Find()
```

### First (Get First Record)
```go
product, err := gormx.NewQuery[Product]().
    Where("id", 1).
    First()
```

### Last (Get Last Record)
```go
product, err := gormx.NewQuery[Product]().
    OrderByAsc("created_at").
    Last()
```

### Count
```go
count, err := gormx.NewQuery[Product]().
    Where("category", "Electronics").
    Count()
```

### Exists
```go
exists, err := gormx.NewQuery[Product]().
    Where("sku", "ABC123").
    Exists()
```

### Paginate (with Total Count)
```go
products, total, err := gormx.NewQuery[Product]().
    Where("in_stock", true).
    OrderByDesc("created_at").
    Paginate(1, 20)  // Page 1, 20 items per page

fmt.Printf("Found %d products (Total: %d)\n", len(products), total)
```

## Aggregate Methods

### Sum
```go
totalValue, err := gormx.NewQuery[Product]().
    Where("category", "Electronics").
    Sum("price")
```

### Avg
```go
avgPrice, err := gormx.NewQuery[Product]().
    Where("in_stock", true).
    Avg("price")
```

### Min
```go
minPrice, err := gormx.NewQuery[Product]().
    Min("price")
```

### Max
```go
maxPrice, err := gormx.NewQuery[Product]().
    Max("price")
```

### Pluck (Get Column Values)
```go
var names []string
err := gormx.NewQuery[Product]().
    Where("category", "Books").
    Pluck("name", &names)
```

## Data Modification

### Create
```go
product := &Product{
    Name:     "New Product",
    Category: "Electronics",
    Price:    299.99,
}

err := gormx.NewQuery[Product]().Create(product)
```

### Update
```go
err := gormx.NewQuery[Product]().
    Where("category", "Electronics").
    Update(map[string]interface{}{
        "in_stock": false,
    })
```

### Delete
```go
err := gormx.NewQuery[Product]().
    Where("quantity", 0).
    Delete()
```

### Save
```go
product.Price = 349.99
err := gormx.NewQuery[Product]().Save(&product)
```

## Batch Operations

### CreateInBatches
```go
products := []*Product{
    {Name: "Product 1", Price: 100},
    {Name: "Product 2", Price: 200},
    // ... more products
}

err := gormx.NewQuery[Product]().CreateInBatches(products, 100)
```

### Chunk Processing
```go
err := gormx.NewQuery[Product]().
    Where("in_stock", true).
    Chunk(100, func(products []*Product) error {
        // Process each chunk of 100 products
        for _, p := range products {
            // Do something with each product
        }
        return nil
    })
```

## Transactions

```go
err := gormx.NewQuery[Product]().Transaction(func(tx *gormx.QueryBuilder[Product]) error {
    // Create a new product
    product := &Product{Name: "New Product", Price: 299.99}
    if err := tx.Create(product); err != nil {
        return err
    }

    // Update other products
    updates := map[string]interface{}{"in_stock": false}
    if err := tx.Where("quantity", 0).Update(updates); err != nil {
        return err
    }

    // If any error occurs, the transaction will be rolled back
    return nil
})
```

## Advanced Usage

### Complex Chained Query
```go
products, err := gormx.NewQuery[Product]().
    WhereIn("category", []string{"Electronics", "Books"}).
    Where("in_stock", true).
    WhereLike("name", "Pro").
    OrderByDesc("price").
    OrderByAsc("name").
    Limit(10).
    Select("name", "price", "category").
    Find()
```

### Clone Query Builder
```go
baseQuery := gormx.NewQuery[Product]().
    Where("category", "Electronics").
    Where("in_stock", true)

// Clone for different operations
expensiveQuery := baseQuery.Clone().Where("price", 1000)
cheapQuery := baseQuery.Clone().Where("price", 100)

expensive, _ := expensiveQuery.Find()
cheap, _ := cheapQuery.Find()
```

### Get Underlying GORM DB
```go
db := gormx.NewQuery[Product]().
    Where("category", "Electronics").
    GetDB()

// Use GORM methods directly
db.Scan(&customResult)
```

## Real-World Examples

### Product Search
```go
func SearchProducts(keyword string, category string, minPrice, maxPrice float64, page, pageSize int) ([]*Product, int64, error) {
    query := gormx.NewQuery[Product]()
    
    if keyword != "" {
        query.WhereLike("name", keyword)
    }
    
    if category != "" {
        query.Where("category", category)
    }
    
    if minPrice > 0 {
        // Note: Would need custom implementation for >= comparison
        query.Where("price", minPrice)
    }
    
    if maxPrice > 0 {
        // Note: Would need custom implementation for <= comparison
        query.Where("price", maxPrice)
    }
    
    return query.
        Where("in_stock", true).
        OrderByDesc("created_at").
        Paginate(page, pageSize)
}
```

### Category Statistics
```go
func GetCategoryStats(category string) (count int64, total, avg float64, err error) {
    query := gormx.NewQuery[Product]().Where("category", category)
    
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
```

### Bulk Update Low Stock Products
```go
func UpdateLowStockStatus(threshold int) error {
    return gormx.NewQuery[Product]().
        Where("quantity", threshold).  // < threshold
        Update(map[string]interface{}{
            "status":   "low_stock",
            "alert_sent": true,
        })
}
```

## Performance Tips

1. **Use Select**: Only select the columns you need
2. **Use Indexes**: Ensure database indexes on fields used in WHERE clauses
3. **Use Pagination**: Don't load all records at once
4. **Use Chunk**: Process large datasets in chunks
5. **Use Transactions**: Group related operations for consistency and performance
6. **Clone Wisely**: Clone queries when you need to reuse a base query

## Comparison with Traditional GORM

### Before (Traditional GORM)
```go
var products []*Product
db := gormx.GetDB()
db = db.Where("category = ?", "Electronics")
db = db.Where("in_stock = ?", true)
db = db.Order("price DESC")
db = db.Limit(10)
err := db.Find(&products).Error
```

### After (Chain Query)
```go
products, err := gormx.NewQuery[Product]().
    Where("category", "Electronics").
    Where("in_stock", true).
    OrderByDesc("price").
    Limit(10).
    Find()
```

The chain query builder provides a cleaner, more intuitive API while maintaining type safety through Go generics.
