# GORMX Aggregate Query Guide

This guide explains how to use the aggregate query functions in GORMX.

## Available Aggregate Functions

### Basic Aggregates

#### Sum
Calculates product of all values in a field.

```go
total, err := gormx.Sum[Product]("price", nil)
// Returns: float64, error
```

#### Average
Calculates the average of all values in a field.

```go
avg, err := gormx.Avg[Product]("price", nil)
// Returns: float64, error
```

#### Min/Max
Finds the minimum or maximum value in a field.

```go
min, err := gormx.Min[Product]("price", nil)
max, err := gormx.Max[Product]("price", nil)
// Returns: interface{}, error
```

#### Count Distinct
Counts the number of distinct values in a field.

```go
count, err := gormx.CountDistinct[Product]("category", nil)
// Returns: int64, error
```

### Advanced Aggregates

#### Group By
Groups records by specified fields and performs aggregations.

```go
results, err := gormx.GroupBy[Product](
    []string{"category"}, 
    nil, 
    []string{"COUNT(*) as count", "SUM(price) as sum", "AVG(price) as avg"},
)
// Returns: []GroupByResult, error
```

#### Multiple Aggregations
Performs multiple aggregate operations in a single query.

```go
aggregates, err := gormx.Aggregate[Product](
    "price", 
    nil, 
    []string{"sum", "avg", "min", "max", "count"},
)
// Returns: map[string]interface{}, error
```

## Using Where Conditions

All aggregate functions support the `Where` parameter for filtering data:

```go
// Create where conditions
where := gormx.NewWhere()
where.Set("category", "Electronics")
where.Set("price", 100, &gormx.SetWhereOptions{IsFuzzy: true}) // price > 100

// Use with aggregates
total, err := gormx.Sum[Product]("price", where)
```

## Supported Where Conditions

- **Equal**: `where.Set("field", value)`
- **Not Equal**: `where.Set("field", value, &gormx.SetWhereOptions{IsNotEqual: true})`
- **Fuzzy Search**: `where.Set("field", value, &gormx.SetWhereOptions{IsFuzzy: true})`
- **IN**: `where.Set("field", values, &gormx.SetWhereOptions{IsIn: true})`
- **NOT IN**: `where.Set("field", values, &gormx.SetWhereOptions{IsNotIn: true})`

## Group By Results

The `GroupBy` function returns `GroupByResult` structs with the following fields:

```go
type GroupByResult struct {
    Group map[string]interface{} // Group by field values
    Count int64                  // Count of records in group
    Sum   float64               // Sum of numeric field
    Avg   float64               // Average of numeric field
    Min   interface{}           // Minimum value
    Max   interface{}           // Maximum value
}
```

## Supported Aggregate Operations

When using the `Aggregate` function, you can specify these operations:

- `"sum"` - Sum of values
- `"avg"` - Average of values
- `"min"` - Minimum value
- `"max"` - Maximum value
- `"count"` - Count of records
- `"count_distinct"` - Count of distinct values

## Performance Considerations

1. **Indexes**: Ensure your database has appropriate indexes on fields used in WHERE conditions and GROUP BY clauses.

2. **Large Datasets**: For large datasets, consider using pagination or limiting the scope of your queries.

3. **Multiple Aggregations**: The `Aggregate` function is more efficient than running multiple separate aggregate queries.

## Error Handling

All aggregate functions return errors that should be handled:

```go
total, err := gormx.Sum[Product]("price", nil)
if err != nil {
    log.Printf("Sum query failed: %v", err)
    return
}
```

## Examples

See the `examples/aggregate_example.go` file for complete working examples of all aggregate functions.
