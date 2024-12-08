package gormx

import (
	"gorm.io/datatypes"
)

// JSON is a JSON.
type JSON = datatypes.JSON

// JSONObject is a JSON object.
type JSONObject = datatypes.JSONMap

// JSONArray is a JSON array.
// @TODO
type JSONArray[T any] datatypes.JSONSlice[T]

// UUID is a UUID.
type UUID = datatypes.UUID

// Date is a date.
type Date = datatypes.Date
