package gormx

import "fmt"

// FindOne finds one record.
// Supports both map[any]any and *Where as where condition.
func FindOne[T any, W WhereCondition](where W) (*T, error) {
	var f T

	// Convert to *Where for unified processing
	w := ToWhere(where)

	// If it's a simple map (has no complex conditions), use GORM's native map support
	if len(w.Items) > 0 {
		// Check if all items are simple equality conditions
		simpleMap := make(map[any]any)
		isSimple := true

		for _, item := range w.Items {
			if item.IsFuzzy || item.IsIn || item.IsNotIn || item.IsPlain || item.IsFullTextSearch || item.IsNotEqual {
				isSimple = false
				break
			}
			simpleMap[item.Key] = item.Value
		}

		if isSimple {
			// Use simple map query
			if err := GetDB().First(&f, simpleMap).Error; err != nil {
				return nil, err
			}
			return &f, nil
		}
	}

	// Use complex conditions
	return FindOneWithComplexConditions[T](w, nil)
}

// FindOneWithComplexConditions finds one record.
func FindOneWithComplexConditions[T any](where *Where, orderBy *OrderBy) (*T, error) {
	var f T
	dataTx := GetDB()

	if where != nil {
		whereClause, whereValues, errx := where.Build()
		if errx != nil {
			return nil, errx
		}

		if whereClause != "" {
			dataTx = dataTx.Where(whereClause, whereValues...)
		}
	}

	if orderBy != nil {
		for _, order := range *orderBy {
			// fmt.Println("order by:", order.Key, order.IsDESC)
			orderMod := "ASC"
			if order.IsDESC {
				orderMod = "DESC"
			}

			orderStr := fmt.Sprintf("%s %s", order.Key, orderMod)
			dataTx = dataTx.Order(orderStr)
		}
	}

	if err := dataTx.First(&f).Error; err != nil {
		return nil, err
	}

	return &f, nil
}
