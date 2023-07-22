package gormx

import "fmt"

// List lists records.
func List[T any](page, pageSize uint, where *Where, orderBy *OrderBy) (data []*T, total int64, err error) {
	offset := int((page - 1) * pageSize)
	limit := int(pageSize)

	// whereClauses := []string{}
	// whereValues := []interface{}{}
	// for _, w := range *where {
	// 	if w.IsFuzzy {
	// 		whereClauses = append(whereClauses, fmt.Sprintf("%s ILike ?", w.Key))
	// 		whereValues = append(whereValues, fmt.Sprintf("%%%s%%", w.Value))
	// 	} else if w.isNot {
	// 		whereClauses = append(whereClauses, fmt.Sprintf("%s != ?", w.Key))
	// 		whereValues = append(whereValues, w.Value)
	// 	} else if w.isIn {
	// 		whereClauses = append(whereClauses, fmt.Sprintf("%s in (?)", w.Key))
	// 		whereValues = append(whereValues, w.Value)
	// 	} else {
	// 		whereClauses = append(whereClauses, fmt.Sprintf("%s = ?", w.Key))
	// 		whereValues = append(whereValues, w.Value)
	// 	}
	// }
	// whereClause := strings.Join(whereClauses, " AND ")

	whereClause, whereValues := where.Build()

	countTx := GetDB().Model(new(T))
	dataTx := GetDB()

	if orderBy != nil {
		for _, order := range *orderBy {
			// fmt.Println("order by:", order.Key, order.IsDESC)
			orderMod := "ASC"
			if order.IsDESC {
				orderMod = "DESC"
			}

			orderStr := fmt.Sprintf("%s %s", order.Key, orderMod)
			countTx = countTx.Order(orderStr)
			dataTx = dataTx.Order(orderStr)
		}
	}
	if whereClause != "" {
		countTx = countTx.Where(whereClause, whereValues...)
		dataTx = dataTx.Where(whereClause, whereValues...)
	}

	err = countTx.
		Count(&total).
		Error
	if err != nil {
		return
	}

	err = dataTx.
		Offset(offset).
		Limit(limit).
		Find(&data).
		Error

	return
}
