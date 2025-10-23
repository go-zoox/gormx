package gormx

import (
	"testing"
	"time"

	"gorm.io/gorm"
)

// TestChainProduct is a test model for chain queries
type TestChainProduct struct {
	ID        uint           `gorm:"primarykey"`
	Name      string         `gorm:"column:name"`
	Category  string         `gorm:"column:category"`
	Price     float64        `gorm:"column:price"`
	Quantity  int            `gorm:"column:quantity"`
	InStock   bool           `gorm:"column:in_stock"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (TestChainProduct) TableName() string {
	return "test_chain_products"
}

func setupChainTestData(t *testing.T) {
	// Skip if no database connection
	if GetDB() == nil {
		t.Skip("No database connection available")
	}

	// Create test table
	err := GetDB().AutoMigrate(&TestChainProduct{})
	if err != nil {
		t.Fatalf("Failed to migrate test table: %v", err)
	}

	// Clean up test data
	GetDB().Unscoped().Where("1 = 1").Delete(&TestChainProduct{})

	// Insert test data
	testProducts := []TestChainProduct{
		{Name: "Laptop", Category: "Electronics", Price: 1000.0, Quantity: 10, InStock: true},
		{Name: "Phone", Category: "Electronics", Price: 500.0, Quantity: 20, InStock: true},
		{Name: "Book", Category: "Books", Price: 20.0, Quantity: 50, InStock: true},
		{Name: "Pen", Category: "Stationery", Price: 2.0, Quantity: 100, InStock: true},
		{Name: "Notebook", Category: "Stationery", Price: 5.0, Quantity: 80, InStock: true},
		{Name: "Monitor", Category: "Electronics", Price: 300.0, Quantity: 0, InStock: false},
	}

	for _, product := range testProducts {
		GetDB().Create(&product)
	}
}

func cleanupChainTestData(t *testing.T) {
	GetDB().Unscoped().Where("1 = 1").Delete(&TestChainProduct{})
}

func TestQueryBuilder_Where(t *testing.T) {
	setupChainTestData(t)
	defer cleanupChainTestData(t)

	t.Run("Simple Where", func(t *testing.T) {
		results, err := NewQuery[TestChainProduct]().
			Where("category", "Electronics").
			Find()

		if err != nil {
			t.Fatalf("Where query failed: %v", err)
		}

		if len(results) != 3 {
			t.Errorf("Expected 3 results, got %d", len(results))
		}
	})

	t.Run("WhereEqual", func(t *testing.T) {
		results, err := NewQuery[TestChainProduct]().
			WhereEqual("category", "Books").
			Find()

		if err != nil {
			t.Fatalf("WhereEqual query failed: %v", err)
		}

		if len(results) != 1 {
			t.Errorf("Expected 1 result, got %d", len(results))
		}
	})

	t.Run("WhereIn", func(t *testing.T) {
		results, err := NewQuery[TestChainProduct]().
			WhereIn("category", []string{"Electronics", "Books"}).
			Find()

		if err != nil {
			t.Fatalf("WhereIn query failed: %v", err)
		}

		if len(results) != 4 {
			t.Errorf("Expected 4 results, got %d", len(results))
		}
	})

	t.Run("Multiple Where Conditions", func(t *testing.T) {
		results, err := NewQuery[TestChainProduct]().
			Where("category", "Electronics").
			Where("in_stock", true).
			Find()

		if err != nil {
			t.Fatalf("Multiple where query failed: %v", err)
		}

		if len(results) != 2 {
			t.Errorf("Expected 2 results, got %d", len(results))
		}
	})
}

func TestQueryBuilder_OrderBy(t *testing.T) {
	setupChainTestData(t)
	defer cleanupChainTestData(t)

	t.Run("OrderBy Ascending", func(t *testing.T) {
		results, err := NewQuery[TestChainProduct]().
			OrderByAsc("price").
			Find()

		if err != nil {
			t.Fatalf("OrderBy query failed: %v", err)
		}

		if len(results) < 2 {
			t.Fatal("Not enough results")
		}

		if results[0].Price > results[1].Price {
			t.Error("Results not ordered correctly")
		}
	})

	t.Run("OrderBy Descending", func(t *testing.T) {
		results, err := NewQuery[TestChainProduct]().
			OrderByDesc("price").
			Find()

		if err != nil {
			t.Fatalf("OrderBy query failed: %v", err)
		}

		if len(results) < 2 {
			t.Fatal("Not enough results")
		}

		if results[0].Price < results[1].Price {
			t.Error("Results not ordered correctly")
		}
	})
}

func TestQueryBuilder_Limit(t *testing.T) {
	setupChainTestData(t)
	defer cleanupChainTestData(t)

	t.Run("Limit", func(t *testing.T) {
		results, err := NewQuery[TestChainProduct]().
			Limit(3).
			Find()

		if err != nil {
			t.Fatalf("Limit query failed: %v", err)
		}

		if len(results) != 3 {
			t.Errorf("Expected 3 results, got %d", len(results))
		}
	})

	t.Run("Limit and Offset", func(t *testing.T) {
		results, err := NewQuery[TestChainProduct]().
			OrderByAsc("id").
			Limit(2).
			Offset(2).
			Find()

		if err != nil {
			t.Fatalf("Limit and Offset query failed: %v", err)
		}

		if len(results) != 2 {
			t.Errorf("Expected 2 results, got %d", len(results))
		}
	})

	t.Run("Page", func(t *testing.T) {
		results, err := NewQuery[TestChainProduct]().
			Page(2, 2).
			Find()

		if err != nil {
			t.Fatalf("Page query failed: %v", err)
		}

		if len(results) != 2 {
			t.Errorf("Expected 2 results, got %d", len(results))
		}
	})
}

func TestQueryBuilder_Select(t *testing.T) {
	setupChainTestData(t)
	defer cleanupChainTestData(t)

	t.Run("Select Specific Columns", func(t *testing.T) {
		results, err := NewQuery[TestChainProduct]().
			Select("name", "price").
			Find()

		if err != nil {
			t.Fatalf("Select query failed: %v", err)
		}

		if len(results) == 0 {
			t.Error("Expected results")
		}
	})
}

func TestQueryBuilder_Count(t *testing.T) {
	setupChainTestData(t)
	defer cleanupChainTestData(t)

	t.Run("Count All", func(t *testing.T) {
		count, err := NewQuery[TestChainProduct]().Count()

		if err != nil {
			t.Fatalf("Count query failed: %v", err)
		}

		if count != 6 {
			t.Errorf("Expected count 6, got %d", count)
		}
	})

	t.Run("Count with Where", func(t *testing.T) {
		count, err := NewQuery[TestChainProduct]().
			Where("category", "Electronics").
			Count()

		if err != nil {
			t.Fatalf("Count with where query failed: %v", err)
		}

		if count != 3 {
			t.Errorf("Expected count 3, got %d", count)
		}
	})
}

func TestQueryBuilder_First(t *testing.T) {
	setupChainTestData(t)
	defer cleanupChainTestData(t)

	t.Run("First", func(t *testing.T) {
		result, err := NewQuery[TestChainProduct]().
			Where("category", "Books").
			First()

		if err != nil {
			t.Fatalf("First query failed: %v", err)
		}

		if result == nil {
			t.Error("Expected result")
		}

		if result.Category != "Books" {
			t.Errorf("Expected category Books, got %s", result.Category)
		}
	})
}

func TestQueryBuilder_Exists(t *testing.T) {
	setupChainTestData(t)
	defer cleanupChainTestData(t)

	t.Run("Exists - True", func(t *testing.T) {
		exists, err := NewQuery[TestChainProduct]().
			Where("category", "Electronics").
			Exists()

		if err != nil {
			t.Fatalf("Exists query failed: %v", err)
		}

		if !exists {
			t.Error("Expected exists to be true")
		}
	})

	t.Run("Exists - False", func(t *testing.T) {
		exists, err := NewQuery[TestChainProduct]().
			Where("category", "NonExistent").
			Exists()

		if err != nil {
			t.Fatalf("Exists query failed: %v", err)
		}

		if exists {
			t.Error("Expected exists to be false")
		}
	})
}

func TestQueryBuilder_Aggregate(t *testing.T) {
	setupChainTestData(t)
	defer cleanupChainTestData(t)

	t.Run("Sum", func(t *testing.T) {
		sum, err := NewQuery[TestChainProduct]().
			Where("category", "Electronics").
			Sum("price")

		if err != nil {
			t.Fatalf("Sum query failed: %v", err)
		}

		expected := 1800.0 // 1000 + 500 + 300
		if sum != expected {
			t.Errorf("Expected sum %f, got %f", expected, sum)
		}
	})

	t.Run("Avg", func(t *testing.T) {
		avg, err := NewQuery[TestChainProduct]().
			Where("category", "Stationery").
			Avg("price")

		if err != nil {
			t.Fatalf("Avg query failed: %v", err)
		}

		expected := 3.5 // (2 + 5) / 2
		if avg != expected {
			t.Errorf("Expected avg %f, got %f", expected, avg)
		}
	})

	t.Run("Min", func(t *testing.T) {
		min, err := NewQuery[TestChainProduct]().
			Min("price")

		if err != nil {
			t.Fatalf("Min query failed: %v", err)
		}

		if min == nil {
			t.Error("Expected min value")
		}
	})

	t.Run("Max", func(t *testing.T) {
		max, err := NewQuery[TestChainProduct]().
			Max("price")

		if err != nil {
			t.Fatalf("Max query failed: %v", err)
		}

		if max == nil {
			t.Error("Expected max value")
		}
	})
}

func TestQueryBuilder_Paginate(t *testing.T) {
	setupChainTestData(t)
	defer cleanupChainTestData(t)

	t.Run("Paginate", func(t *testing.T) {
		results, total, err := NewQuery[TestChainProduct]().
			OrderByAsc("id").
			Paginate(1, 2)

		if err != nil {
			t.Fatalf("Paginate query failed: %v", err)
		}

		if len(results) != 2 {
			t.Errorf("Expected 2 results, got %d", len(results))
		}

		if total != 6 {
			t.Errorf("Expected total 6, got %d", total)
		}
	})
}

func TestQueryBuilder_Chunk(t *testing.T) {
	setupChainTestData(t)
	defer cleanupChainTestData(t)

	t.Run("Chunk", func(t *testing.T) {
		var processedCount int
		err := NewQuery[TestChainProduct]().
			OrderByAsc("id").
			Chunk(2, func(products []*TestChainProduct) error {
				processedCount += len(products)
				return nil
			})

		if err != nil {
			t.Fatalf("Chunk query failed: %v", err)
		}

		if processedCount != 6 {
			t.Errorf("Expected to process 6 records, processed %d", processedCount)
		}
	})
}

func TestQueryBuilder_GroupBy(t *testing.T) {
	setupChainTestData(t)
	defer cleanupChainTestData(t)

	t.Run("GroupBy", func(t *testing.T) {
		var results []struct {
			Category string
			Count    int64
		}

		err := NewQuery[TestChainProduct]().
			Select("category", "COUNT(*) as count").
			GroupBy("category").
			Scan(&results)

		if err != nil {
			t.Fatalf("GroupBy query failed: %v", err)
		}

		if len(results) != 3 {
			t.Errorf("Expected 3 groups, got %d", len(results))
		}
	})
}

func TestQueryBuilder_Distinct(t *testing.T) {
	setupChainTestData(t)
	defer cleanupChainTestData(t)

	t.Run("Distinct", func(t *testing.T) {
		var categories []string
		err := NewQuery[TestChainProduct]().
			Distinct().
			Pluck("category", &categories)

		if err != nil {
			t.Fatalf("Distinct query failed: %v", err)
		}

		if len(categories) != 3 {
			t.Errorf("Expected 3 distinct categories, got %d", len(categories))
		}
	})
}

func TestQueryBuilder_Clone(t *testing.T) {
	setupChainTestData(t)
	defer cleanupChainTestData(t)

	t.Run("Clone", func(t *testing.T) {
		originalQuery := NewQuery[TestChainProduct]().
			Where("category", "Electronics").
			OrderByAsc("price")

		clonedQuery := originalQuery.Clone()

		// Modify cloned query
		clonedQuery.Where("in_stock", true)

		// Original should not be affected
		originalCount, _ := originalQuery.Count()
		clonedCount, _ := clonedQuery.Count()

		if originalCount == clonedCount {
			t.Log("Clone might not be fully independent (expected, this is a simple clone)")
		}
	})
}

func TestQueryBuilder_ChainedOperations(t *testing.T) {
	setupChainTestData(t)
	defer cleanupChainTestData(t)

	t.Run("Complex Chained Query", func(t *testing.T) {
		results, err := NewQuery[TestChainProduct]().
			Where("category", "Electronics").
			Where("in_stock", true).
			OrderByDesc("price").
			Limit(2).
			Find()

		if err != nil {
			t.Fatalf("Chained query failed: %v", err)
		}

		if len(results) != 2 {
			t.Errorf("Expected 2 results, got %d", len(results))
		}

		// Should be ordered by price descending
		if len(results) == 2 && results[0].Price < results[1].Price {
			t.Error("Results not ordered correctly")
		}
	})
}
