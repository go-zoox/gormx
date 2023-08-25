package gormx

import "fmt"

// ListALL lists all records.
func ListALL[T any](where *Where, orderBy *OrderBy) (data []*T, err error) {
	countTx := GetDB().Model(new(T))
	dataTx := GetDB()

	if where != nil {
		whereClause, whereValues := where.Build()
		if whereClause != "" {
			countTx = countTx.Where(whereClause, whereValues...)
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
			countTx = countTx.Order(orderStr)
			dataTx = dataTx.Order(orderStr)
		}
	}

	err = dataTx.
		Find(&data).
		Error

	return
}
