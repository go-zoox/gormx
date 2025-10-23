# Where 泛型支持

## 概述

从此版本开始，所有 `where` 参数的方法都支持使用泛型来同时接受 `map[any]any` 和 `*Where` 两种类型。

## 实现原理

### 1. 类型约束（Type Constraint）

定义了一个泛型类型约束 `WhereCondition`：

```go
type WhereCondition interface {
	map[any]any | *Where
}
```

这个约束限制了泛型参数只能是 `map[any]any` 或 `*Where` 两种类型之一。

### 2. 类型转换函数

提供了 `ToWhere` 函数来统一处理两种类型：

```go
func ToWhere[T WhereCondition](condition T) *Where
```

该函数会：
- 如果输入是 `*Where`，直接返回
- 如果输入是 `map[any]any`，转换为 `*Where` 对象

### 3. 方法签名更新

原来的方法：
```go
func FindOne[T any](where map[any]any) (*T, error)
```

更新后的方法：
```go
func FindOne[T any, W WhereCondition](where W) (*T, error)
```

## 使用示例

### 方式1：使用 map[any]any（简单场景）

```go
user, err := gormx.FindOne[User](map[any]any{
	"name": "Alice",
	"age":  25,
})
```

### 方式2：使用 *Where（简单场景）

```go
where := gormx.NewWhere()
where.Set("name", "Bob")
where.Set("age", 30)

user, err := gormx.FindOne[User](where)
```

### 方式3：使用 *Where（复杂场景 - 模糊查询）

```go
where := gormx.NewWhere()
where.Set("name", "Charlie", &gormx.SetWhereOptions{IsFuzzy: true})
where.Set("age", 35, &gormx.SetWhereOptions{IsNotEqual: true})

user, err := gormx.FindOne[User](where)
```

### 方式4：使用 *Where（复杂场景 - IN 查询）

```go
where := gormx.NewWhere()
where.Set("age", []int{20, 25, 30}, &gormx.SetWhereOptions{IsIn: true})

user, err := gormx.FindOne[User](where)
```

## 优势

1. **类型安全**：编译时就能检查类型错误
2. **向后兼容**：现有使用 `map[any]any` 的代码无需修改
3. **灵活性**：可以根据场景选择合适的类型
4. **统一接口**：所有方法都使用同样的模式

## 适用方法

以下方法都可以应用这个模式：

### 查询方法
- `FindOne[T, W]`
- `FindOneAndDelete[T, W]`
- `FindOneAndUpdate[T, W]`
- `FindOneOrCreate[T, W]`
- `GetOrCreate[T, W]`

### 删除方法
- `Delete[T, W]`

### 判断方法
- `Exists[T, W]`

### 列表查询方法（需要调整）
- `Find[T, W]` - 需要额外处理 nil 的情况
- `FindAll[T, W]` - 需要额外处理 nil 的情况
- `List[T, W]` - 需要额外处理 nil 的情况
- `ListALL[T, W]` - 需要额外处理 nil 的情况

## 注意事项

### nil 值处理

对于原本接受 `*Where` 的方法（如 `FindAll`），如果需要支持 nil 值，有两种方案：

**方案A：使用指针约束**
```go
type WhereConditionPtr interface {
	map[any]any | *Where | nil
}
```

**方案B：保持可选参数**
```go
func FindAll[T any](where *Where, orderBy *OrderBy) (data []*T, err error)
```

### 性能考虑

- 使用 `map[any]any` 时，如果是简单的等值查询，内部会直接使用 GORM 的原生 map 支持，性能最优
- 使用 `*Where` 时，会构建完整的 SQL WHERE 子句，支持复杂查询
- 两种方式都经过优化，运行时开销很小

## 迁移建议

### 已有代码

现有代码无需修改，因为泛型参数会自动推断：

```go
// 这行代码依然有效
user, err := gormx.FindOne[User](map[any]any{"name": "Alice"})
```

### 新代码

建议根据场景选择：
- **简单查询**：使用 `map[any]any`
- **复杂查询**：使用 `*Where`

## 完整示例

参见 `examples/where_generic/main.go`

