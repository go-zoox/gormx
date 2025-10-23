package main

import (
	"fmt"
	"log"

	"github.com/go-zoox/gormx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID    uint `gorm:"primarykey"`
	Name  string
	Email string
	Age   int
}

func main() {
	// Initialize database
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto migrate
	db.AutoMigrate(&User{})

	// Set the database
	gormx.SetDB(db)

	fmt.Println("=== Where Generic 完整示例 ===\n")

	// 1. FindOne - 使用 map[any]any
	fmt.Println("1. FindOne with map[any]any:")
	db.Create(&User{Name: "Alice", Email: "alice@example.com", Age: 25})
	user1, err := gormx.FindOne[User](map[any]any{"name": "Alice"})
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Found: %+v\n", user1)
	}

	// 2. FindOne - 使用 *Where
	fmt.Println("\n2. FindOne with *Where:")
	where2 := gormx.NewWhere()
	where2.Set("age", 25)
	user2, err := gormx.FindOne[User](where2)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Found: %+v\n", user2)
	}

	// 3. Exists - 使用 map[any]any
	fmt.Println("\n3. Exists with map[any]any:")
	exists1, err := gormx.Exists[User](map[any]any{"name": "Alice"})
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Exists: %v\n", exists1)
	}

	// 4. Exists - 使用 *Where
	fmt.Println("\n4. Exists with *Where:")
	where4 := gormx.NewWhere()
	where4.Set("name", "Bob")
	exists2, err := gormx.Exists[User](where4)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Exists: %v\n", exists2)
	}

	// 5. FindOneOrCreate - 使用 map[any]any
	fmt.Println("\n5. FindOneOrCreate with map[any]any:")
	user5, err := gormx.FindOneOrCreate[User](
		map[any]any{"email": "bob@example.com"},
		func(u *User) {
			u.Name = "Bob"
			u.Email = "bob@example.com"
			u.Age = 30
		},
	)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Created: %+v\n", user5)
	}

	// 6. FindOneOrCreate - 使用 *Where
	fmt.Println("\n6. FindOneOrCreate with *Where:")
	where6 := gormx.NewWhere()
	where6.Set("email", "charlie@example.com")
	user6, err := gormx.FindOneOrCreate[User](
		where6,
		func(u *User) {
			u.Name = "Charlie"
			u.Email = "charlie@example.com"
			u.Age = 35
		},
	)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Created: %+v\n", user6)
	}

	// 7. GetOrCreate - 使用 map[any]any
	fmt.Println("\n7. GetOrCreate with map[any]any:")
	user7, err := gormx.GetOrCreate[User](
		map[any]any{"email": "david@example.com"},
		func(u *User) {
			u.Name = "David"
			u.Email = "david@example.com"
			u.Age = 40
		},
	)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Created: %+v\n", user7)
	}

	// 8. FindOneAndUpdate - 使用 map[any]any
	fmt.Println("\n8. FindOneAndUpdate with map[any]any:")
	user8, err := gormx.FindOneAndUpdate[User](
		map[any]any{"name": "Alice"},
		func(u *User) {
			u.Age = 26
			db.Save(u)
		},
	)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Updated: %+v\n", user8)
	}

	// 9. FindOneAndUpdate - 使用 *Where（复杂条件）
	fmt.Println("\n9. FindOneAndUpdate with *Where (complex):")
	where9 := gormx.NewWhere()
	where9.Set("age", 30, &gormx.SetWhereOptions{IsNotEqual: true})
	user9, err := gormx.FindOneAndUpdate[User](
		where9,
		func(u *User) {
			u.Age = u.Age + 1
			db.Save(u)
		},
	)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Updated: %+v\n", user9)
	}

	// 10. FindOneAndDelete - 使用 map[any]any
	fmt.Println("\n10. FindOneAndDelete with map[any]any:")
	db.Create(&User{Name: "ToDelete", Email: "delete@example.com", Age: 50})
	user10, err := gormx.FindOneAndDelete[User](map[any]any{"name": "ToDelete"})
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Deleted: %+v\n", user10)
	}

	// 11. Delete - 使用 *Where
	fmt.Println("\n11. Delete with *Where:")
	db.Create(&User{Name: "ToDelete2", Email: "delete2@example.com", Age: 60})
	where11 := gormx.NewWhere()
	where11.Set("email", "delete2@example.com")
	err = gormx.Delete[User](where11)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Deleted successfully\n")
	}

	// 显示最终数据
	fmt.Println("\n=== 最终数据库中的用户 ===")
	var allUsers []User
	db.Find(&allUsers)
	for i, u := range allUsers {
		fmt.Printf("%d. %+v\n", i+1, u)
	}
}
