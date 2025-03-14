package gormx

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

	whereClause, whereValues, errx := where.Build()
	if errx != nil {
		return nil, 0, errx
	}

	dataTx := GetDB().Model(new(T))

	if orderBy != nil {
		for _, order := range *orderBy {
			dataTx = dataTx.Order(order.Clause())
		}
	}
	if whereClause != "" {
		dataTx = dataTx.Where(whereClause, whereValues...)
	}

	err = dataTx.
		Count(&total).
		Offset(offset).
		Limit(limit).
		Find(&data).
		Error

	return
}
