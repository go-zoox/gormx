# GORMX Features

A comprehensive list of features provided by GORMX - a powerful GORM extension library.

## Core Features

### 1. Basic CRUD Operations
- ✅ `Create[T any](one *T) (*T, error)` - Create a single record
- ✅ `Save[T any](one *T) error` - Save a record (insert or update)
- ✅ `Update[T any](id uint, uc func(*T)) error` - Update by ID with callback
- ✅ `Delete[T any](where map[any]any) error` - Delete with conditions
- ✅ `DeleteOneByID[T any](id uint) error` - Delete by ID
- ✅ `FindByID[T any](id uint) (*T, error)` - Find by ID
- ✅ `FindOne[T any](where map[any]any) (*T, error)` - Find one record
- ✅ `FindAll[T any](where *Where, orderBy *OrderBy) ([]*T, error)` - Find all records
- ✅ `Find[T any](page, pageSize uint, where *Where, orderBy *OrderBy) ([]*T, int64, error)` - Find with pagination
- ✅ `Retrieve[T any](id uint) (*T, error)` - Retrieve by ID (alias for FindByID)

### 2. Advanced Query Operations
- ✅ `List[T any](page, pageSize uint, where *Where, orderBy *OrderBy) ([]*T, int64, error)` - List with pagination
- ✅ `ListALL[T any](where *Where, orderBy *OrderBy) ([]*T, error)` - List all records
- ✅ `GetMany[T any](ids []uint) ([]*T, error)` - Get multiple records by IDs
- ✅ `Exists[T any](where map[any]any) (bool, error)` - Check if record exists
- ✅ `Has[T any](where map[string]any) bool` - Check if record exists (simplified)
- ✅ `Count[T any](where *Where) (int64, error)` - Count records
- ✅ `CountALL[T any]() (int64, error)` - Count all records

### 3. Atomic Operations
- ✅ `FindOneOrCreate[T any](where map[any]any, callback func(*T)) (*T, error)` - Find or create
- ✅ `FindOneByIDOrCreate[T any](id uint, callback func(*T)) (*T, error)` - Find by ID or create
- ✅ `FindOneAndUpdate[T any](where map[any]any, callback func(*T)) (*T, error)` - Find and update
- ✅ `FindOneByIDAndUpdate[T any](id uint, callback func(*T)) (*T, error)` - Find by ID and update
- ✅ `FindOneAndDelete[T any](where map[any]any) (*T, error)` - Find and delete
- ✅ `FindOneByIDAndDelete[T any](id uint) (*T, error)` - Find by ID and delete
- ✅ `GetOrCreate[T any](where map[any]any, callback func(*T)) (*T, error)` - Get or create (alias)

### 4. Aggregate Functions (NEW!)
- ✅ `Sum[T any](field string, where *Where) (float64, error)` - Calculate sum
- ✅ `Avg[T any](field string, where *Where) (float64, error)` - Calculate average
- ✅ `Min[T any](field string, where *Where) (interface{}, error)` - Find minimum value
- ✅ `Max[T any](field string, where *Where) (interface{}, error)` - Find maximum value
- ✅ `CountDistinct[T any](field string, where *Where) (int64, error)` - Count distinct values
- ✅ `GroupBy[T any](fields []string, where *Where, aggregates []string) ([]GroupByResult, error)` - Group by with aggregates
- ✅ `Aggregate[T any](field string, where *Where, operations []string) (map[string]interface{}, error)` - Multiple aggregations

### 5. Chain Query Builder (NEW!)

#### Query Building Methods
- ✅ `NewQuery[T any]() *QueryBuilder[T]` - Create new query builder
- ✅ `Where(field string, value interface{}, opts ...*SetWhereOptions) *QueryBuilder[T]` - Add where condition
- ✅ `WhereEqual(field string, value interface{}) *QueryBuilder[T]` - Equal condition
- ✅ `WhereNotEqual(field string, value interface{}) *QueryBuilder[T]` - Not equal condition
- ✅ `WhereIn(field string, values interface{}) *QueryBuilder[T]` - IN condition
- ✅ `WhereNotIn(field string, values interface{}) *QueryBuilder[T]` - NOT IN condition
- ✅ `WhereLike(field string, value string) *QueryBuilder[T]` - LIKE condition
- ✅ `WhereBetween(field string, start, end interface{}) *QueryBuilder[T]` - BETWEEN condition
- ✅ `WhereRaw(sql string, args ...interface{}) *QueryBuilder[T]` - Raw SQL where

#### Selection and Ordering
- ✅ `Select(columns ...string) *QueryBuilder[T]` - Select specific columns
- ✅ `OrderBy(field string, desc ...bool) *QueryBuilder[T]` - Order by field
- ✅ `OrderByAsc(field string) *QueryBuilder[T]` - Order ascending
- ✅ `OrderByDesc(field string) *QueryBuilder[T]` - Order descending
- ✅ `Distinct() *QueryBuilder[T]` - Select distinct

#### Pagination and Limiting
- ✅ `Limit(limit int) *QueryBuilder[T]` - Limit results
- ✅ `Offset(offset int) *QueryBuilder[T]` - Offset results
- ✅ `Page(page, pageSize int) *QueryBuilder[T]` - Pagination helper

#### Joins and Relationships
- ✅ `Join(table, condition string, args ...interface{}) *QueryBuilder[T]` - Inner join
- ✅ `LeftJoin(table, condition string, args ...interface{}) *QueryBuilder[T]` - Left join
- ✅ `RightJoin(table, condition string, args ...interface{}) *QueryBuilder[T]` - Right join
- ✅ `Preload(associations ...string) *QueryBuilder[T]` - Eager load associations

#### Grouping
- ✅ `GroupBy(fields ...string) *QueryBuilder[T]` - Group by fields
- ✅ `Having(field string, value interface{}, opts ...*SetWhereOptions) *QueryBuilder[T]` - Having clause

#### Execution Methods
- ✅ `Find() ([]*T, error)` - Execute and get all results
- ✅ `First() (*T, error)` - Execute and get first result
- ✅ `Last() (*T, error)` - Execute and get last result
- ✅ `Count() (int64, error)` - Execute and count
- ✅ `Exists() (bool, error)` - Check if results exist
- ✅ `Paginate(page, pageSize int) ([]*T, int64, error)` - Paginate with total count

#### Aggregate Methods
- ✅ `Sum(field string) (float64, error)` - Calculate sum
- ✅ `Avg(field string) (float64, error)` - Calculate average
- ✅ `Min(field string) (interface{}, error)` - Find minimum
- ✅ `Max(field string) (interface{}, error)` - Find maximum
- ✅ `Pluck(column string, dest interface{}) error` - Get column values

#### Data Modification
- ✅ `Create(value *T) error` - Create record
- ✅ `Save(value *T) error` - Save record
- ✅ `Update(updates map[string]interface{}) error` - Update records
- ✅ `UpdateColumn(updates map[string]interface{}) error` - Update columns
- ✅ `Delete() error` - Delete records

#### Batch Operations
- ✅ `CreateInBatches(values []*T, batchSize int) error` - Batch insert
- ✅ `FindInBatches(batchSize int, fn func(tx *gorm.DB, batch int) error) error` - Process in batches
- ✅ `Chunk(chunkSize int, callback func([]*T) error) error` - Process chunks

#### Transaction Support
- ✅ `Transaction(fn func(tx *QueryBuilder[T]) error) error` - Execute in transaction

#### Utility Methods
- ✅ `Clone() *QueryBuilder[T]` - Clone query builder
- ✅ `GetDB() *gorm.DB` - Get underlying GORM DB
- ✅ `GetTableName() string` - Get table name
- ✅ `ToSQL() (string, error)` - Get SQL query string
- ✅ `Scan(dest interface{}) error` - Scan results
- ✅ `Raw(sql string, values ...interface{}) *QueryBuilder[T]` - Raw SQL query

### 6. Where Condition Builder
- ✅ `NewWhere() *Where` - Create new where builder
- ✅ `Set(key string, value interface{}, opts ...*SetWhereOptions)` - Set condition
- ✅ `Add(key string, value interface{}, opts ...*SetWhereOptions)` - Add condition
- ✅ `Get(key string) (interface{}, bool)` - Get condition
- ✅ `Del(key string)` - Delete condition
- ✅ `Build() (string, []interface{}, error)` - Build where clause

#### Where Options
- ✅ Equal (`IsEqual`)
- ✅ Not Equal (`IsNotEqual`)
- ✅ Fuzzy Search / LIKE (`IsFuzzy`)
- ✅ IN clause (`IsIn`)
- ✅ NOT IN clause (`IsNotIn`)
- ✅ Plain SQL (`IsPlain`)
- ✅ Full Text Search (`IsFullTextSearch`)

### 7. OrderBy Builder
- ✅ `Set(key string, isDESC bool)` - Set order
- ✅ `Get(key string) (bool, bool)` - Get order
- ✅ `Del(key string)` - Delete order
- ✅ Multiple field ordering support

### 8. Database Management
- ✅ `LoadDB(engine string, dsn string) error` - Load database connection
- ✅ `Connect(engine string, dsn string, opts ...func(*LoadDBOptions)) (*gorm.DB, error)` - Connect to database
- ✅ `GetDB() *gorm.DB` - Get database instance
- ✅ `SetDB(d *gorm.DB)` - Set database instance
- ✅ `GetEngine() string` - Get database engine
- ✅ `GetDSN() string` - Get database DSN
- ✅ `Migrate()` - Auto migrate models

#### Supported Databases
- ✅ MySQL
- ✅ PostgreSQL
- ✅ SQLite
- ✅ MongoDB (partial support via DDL)

### 9. DDL Operations
- ✅ `CreateDatabase(engine, dsn, name string) error` - Create database
- ✅ `DeleteDatabase(engine, dsn, name string) error` - Delete database
- ✅ `CreateTableInDatabase(database, name string) error` - Create table
- ✅ `DeleteTableInDatabase(database, name string) error` - Delete table
- ✅ `RenameTableInDatabase(database, oldName, newName string) error` - Rename table

#### User Management
- ✅ `CreateUser(engine, dsn, username, password string) error` - Create user
- ✅ `DeleteUser(engine, dsn, username string) error` - Delete user
- ✅ `UpdateUserPassword(username, password string) error` - Update password
- ✅ `GrantUserPrivileges(username string) error` - Grant privileges
- ✅ `RevokeUserPrivileges(username string) error` - Revoke privileges
- ✅ `GrantUserPrivilegesToDatabase(username, database string) error` - Grant database privileges
- ✅ `RevokeUserPrivilegesFromDatabase(username, database string) error` - Revoke database privileges
- ✅ `ReadOnlyUserToDatabase(username, database string) error` - Grant read-only access

### 10. Data Types
- ✅ `JSON` - JSON field type
- ✅ `JSONObject` - JSON object type
- ✅ `JSONArray[T any]` - JSON array type
- ✅ `UUID` - UUID type
- ✅ `Date` - Date type

### 11. Error Handling
- ✅ All GORM error constants exposed
- ✅ Error checking helper functions:
  - `IsRecordNotFoundError(err error) bool`
  - `IsInvalidTransactionError(err error) bool`
  - `IsDuplicatedKeyError(err error) bool`
  - `IsForeignKeyViolatedError(err error) bool`
  - And 12 more error checking functions

### 12. Generic Support
- ✅ Full Go generics support for type safety
- ✅ `Model` interface for model management
- ✅ `ModelImpl` base struct with common fields:
  - ID (uint, primary key)
  - CreatedAt (time.Time)
  - UpdatedAt (time.Time)
  - DeletedAt (soft delete)
  - Creator (uint)
  - Modifier (uint)
- ✅ `ModelGeneric[T any]` - Generic model wrapper

### 13. HTTP Integration
- ✅ `Params` - HTTP parameter parser
- ✅ `Page` - Pagination parameters
- ✅ `GetList()` - Parse list parameters from HTTP request
- ✅ Automatic query parameter parsing:
  - page, pageSize
  - where conditions
  - orderBy
  - Fuzzy search (`:*`)
  - Not equal (`:!`)
  - IN clause (`:in`)
  - NOT IN clause (`:!in`)

### 14. IoC Container
- ✅ Model registration and management
- ✅ Automatic model discovery

### 15. Raw SQL Support
- ✅ `SQL[T any](sql string, values ...any) (*T, error)` - Execute raw SQL

## Feature Comparison

| Feature | Traditional GORM | GORMX | GORMX Chain |
|---------|-----------------|-------|-------------|
| Basic CRUD | ✅ | ✅ | ✅ |
| Pagination | ✅ | ✅ | ✅ |
| Aggregates | ✅ | ✅ | ✅ |
| Group By | ✅ | ✅ | ✅ |
| Transactions | ✅ | ⚠️ | ✅ |
| Fluent API | ❌ | ⚠️ | ✅ |
| Type Safety | ⚠️ | ✅ | ✅ |
| HTTP Integration | ❌ | ✅ | ❌ |
| Atomic Operations | ❌ | ✅ | ⚠️ |
| Chunk Processing | ⚠️ | ❌ | ✅ |
| DDL Operations | ❌ | ✅ | ❌ |

Legend:
- ✅ Fully supported
- ⚠️ Partially supported or needs improvement
- ❌ Not supported

## Performance Features

1. **Batch Operations**: Insert/update/delete multiple records efficiently
2. **Chunk Processing**: Process large datasets without memory issues
3. **Query Optimization**: Automatic query building and optimization
4. **Connection Pooling**: Built on GORM's connection pooling
5. **Lazy Loading**: Load data only when needed
6. **Eager Loading**: Preload associations to avoid N+1 queries

## Developer Experience

1. **Type Safety**: Full Go generics support
2. **Fluent API**: Chain query builder for readable code
3. **Auto-completion**: IDE support for all methods
4. **Error Handling**: Comprehensive error checking
5. **Documentation**: Complete API documentation and guides
6. **Examples**: Real-world usage examples

## Testing

- ✅ Unit tests for core functions
- ✅ Unit tests for aggregate functions
- ✅ Unit tests for chain query builder
- ✅ Test helpers and fixtures

## Documentation

- ✅ README.md - Overview and basic usage
- ✅ QUICK_START.md - Quick start guide
- ✅ CHAIN.md - Chain query builder documentation
- ✅ AGGREGATE.md - Aggregate query documentation
- ✅ FEATURES.md - Feature list (this file)
- ✅ Code examples in `examples/` directory

## Future Enhancements

Potential features for future versions:

- [ ] Query caching layer
- [ ] Advanced transaction management
- [ ] Query performance monitoring
- [ ] Database backup/restore utilities
- [ ] Schema migration tools
- [ ] Full-text search support
- [ ] Geospatial queries
- [ ] Time-series data support
- [ ] Advanced indexing management
- [ ] Query debugging tools
- [ ] Connection pool management
- [ ] Read/write splitting
- [ ] Sharding support
- [ ] Optimistic locking
- [ ] Pessimistic locking

## License

MIT License - See LICENSE file for details
