package gormx

import "fmt"

// ListAll lists all records.
func ListAll[T any](where *Where, orderBy *OrderBy) (data []*T, err error) {
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

	err = dataTx.
		Find(&data).
		Error

	return
}
