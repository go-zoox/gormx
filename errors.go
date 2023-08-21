package gormx

import (
	"errors"

	"gorm.io/gorm"
)

var (
	// ErrRecordNotFound record not found error
	ErrRecordNotFound = gorm.ErrRecordNotFound
	// ErrInvalidTransaction invalid transaction when you are trying to `Commit` or `Rollback`
	ErrInvalidTransaction = gorm.ErrInvalidTransaction
	// ErrNotImplemented not implemented
	ErrNotImplemented = gorm.ErrNotImplemented
	// ErrMissingWhereClause missing where clause
	ErrMissingWhereClause = gorm.ErrMissingWhereClause
	// ErrUnsupportedRelation unsupported relations
	ErrUnsupportedRelation = gorm.ErrUnsupportedRelation
	// ErrPrimaryKeyRequired primary keys required
	ErrPrimaryKeyRequired = gorm.ErrPrimaryKeyRequired
	// ErrModelValueRequired model value required
	ErrModelValueRequired = gorm.ErrModelValueRequired
	// ErrModelAccessibleFieldsRequired model accessible fields required
	ErrModelAccessibleFieldsRequired = gorm.ErrModelAccessibleFieldsRequired
	// ErrSubQueryRequired sub query required
	ErrSubQueryRequired = gorm.ErrSubQueryRequired
	// ErrInvalidData unsupported data
	ErrInvalidData = gorm.ErrInvalidData
	// ErrUnsupportedDriver unsupported driver
	ErrUnsupportedDriver = gorm.ErrUnsupportedDriver
	// ErrRegistered registered
	ErrRegistered = gorm.ErrRegistered
	// ErrInvalidField invalid field
	ErrInvalidField = gorm.ErrInvalidField
	// ErrEmptySlice empty slice found
	ErrEmptySlice = gorm.ErrEmptySlice
	// ErrDryRunModeUnsupported dry run mode unsupported
	ErrDryRunModeUnsupported = gorm.ErrDryRunModeUnsupported
	// ErrInvalidDB invalid db
	ErrInvalidDB = gorm.ErrInvalidDB
	// ErrInvalidValue invalid value
	ErrInvalidValue = gorm.ErrInvalidValue
	// ErrInvalidValueOfLength invalid values do not match length
	ErrInvalidValueOfLength = gorm.ErrInvalidValueOfLength
	// ErrPreloadNotAllowed preload is not allowed when count is used
	ErrPreloadNotAllowed = gorm.ErrPreloadNotAllowed
	// ErrDuplicatedKey occurs when there is a unique key constraint violation
	ErrDuplicatedKey = gorm.ErrDuplicatedKey
	// ErrForeignKeyViolated occurs when there is a foreign key constraint violation
	ErrForeignKeyViolated = gorm.ErrForeignKeyViolated
)

// IsRecordNotFoundError returns true if err is related to record not found error
func IsRecordNotFoundError(err error) bool {
	return errors.Is(err, ErrRecordNotFound)
}

// IsInvalidTransactionError returns true if err is related to invalid transaction error
func IsInvalidTransactionError(err error) bool {
	return errors.Is(err, ErrInvalidTransaction)
}

// IsNotImplementedError returns true if err is related to not implemented error
func IsNotImplementedError(err error) bool {
	return errors.Is(err, ErrNotImplemented)
}

// IsMissingWhereClauseError returns true if err is related to missing where clause error
func IsMissingWhereClauseError(err error) bool {
	return errors.Is(err, ErrMissingWhereClause)
}

// IsUnsupportedRelationError returns true if err is related to unsupported relation error
func IsUnsupportedRelationError(err error) bool {
	return errors.Is(err, ErrUnsupportedRelation)
}

// IsPrimaryKeyRequiredError returns true if err is related to primary key required error
func IsPrimaryKeyRequiredError(err error) bool {
	return errors.Is(err, ErrPrimaryKeyRequired)
}

// IsModelValueRequiredError returns true if err is related to model value required error
func IsModelValueRequiredError(err error) bool {
	return errors.Is(err, ErrModelValueRequired)
}

// IsModelAccessibleFieldsRequiredError returns true if err is related to model accessible fields required error
func IsModelAccessibleFieldsRequiredError(err error) bool {
	return errors.Is(err, ErrModelAccessibleFieldsRequired)
}

// IsSubQueryRequiredError returns true if err is related to sub query required error
func IsSubQueryRequiredError(err error) bool {
	return errors.Is(err, ErrSubQueryRequired)
}

// IsInvalidDataError returns true if err is related to invalid data error
func IsInvalidDataError(err error) bool {
	return errors.Is(err, ErrInvalidData)
}

// IsUnsupportedDriverError returns true if err is related to unsupported driver error
func IsUnsupportedDriverError(err error) bool {
	return errors.Is(err, ErrUnsupportedDriver)
}

// IsRegisteredError returns true if err is related to registered error
func IsRegisteredError(err error) bool {
	return errors.Is(err, ErrRegistered)
}

// IsInvalidFieldError returns true if err is related to invalid field error
func IsInvalidFieldError(err error) bool {
	return errors.Is(err, ErrInvalidField)
}

// IsEmptySliceError returns true if err is related to empty slice error
func IsEmptySliceError(err error) bool {
	return errors.Is(err, ErrEmptySlice)
}

// IsDryRunModeUnsupportedError returns true if err is related to dry run mode unsupported error
func IsDryRunModeUnsupportedError(err error) bool {
	return errors.Is(err, ErrDryRunModeUnsupported)
}

// IsInvalidDBError returns true if err is related to invalid db error
func IsInvalidDBError(err error) bool {
	return errors.Is(err, ErrInvalidDB)
}

// IsInvalidValueError returns true if err is related to invalid value error
func IsInvalidValueError(err error) bool {
	return errors.Is(err, ErrInvalidValue)
}

// IsInvalidValueOfLengthError returns true if err is related to invalid value of length error
func IsInvalidValueOfLengthError(err error) bool {
	return errors.Is(err, ErrInvalidValueOfLength)
}

// IsPreloadNotAllowedError returns true if err is related to preload not allowed error
func IsPreloadNotAllowedError(err error) bool {
	return errors.Is(err, ErrPreloadNotAllowed)
}

// IsDuplicatedKeyError returns true if err is related to duplicated key error
func IsDuplicatedKeyError(err error) bool {
	return errors.Is(err, ErrDuplicatedKey)
}

// IsForeignKeyViolatedError returns true if err is related to foreign key violated error
func IsForeignKeyViolatedError(err error) bool {
	return errors.Is(err, ErrForeignKeyViolated)
}
