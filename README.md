# GORMX - GORM Utils

## Functions

```go
package gormx // import "github.com/go-zoox/gormx"

var Version = "1.0.0"
func Create[T any](one *T) (*T, error)
func Delete[T any](where map[any]any) (err error)
func DeleteOneByID[T any](id uint) (err error)
func Exists[T any](where map[any]any) (bool, error)
func Find[T any](page, pageSize uint, where *Where, orderBy *OrderBy) (data []*T, total int64, err error)
func FindAll[T any](where *Where, orderBy *OrderBy) (data []*T, err error)
func FindByID[T any](id uint) (*T, error)
func FindOne[T any](where map[any]any) (*T, error)
func FindOneAndDelete[T any](where map[any]any) (*T, error)
func FindOneAndUpdate[T any](where map[any]any, callback func(*T)) (*T, error)
func FindOneByIDAndDelete[T any](id uint) (*T, error)
func FindOneByIDAndUpdate[T any](id uint, callback func(*T)) (*T, error)
func FindOneByIDOrCreate[T any](id uint, callback func(*T)) (*T, error)
func FindOneOrCreate[T any](where map[any]any, callback func(*T)) (*T, error)
func GetDB() *gorm.DB
func GetMany[T any](ids []uint) (data []*T, err error)
func GetOrCreate[T any](where map[any]any, callback func(*T)) (*T, error)
func Has[T any](where map[string]any) bool
func List[T any](page, pageSize uint, where *Where, orderBy *OrderBy) (data []*T, total int64, err error)
func ListALL[T any](where *Where, orderBy *OrderBy) (data []*T, err error)
func LoadDB(engine string, dsn string) (err error)
func Retrieve[T any](id uint) (*T, error)
func Save[T any](one *T) error
func Update[T any](id uint, uc func(*T)) (err error)

// Aggregate Functions
func Sum[T any](field string, where *Where) (float64, error)
func Avg[T any](field string, where *Where) (float64, error)
func Min[T any](field string, where *Where) (interface{}, error)
func Max[T any](field string, where *Where) (interface{}, error)
func CountDistinct[T any](field string, where *Where) (int64, error)
func GroupBy[T any](fields []string, where *Where, aggregates []string) ([]GroupByResult, error)
func Aggregate[T any](field string, where *Where, operations []string) (map[string]interface{}, error)

// Types
type OrderBy []OrderByOne
type OrderByOne struct{ ... }
type Page struct{ ... }
type SetWhereOptions struct{ ... }
type Where []WhereOne
type WhereOne struct{ ... }
type GroupByResult struct{ ... }
type AggregateResult struct{ ... }
```

## Aggregate Query Examples

### Basic Aggregate Functions

```go
// Sum
total, err := gormx.Sum[Product](&gormx.Where{}, "price")

// Average
avgPrice, err := gormx.Avg[Product](&gormx.Where{}, "price")

// Min/Max
minPrice, err := gormx.Min[Product](&gormx.Where{}, "price")
maxPrice, err := gormx.Max[Product](&gormx.Where{}, "price")

// Count Distinct
uniqueCategories, err := gormx.CountDistinct[Product](&gormx.Where{}, "category")
```

### Group By Queries

```go
// Group by category with count
results, err := gormx.GroupBy[Product](
    []string{"category"}, 
    &gormx.Where{}, 
    []string{"COUNT(*) as count", "SUM(price) as sum"}
)

// Multiple aggregations in one query
aggregates, err := gormx.Aggregate[Product](
    "price", 
    &gormx.Where{}, 
    []string{"sum", "avg", "min", "max", "count"}
)
```

### With Where Conditions

```go
where := &gormx.Where{}
where.Set("category", "electronics")
where.Set("price", 100, &gormx.SetWhereOptions{IsFuzzy: true})

// Sum with conditions
total, err := gormx.Sum[Product](where, "price")
```
