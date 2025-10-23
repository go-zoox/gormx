# GORMX Examples

This directory contains example applications demonstrating the usage of GORMX features.

## Directory Structure

```
examples/
├── aggregate/     - Aggregate query examples
│   └── main.go    - Demonstrates Sum, Avg, Min, Max, GroupBy operations
└── chain/         - Chain query builder examples
    └── main.go    - Demonstrates fluent API for building queries
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
- [FEATURES.md](../FEATURES.md) - Complete feature list
