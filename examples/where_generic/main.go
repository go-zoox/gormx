package main

import (
	"fmt"

	"github.com/go-zoox/gormx"
)

type User struct {
	ID   uint `gorm:"primarykey"`
	Name string
	Age  int
}

func main() {
	// 示例1：使用 map[any]any
	user1, err := gormx.FindOne[User](map[any]any{
		"name": "Alice",
		"age":  25,
	})
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("User1: %+v\n", user1)
	}

	// 示例2：使用 *Where（简单条件）
	where2 := gormx.NewWhere()
	where2.Set("name", "Bob")
	where2.Set("age", 30)

	user2, err := gormx.FindOne[User](where2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("User2: %+v\n", user2)
	}

	// 示例3：使用 *Where（复杂条件 - 模糊查询）
	where3 := gormx.NewWhere()
	where3.Set("name", "Charlie", &gormx.SetWhereOptions{IsFuzzy: true})
	where3.Set("age", 35, &gormx.SetWhereOptions{IsNotEqual: true})

	user3, err := gormx.FindOne[User](where3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("User3: %+v\n", user3)
	}

	// 示例4：使用 *Where（IN 查询）
	where4 := gormx.NewWhere()
	where4.Set("age", []int{20, 25, 30}, &gormx.SetWhereOptions{IsIn: true})

	user4, err := gormx.FindOne[User](where4)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("User4: %+v\n", user4)
	}
}
