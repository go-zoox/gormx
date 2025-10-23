package gormx

import (
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

// TestProduct is a test model for aggregate queries
type TestProduct struct {
	ID        uint           `gorm:"primarykey"`
	Name      string         `gorm:"column:name"`
	Category  string         `gorm:"column:category"`
	Price     float64        `gorm:"column:price"`
	Quantity  int            `gorm:"column:quantity"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (TestProduct) TableName() string {
	return "test_products"
}

func TestAggregateFunctions(t *testing.T) {
	// Skip if no database connection
	if GetDB() == nil {
		t.Skip("No database connection available")
	}

	// Create test table
	err := GetDB().AutoMigrate(&TestProduct{})
	if err != nil {
		t.Fatalf("Failed to migrate test table: %v", err)
	}

	// Clean up test data
	GetDB().Unscoped().Where("1 = 1").Delete(&TestProduct{})

	// Insert test data
	testProducts := []TestProduct{
		{Name: "Product 1", Category: "Electronics", Price: 100.0, Quantity: 10},
		{Name: "Product 2", Category: "Electronics", Price: 200.0, Quantity: 5},
		{Name: "Product 3", Category: "Books", Price: 50.0, Quantity: 20},
		{Name: "Product 4", Category: "Books", Price: 75.0, Quantity: 15},
		{Name: "Product 5", Category: "Clothing", Price: 150.0, Quantity: 8},
	}

	for _, product := range testProducts {
		GetDB().Create(&product)
	}

	t.Run("Test Sum", func(t *testing.T) {
		total, err := Sum[TestProduct]("price", nil)
		if err != nil {
			t.Fatalf("Sum failed: %v", err)
		}
		expected := 575.0 // 100 + 200 + 50 + 75 + 150
		if total != expected {
			t.Errorf("Expected sum %f, got %f", expected, total)
		}
	})

	t.Run("Test Avg", func(t *testing.T) {
		avg, err := Avg[TestProduct]("price", nil)
		if err != nil {
			t.Fatalf("Avg failed: %v", err)
		}
		expected := 115.0 // 575 / 5
		if avg != expected {
			t.Errorf("Expected avg %f, got %f", expected, avg)
		}
	})

	t.Run("Test Min", func(t *testing.T) {
		min, err := Min[TestProduct]("price", nil)
		if err != nil {
			t.Fatalf("Min failed: %v", err)
		}
		expected := 50.0
		if min != expected {
			t.Errorf("Expected min %f, got %v", expected, min)
		}
	})

	t.Run("Test Max", func(t *testing.T) {
		max, err := Max[TestProduct]("price", nil)
		if err != nil {
			t.Fatalf("Max failed: %v", err)
		}
		expected := 200.0
		if max != expected {
			t.Errorf("Expected max %f, got %v", expected, max)
		}
	})

	t.Run("Test CountDistinct", func(t *testing.T) {
		count, err := CountDistinct[TestProduct]("category", nil)
		if err != nil {
			t.Fatalf("CountDistinct failed: %v", err)
		}
		expected := int64(3) // Electronics, Books, Clothing
		if count != expected {
			t.Errorf("Expected count %d, got %d", expected, count)
		}
	})

	t.Run("Test Sum with Where", func(t *testing.T) {
		where := NewWhere()
		where.Set("category", "Electronics")

		total, err := Sum[TestProduct]("price", where)
		if err != nil {
			t.Fatalf("Sum with where failed: %v", err)
		}
		expected := 300.0 // 100 + 200
		if total != expected {
			t.Errorf("Expected sum %f, got %f", expected, total)
		}
	})

	t.Run("Test GroupBy", func(t *testing.T) {
		results, err := GroupBy[TestProduct](
			[]string{"category"},
			nil,
			[]string{"COUNT(*) as count", "SUM(price) as sum"},
		)
		if err != nil {
			t.Fatalf("GroupBy failed: %v", err)
		}

		if len(results) != 3 {
			t.Errorf("Expected 3 groups, got %d", len(results))
		}

		// Find Electronics group
		var electronicsGroup *GroupByResult
		for _, result := range results {
			if result.Group["category"] == "Electronics" {
				electronicsGroup = &result
				break
			}
		}

		if electronicsGroup == nil {
			t.Fatal("Electronics group not found")
		}

		if electronicsGroup.Count != 2 {
			t.Errorf("Expected Electronics count 2, got %d", electronicsGroup.Count)
		}

		if electronicsGroup.Sum != 300.0 {
			t.Errorf("Expected Electronics sum 300.0, got %f", electronicsGroup.Sum)
		}
	})

	t.Run("Test Aggregate", func(t *testing.T) {
		results, err := Aggregate[TestProduct](
			"price",
			nil,
			[]string{"sum", "avg", "min", "max", "count"},
		)
		if err != nil {
			t.Fatalf("Aggregate failed: %v", err)
		}

		if results["sum"] != 575.0 {
			t.Errorf("Expected sum 575.0, got %v", results["sum"])
		}

		if results["avg"] != 115.0 {
			t.Errorf("Expected avg 115.0, got %v", results["avg"])
		}

		if results["min"] != 50.0 {
			t.Errorf("Expected min 50.0, got %v", results["min"])
		}

		if results["max"] != 200.0 {
			t.Errorf("Expected max 200.0, got %v", results["max"])
		}

		if results["count"] != int64(5) {
			t.Errorf("Expected count 5, got %v", results["count"])
		}
	})

	// Clean up test data
	GetDB().Unscoped().Where("1 = 1").Delete(&TestProduct{})
}

func TestGetFieldType(t *testing.T) {
	fieldType, err := GetFieldType[TestProduct]("Price")
	if err != nil {
		t.Fatalf("GetFieldType failed: %v", err)
	}

	if fieldType.Kind() != reflect.Float64 {
		t.Errorf("Expected field type Float64, got %v", fieldType.Kind())
	}

	// Test non-existent field
	_, err = GetFieldType[TestProduct]("NonExistentField")
	if err == nil {
		t.Error("Expected error for non-existent field")
	}
}
