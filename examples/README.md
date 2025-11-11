# GORMX Examples

This directory contains example applications demonstrating the usage of GORMX features.

## Directory Structure

```
examples/
├── aggregate/          - Aggregate query examples
│   └── main.go         - Demonstrates Sum, Avg, Min, Max, GroupBy operations
├── chain/              - Chain query builder examples
│   └── main.go         - Demonstrates fluent API for building queries
├── where_generic/      - Where generic type examples
│   └── main.go         - Demonstrates using map and *Where types
├── where_generic_all/  - Advanced where generic examples
│   └── main.go         - Comprehensive where condition examples
├── where_range/        - Range query examples
│   └── main.go         - Demonstrates BETWEEN queries for dates, numbers, etc.
├── where_range_modes/  - Range mode examples
│   └── main.go         - Demonstrates [A,B], (A,B), [A,B), (A,B] modes
└── where_comparison/   - Comparison operator examples
    └── main.go         - Demonstrates >, <, >=, <= operators
```

## Running Examples

### Aggregate Query Example

```bash
cd examples/aggregate
go run main.go
```

This example demonstrates:
- Basic aggregate functions (Sum, Avg, Min, Max)
- Count and CountDistinct
- GroupBy with aggregations
- Multiple aggregations in one query
- Using aggregate functions with WHERE conditions

### Chain Query Builder Example

```bash
cd examples/chain
go run main.go
```

This example demonstrates:
- Simple queries with Where conditions
- Multiple where conditions
- Ordering and limiting results
- Pagination
- Aggregate queries via chain
- Complex chained queries
- GroupBy with chain query
- Chunk processing
- Transactions
- Join operations

### Where Generic Type Example

```bash
cd examples/where_generic
go run main.go
```

This example demonstrates:
- Using `map[any]any` for simple queries
- Using `*Where` for simple queries
- Using `*Where` with complex conditions (fuzzy search, not equal)
- Using `*Where` with IN queries

### Where Range Query Example

```bash
cd examples/where_range
go run main.go
```

This example demonstrates:
- Age range queries with integers
- Date range queries with strings
- Date range queries with time.Time
- Multiple conditions with ranges
- Combining range with other conditions (fuzzy search, etc.)
- Price range queries with floats
- URL query parameter syntax for range queries

### Where Range Modes Example

```bash
cd examples/where_range_modes
go run main.go
```

This example demonstrates:
- Closed range [A, B] - both inclusive (default)
- Open range (A, B) - both exclusive
- Left closed range [A, B) - left inclusive, right exclusive
- Right closed range (A, B] - left exclusive, right inclusive
- Comparison of all four modes
- Real-world use cases for each mode
- Multiple conditions with different range modes

### Where Comparison Operator Example

```bash
cd examples/where_comparison
go run main.go
```

This example demonstrates:
- Greater than (>) queries
- Less than (<) queries
- Greater or equal (>=) queries
- Less or equal (<=) queries
- Date comparisons
- Multiple comparison conditions
- Mixing comparison operators with other conditions
- Real-world use cases
- URL query parameter syntax for comparison operators

## Note

These examples show the API usage through commented code and print statements. To actually run the queries, you need to:

1. Set up a database connection
2. Create the appropriate tables
3. Uncomment the actual query code
4. Run the examples

For example:

```go
// Uncomment this section in the example files
err := gormx.LoadDB("mysql", "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
if err != nil {
    panic(err)
}

// Migrate tables
err = gormx.GetDB().AutoMigrate(&Product{})
if err != nil {
    panic(err)
}
```

## Documentation

For more detailed information, see:
- [QUICK_START.md](../QUICK_START.md) - Quick start guide
- [CHAIN.md](../CHAIN.md) - Chain query builder documentation
- [AGGREGATE.md](../AGGREGATE.md) - Aggregate query documentation
- [WHERE_GENERIC.md](../WHERE_GENERIC.md) - Where generic type documentation
- [FEATURES.md](../FEATURES.md) - Complete feature list
