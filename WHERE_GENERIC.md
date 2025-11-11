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

### 方式5：使用 *Where（复杂场景 - 范围查询）

```go
where := gormx.NewWhere()

// 年龄范围 18-65（默认是闭区间 [A, B]）
where.Set("age", []int{18, 65}, &gormx.SetWhereOptions{IsRange: true})
// 生成: age BETWEEN 18 AND 65

// 指定范围模式
where.Set("age", []int{18, 65}, &gormx.SetWhereOptions{
    IsRange:   true,
    RangeMode: gormx.RangeModeClosed, // [A, B] 闭区间（默认）
})

// 日期范围（字符串）
where.Set("created_at", []string{"2023-01-01", "2023-12-31"}, &gormx.SetWhereOptions{IsRange: true})

// 日期范围（time.Time）
start := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)
where.Set("created_at", []interface{}{start, end}, &gormx.SetWhereOptions{IsRange: true})

users, total, err := gormx.Find[User](1, 10, where, nil)
```

#### 范围模式（Range Modes）

IsRange 支持四种范围模式：

```go
// 1. RangeModeClosed - [A, B] 闭区间（默认）
// 包含两端：A <= field <= B
where.Set("age", []int{18, 65}, &gormx.SetWhereOptions{
    IsRange:   true,
    RangeMode: gormx.RangeModeClosed,
})
// 生成: age BETWEEN 18 AND 65

// 2. RangeModeOpen - (A, B) 开区间
// 不包含两端：A < field < B
where.Set("age", []int{18, 65}, &gormx.SetWhereOptions{
    IsRange:   true,
    RangeMode: gormx.RangeModeOpen,
})
// 生成: (age > 18 AND age < 65)

// 3. RangeModeLeftClosed - [A, B) 左闭右开
// 包含左端，不包含右端：A <= field < B
where.Set("age", []int{18, 65}, &gormx.SetWhereOptions{
    IsRange:   true,
    RangeMode: gormx.RangeModeLeftClosed,
})
// 生成: (age >= 18 AND age < 65)

// 4. RangeModeRightClosed - (A, B] 左开右闭
// 不包含左端，包含右端：A < field <= B
where.Set("age", []int{18, 65}, &gormx.SetWhereOptions{
    IsRange:   true,
    RangeMode: gormx.RangeModeRightClosed,
})
// 生成: (age > 18 AND age <= 65)
```

**使用场景：**
- **RangeModeClosed** `[A, B]`：通用范围查询，如年龄段 18-65 岁
- **RangeModeOpen** `(A, B)`：需要排除边界值的场景
- **RangeModeLeftClosed** `[A, B)`：时间范围查询，如某一天的所有数据
- **RangeModeRightClosed** `(A, B]`：分页偏移等场景

### 方式6：使用 *Where（复杂场景 - 比较操作符）

```go
where := gormx.NewWhere()

// 大于 (>)
where.Set("age", 18, &gormx.SetWhereOptions{IsGreaterThan: true})

// 小于 (<)
where.Set("age", 65, &gormx.SetWhereOptions{IsLessThan: true})

// 大于等于 (>=)
where.Set("score", 60, &gormx.SetWhereOptions{IsGreaterOrEqualThan: true})

// 小于等于 (<=)
where.Set("price", 100, &gormx.SetWhereOptions{IsLessOrEqualThan: true})

// 组合使用：age >= 18 AND age <= 65
where := gormx.NewWhere()
where.Add("age", 18, &gormx.SetWhereOptions{IsGreaterOrEqualThan: true})
where.Add("age", 65, &gormx.SetWhereOptions{IsLessOrEqualThan: true})

users, total, err := gormx.Find[User](1, 10, where, nil)
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

## 支持的查询类型

### 基本查询
- **IsEqual**: 等值查询 `field = ?`
- **IsNotEqual**: 不等查询 `field != ?`
- **IsFuzzy**: 模糊查询 `field ILike ?`
- **IsIn**: IN 查询 `field in (?)`
- **IsNotIn**: NOT IN 查询 `field not in (?)`
- **IsRange**: 范围查询 `field BETWEEN ? AND ?`
- **IsGreaterThan**: 大于查询 `field > ?`
- **IsLessThan**: 小于查询 `field < ?`
- **IsGreaterOrEqualThan**: 大于等于查询 `field >= ?`
- **IsLessOrEqualThan**: 小于等于查询 `field <= ?`
- **IsPlain**: 原生 SQL 条件

### URL 查询参数语法（通过 Params）

在 HTTP 请求中，可以通过以下语法使用查询条件：

```
# 等值查询
?name=Alice

# 模糊查询
?name=Alice:*

# 不等查询
?status=inactive:!

# IN 查询
?age=20,25,30:in

# NOT IN 查询
?age=20,25,30:!in

# 范围查询
?age=18,65:range      # 默认闭区间 [A, B]
?age=18,65:range[]    # 闭区间 [A, B]
?age=18,65:range()    # 开区间 (A, B)
?age=18,65:range[)    # 左闭右开 [A, B)
?age=18,65:range(]    # 左开右闭 (A, B]
?created_at=2023-01-01,2023-12-31:range[)  # 日期范围 [A, B)

# 比较操作符
?age=18:>          # age > 18
?age=65:<          # age < 65
?score=60:>=       # score >= 60
?price=100:<=      # price <= 100
```

示例：
```go
// 在 HTTP Handler 中
params := gormx.NewParams(ctx)
where := params.Where()

// URL: /users?age=18,65:range&status=active
// 会自动转换为：
// WHERE age BETWEEN 18 AND 65 AND status = 'active'

// URL: /users?age=18,65:range()&status=active
// 会自动转换为：
// WHERE (age > 18 AND age < 65) AND status = 'active'

// URL: /orders?created_at=2023-01-01,2023-12-31:range[)
// 会自动转换为：
// WHERE (created_at >= '2023-01-01' AND created_at < '2023-12-31')

// URL: /users?age=18:>&score=60:>=&status=active
// 会自动转换为：
// WHERE age > 18 AND score >= 60 AND status = 'active'

// URL: /products?price=100:<=&quantity=10:>
// 会自动转换为：
// WHERE price <= 100 AND quantity > 10
```

## 完整示例

### 基本用法
参见 `examples/where_generic/main.go`

### 范围查询
参见 `examples/where_range/main.go`

### 范围模式
参见 `examples/where_range_modes/main.go`

### 比较操作符
参见 `examples/where_comparison/main.go`

