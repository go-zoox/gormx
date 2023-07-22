package gormx

// Page is the page query.
type Page struct {
	Page     int64 `query:"page,default=1"`
	PageSize int64 `query:"pageSize,default=10"`
}
