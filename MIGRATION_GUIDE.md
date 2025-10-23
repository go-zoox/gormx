# Where Generic 迁移指南

## 快速开始

**好消息**：你的现有代码无需修改就能继续工作！

## 背景

我们升级了以下方法以支持泛型，使它们能同时接受 `map[any]any` 和 `*Where` 类型：

- `FindOne`
- `FindOneAndDelete`
- `FindOneAndUpdate`
- `FindOneOrCreate`
- `GetOrCreate`
- `Delete`
- `Exists`

## 迁移步骤

### 步骤 1: 无需修改现有代码 ✅

你的现有代码会继续工作：

```go
// ✅ 这段代码无需修改，继续正常工作
user, err := gormx.FindOne[User](map[any]any{"name": "Alice"})
```

### 步骤 2: 选择性地使用新功能（可选）

如果你需要复杂查询，现在可以使用 `*Where`：

```go
// 🆕 新功能：使用 *Where 进行复杂查询
where := gormx.NewWhere()
where.Set("name", "Alice", &gormx.SetWhereOptions{IsFuzzy: true})
where.Set("age", 25, &gormx.SetWhereOptions{IsNotEqual: true})
user, err := gormx.FindOne[User](where)
```

## 常见场景

### 场景 1: 简单等值查询

**推荐**：继续使用 `map[any]any`

```go
// ✅ 简单直观
user, err := gormx.FindOne[User](map[any]any{
    "name": "Alice",
    "age": 25,
})
```

### 场景 2: 模糊查询

**推荐**：使用 `*Where`

```go
// ✅ 强大的模糊查询
where := gormx.NewWhere()
where.Set("name", "Alice", &gormx.SetWhereOptions{IsFuzzy: true})
user, err := gormx.FindOne[User](where)
```

### 场景 3: IN 查询

**推荐**：使用 `*Where`

```go
// ✅ IN 查询
where := gormx.NewWhere()
where.Set("status", []string{"active", "pending"}, &gormx.SetWhereOptions{IsIn: true})
users, err := gormx.FindOne[User](where)
```

### 场景 4: NOT 查询

**推荐**：使用 `*Where`

```go
// ✅ NOT 查询
where := gormx.NewWhere()
where.Set("age", 18, &gormx.SetWhereOptions{IsNotEqual: true})
user, err := gormx.FindOne[User](where)
```

### 场景 5: 多条件组合

**推荐**：使用 `*Where`

```go
// ✅ 复杂的多条件查询
where := gormx.NewWhere()
where.Set("category", "electronics")
where.Set("name", "Pro", &gormx.SetWhereOptions{IsFuzzy: true})
where.Set("price", []int{999, 1999, 2999}, &gormx.SetWhereOptions{IsIn: true})
where.Set("status", "draft", &gormx.SetWhereOptions{IsNotEqual: true})
products, err := gormx.FindOne[Product](where)
```

## 选择指南

```
需要复杂查询？
    ├─ 是 → 使用 *Where
    │   ├─ 模糊查询（LIKE）
    │   ├─ IN / NOT IN
    │   ├─ 不等于（!=）
    │   ├─ 多条件组合
    │   └─ 全文搜索
    │
    └─ 否 → 使用 map[any]any
        └─ 简单等值查询（=）
```

## 性能考虑

### map[any]any
- ✅ 简单等值查询性能最优
- ✅ GORM 原生支持
- ✅ 最小化开销

### *Where
- ✅ 支持复杂查询
- ✅ SQL 构建优化
- ✅ 灵活性最高

## 代码示例对比

### Before（继续有效）

```go
// 查找
user, err := gormx.FindOne[User](map[any]any{"name": "Alice"})

// 删除
err = gormx.Delete[User](map[any]any{"id": 1})

// 检查存在
exists, err := gormx.Exists[User](map[any]any{"email": "alice@example.com"})

// 查找或创建
user, err := gormx.FindOneOrCreate[User](
    map[any]any{"email": "bob@example.com"},
    func(u *User) {
        u.Name = "Bob"
        u.Email = "bob@example.com"
    },
)
```

### After（新功能）

```go
// 使用 *Where 的复杂查询
where := gormx.NewWhere()
where.Set("name", "Alice", &gormx.SetWhereOptions{IsFuzzy: true})
where.Set("age", 18, &gormx.SetWhereOptions{IsNotEqual: true})
where.Set("status", []string{"active", "pending"}, &gormx.SetWhereOptions{IsIn: true})

user, err := gormx.FindOne[User](where)
```

## 常见问题

### Q1: 我必须修改现有代码吗？

**A**: 不需要！现有代码会继续工作。

### Q2: 什么时候应该使用 `*Where`？

**A**: 当你需要以下任一功能时：
- 模糊查询（LIKE）
- IN / NOT IN
- 不等于（!=）
- 多条件复杂组合
- 全文搜索

### Q3: `map[any]any` 还能用吗？

**A**: 当然可以！对于简单的等值查询，`map[any]any` 依然是最佳选择。

### Q4: 性能会受影响吗？

**A**: 不会。对于 `map[any]any`，我们依然使用 GORM 的原生支持。

### Q5: 如何判断使用哪种方式？

**A**: 
- 简单查询 → `map[any]any`
- 复杂查询 → `*Where`

## 完整示例

查看 `examples/where_generic_all/main.go` 获取完整的使用示例。

## 需要帮助？

- 📖 查看 [WHERE_GENERIC.md](WHERE_GENERIC.md) 获取完整文档
- 📖 查看 [README.md](README.md) 了解项目概览
- 💻 查看 [examples/](examples/) 目录获取更多示例

## 总结

✅ **零成本迁移** - 现有代码无需修改  
✅ **类型安全** - 编译时类型检查  
✅ **更灵活** - 支持复杂查询  
✅ **高性能** - 针对不同场景优化  

立即开始使用新功能，享受更强大、更灵活的查询能力！

