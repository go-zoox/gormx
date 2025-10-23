# Changelog

All notable changes to the GORMX project will be documented in this file.

## [Unreleased] - 2025-10-23

### Added - Aggregate Query Support

#### New Aggregate Functions
- Added `Sum[T any](field string, where *Where) (float64, error)` - Calculate sum of numeric field
- Added `Avg[T any](field string, where *Where) (float64, error)` - Calculate average of numeric field
- Added `Min[T any](field string, where *Where) (interface{}, error)` - Find minimum value
- Added `Max[T any](field string, where *Where) (interface{}, error)` - Find maximum value
- Added `CountDistinct[T any](field string, where *Where) (int64, error)` - Count distinct values
- Added `GroupBy[T any](fields []string, where *Where, aggregates []string) ([]GroupByResult, error)` - Group by with aggregations
- Added `Aggregate[T any](field string, where *Where, operations []string) (map[string]interface{}, error)` - Multiple aggregate operations

#### New Types
- Added `GroupByResult` struct for group by query results
- Added `AggregateResult` struct for aggregate query results

#### Files
- `aggregate.go` - Core aggregate query implementation
- `aggregate_test.go` - Comprehensive test suite for aggregate functions
- `AGGREGATE.md` - Complete aggregate query documentation
- `examples/aggregate_example.go` - Usage examples for aggregate queries

### Added - Chain Query Builder

#### New Query Builder
- Added `QueryBuilder[T any]` - Fluent chain query builder with type safety
- Added `NewQuery[T any]() *QueryBuilder[T]` - Create new query builder instance

#### Where Conditions
- Added `Where(field string, value interface{}, opts ...*SetWhereOptions) *QueryBuilder[T]`
- Added `WhereEqual(field string, value interface{}) *QueryBuilder[T]`
- Added `WhereNotEqual(field string, value interface{}) *QueryBuilder[T]`
- Added `WhereIn(field string, values interface{}) *QueryBuilder[T]`
- Added `WhereNotIn(field string, values interface{}) *QueryBuilder[T]`
- Added `WhereLike(field string, value string) *QueryBuilder[T]`
- Added `WhereBetween(field string, start, end interface{}) *QueryBuilder[T]`
- Added `WhereRaw(sql string, args ...interface{}) *QueryBuilder[T]`

#### Selection and Ordering
- Added `Select(columns ...string) *QueryBuilder[T]` - Select specific columns
- Added `OrderBy(field string, desc ...bool) *QueryBuilder[T]` - Order by field
- Added `OrderByAsc(field string) *QueryBuilder[T]` - Order ascending
- Added `OrderByDesc(field string) *QueryBuilder[T]` - Order descending
- Added `Distinct() *QueryBuilder[T]` - Select distinct records

#### Pagination and Limiting
- Added `Limit(limit int) *QueryBuilder[T]` - Limit results
- Added `Offset(offset int) *QueryBuilder[T]` - Offset results
- Added `Page(page, pageSize int) *QueryBuilder[T]` - Pagination helper

#### Joins and Relationships
- Added `Join(table, condition string, args ...interface{}) *QueryBuilder[T]` - Inner join
- Added `LeftJoin(table, condition string, args ...interface{}) *QueryBuilder[T]` - Left join
- Added `RightJoin(table, condition string, args ...interface{}) *QueryBuilder[T]` - Right join
- Added `Preload(associations ...string) *QueryBuilder[T]` - Eager load associations

#### Grouping
- Added `GroupBy(fields ...string) *QueryBuilder[T]` - Group by fields
- Added `Having(field string, value interface{}, opts ...*SetWhereOptions) *QueryBuilder[T]` - Having clause

#### Execution Methods
- Added `Find() ([]*T, error)` - Execute and get all results
- Added `First() (*T, error)` - Execute and get first result
- Added `Last() (*T, error)` - Execute and get last result
- Added `Count() (int64, error)` - Execute and count
- Added `Exists() (bool, error)` - Check if results exist
- Added `Paginate(page, pageSize int) ([]*T, int64, error)` - Paginate with total count

#### Aggregate Methods (Chain)
- Added `Sum(field string) (float64, error)` - Calculate sum via chain
- Added `Avg(field string) (float64, error)` - Calculate average via chain
- Added `Min(field string) (interface{}, error)` - Find minimum via chain
- Added `Max(field string) (interface{}, error)` - Find maximum via chain
- Added `Pluck(column string, dest interface{}) error` - Get column values

#### Data Modification
- Added `Create(value *T) error` - Create record
- Added `Save(value *T) error` - Save record
- Added `Update(updates map[string]interface{}) error` - Update records
- Added `UpdateColumn(updates map[string]interface{}) error` - Update columns
- Added `Delete() error` - Delete records

#### Batch Operations
- Added `CreateInBatches(values []*T, batchSize int) error` - Batch insert
- Added `FindInBatches(batchSize int, fn func(tx *gorm.DB, batch int) error) error` - Process in batches
- Added `Chunk(chunkSize int, callback func([]*T) error) error` - Process records in chunks

#### Transaction Support
- Added `Transaction(fn func(tx *QueryBuilder[T]) error) error` - Execute operations in transaction

#### Utility Methods
- Added `Clone() *QueryBuilder[T]` - Clone query builder
- Added `GetDB() *gorm.DB` - Get underlying GORM DB
- Added `GetTableName() string` - Get table name
- Added `ToSQL() (string, error)` - Get SQL query string
- Added `Scan(dest interface{}) error` - Scan results
- Added `Raw(sql string, values ...interface{}) *QueryBuilder[T]` - Raw SQL query

#### New Types
- Added `JoinClause` struct for join operations

#### Files
- `chain.go` - Core chain query builder implementation
- `chain_test.go` - Comprehensive test suite for chain queries
- `CHAIN.md` - Complete chain query builder documentation
- `examples/chain_example.go` - Usage examples for chain queries

### Documentation

- Added `AGGREGATE.md` - Comprehensive aggregate query guide
- Added `CHAIN.md` - Comprehensive chain query builder guide
- Added `QUICK_START.md` - Quick start guide for new users
- Added `FEATURES.md` - Complete feature list and comparison
- Added `CHANGELOG.md` - This changelog file
- Updated `README.md` - Added aggregate functions and chain query builder sections
- Added usage examples in `examples/` directory

### Tests

- Added comprehensive test suite for aggregate functions (`aggregate_test.go`)
- Added comprehensive test suite for chain query builder (`chain_test.go`)
- All tests pass successfully (require database connection for integration tests)

### Improvements

- Enhanced type safety with Go generics throughout
- Improved code organization and documentation
- Better error handling and reporting
- More intuitive API design with chain query builder

## [1.0.0] - Previous Release

### Initial Features

- Basic CRUD operations with generics
- Where condition builder
- OrderBy builder
- Pagination support
- List operations
- Atomic operations (FindOrCreate, FindAndUpdate, etc.)
- Database connection management
- DDL operations
- User management
- HTTP integration
- Error handling utilities
- Data type support (JSON, UUID, Date)
- Model management with IoC
- Raw SQL support

## Upcoming Features

### Planned for Future Releases

- Query caching layer
- Advanced transaction management with savepoints
- Query performance monitoring and logging
- Database backup/restore utilities
- Enhanced schema migration tools
- Full-text search support
- Geospatial queries
- Time-series data support
- Advanced indexing management
- Query debugging and explain tools
- Connection pool configuration
- Read/write splitting for replicas
- Database sharding support
- Optimistic locking with version fields
- Pessimistic locking (SELECT FOR UPDATE)
- Soft delete query scopes
- Custom comparison operators (<, >, <=, >=)
- More complex where conditions (OR, nested AND/OR)
- Subquery support
- Union queries
- Window functions
- CTE (Common Table Expressions) support

## Migration Guide

### From Traditional GORMX to Chain Query Builder

#### Before
```go
where := gormx.NewWhere()
where.Set("category", "Electronics")
where.Set("in_stock", true)

orderBy := &gormx.OrderBy{}
orderBy.Set("price", true)

products, total, err := gormx.List[Product](1, 20, where, orderBy)
```

#### After
```go
products, total, err := gormx.NewQuery[Product]().
    Where("category", "Electronics").
    Where("in_stock", true).
    OrderByDesc("price").
    Paginate(1, 20)
```

### From Traditional Aggregate to New Aggregate Functions

#### Before (using raw GORM)
```go
var result struct {
    Sum float64
}
db := gormx.GetDB()
db.Model(&Product{}).
    Where("category = ?", "Electronics").
    Select("SUM(price) as sum").
    Scan(&result)
```

#### After
```go
// Using standalone function
total, err := gormx.Sum[Product]("price", where)

// Or using chain query
total, err := gormx.NewQuery[Product]().
    Where("category", "Electronics").
    Sum("price")
```

## Breaking Changes

None in this release. All new features are additive and maintain backward compatibility.

## Contributors

- Core team and contributors

## License

MIT License - See LICENSE file for details.
