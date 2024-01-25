package gormx

import "fmt"

// FindOne finds one record.
func FindOne[T any](where map[any]any) (*T, error) {
	var f T
	if err := GetDB().First(&f, where).Error; err != nil {
		return nil, err
	}

	return &f, nil
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
