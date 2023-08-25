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
type OrderBy []OrderByOne
type OrderByOne struct{ ... }
type Page struct{ ... }
type SetWhereOptions struct{ ... }
type Where []WhereOne
type WhereOne struct{ ... }
```
