// Code generated by ent, DO NOT EDIT.

package alltypes

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the alltypes type in the database.
	Label = "all_types"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldInt holds the string denoting the int field in the database.
	FieldInt = "int"
	// FieldInt8 holds the string denoting the int8 field in the database.
	FieldInt8 = "int8"
	// FieldInt16 holds the string denoting the int16 field in the database.
	FieldInt16 = "int16"
	// FieldInt32 holds the string denoting the int32 field in the database.
	FieldInt32 = "int32"
	// FieldInt64 holds the string denoting the int64 field in the database.
	FieldInt64 = "int64"
	// FieldUint holds the string denoting the uint field in the database.
	FieldUint = "uint"
	// FieldUint8 holds the string denoting the uint8 field in the database.
	FieldUint8 = "uint8"
	// FieldUint16 holds the string denoting the uint16 field in the database.
	FieldUint16 = "uint16"
	// FieldUint32 holds the string denoting the uint32 field in the database.
	FieldUint32 = "uint32"
	// FieldUint64 holds the string denoting the uint64 field in the database.
	FieldUint64 = "uint64"
	// FieldFloat32 holds the string denoting the float32 field in the database.
	FieldFloat32 = "float32"
	// FieldFloat64 holds the string denoting the float64 field in the database.
	FieldFloat64 = "float64"
	// FieldStringType holds the string denoting the string_type field in the database.
	FieldStringType = "string_type"
	// FieldBool holds the string denoting the bool field in the database.
	FieldBool = "bool"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldBytes holds the string denoting the bytes field in the database.
	FieldBytes = "bytes"
	// FieldNilable holds the string denoting the nilable field in the database.
	FieldNilable = "nilable"
	// Table holds the table name of the alltypes in the database.
	Table = "all_types"
)

// Columns holds all SQL columns for alltypes fields.
var Columns = []string{
	FieldID,
	FieldInt,
	FieldInt8,
	FieldInt16,
	FieldInt32,
	FieldInt64,
	FieldUint,
	FieldUint8,
	FieldUint16,
	FieldUint32,
	FieldUint64,
	FieldFloat32,
	FieldFloat64,
	FieldStringType,
	FieldBool,
	FieldUUID,
	FieldTime,
	FieldText,
	FieldState,
	FieldBytes,
	FieldNilable,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// State defines the type for the "state" enum field.
type State string

// State values.
const (
	StateOn  State = "on"
	StateOff State = "off"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(s State) error {
	switch s {
	case StateOn, StateOff:
		return nil
	default:
		return fmt.Errorf("alltypes: invalid enum value for state field: %q", s)
	}
}

// OrderOption defines the ordering options for the AllTypes queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByInt orders the results by the int field.
func ByInt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInt, opts...).ToFunc()
}

// ByInt8 orders the results by the int8 field.
func ByInt8(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInt8, opts...).ToFunc()
}

// ByInt16 orders the results by the int16 field.
func ByInt16(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInt16, opts...).ToFunc()
}

// ByInt32 orders the results by the int32 field.
func ByInt32(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInt32, opts...).ToFunc()
}

// ByInt64 orders the results by the int64 field.
func ByInt64(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInt64, opts...).ToFunc()
}

// ByUint orders the results by the uint field.
func ByUint(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUint, opts...).ToFunc()
}

// ByUint8 orders the results by the uint8 field.
func ByUint8(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUint8, opts...).ToFunc()
}

// ByUint16 orders the results by the uint16 field.
func ByUint16(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUint16, opts...).ToFunc()
}

// ByUint32 orders the results by the uint32 field.
func ByUint32(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUint32, opts...).ToFunc()
}

// ByUint64 orders the results by the uint64 field.
func ByUint64(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUint64, opts...).ToFunc()
}

// ByFloat32 orders the results by the float32 field.
func ByFloat32(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFloat32, opts...).ToFunc()
}

// ByFloat64 orders the results by the float64 field.
func ByFloat64(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFloat64, opts...).ToFunc()
}

// ByStringType orders the results by the string_type field.
func ByStringType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStringType, opts...).ToFunc()
}

// ByBool orders the results by the bool field.
func ByBool(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBool, opts...).ToFunc()
}

// ByUUID orders the results by the uuid field.
func ByUUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUUID, opts...).ToFunc()
}

// ByTime orders the results by the time field.
func ByTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTime, opts...).ToFunc()
}

// ByText orders the results by the text field.
func ByText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldText, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// ByNilable orders the results by the nilable field.
func ByNilable(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNilable, opts...).ToFunc()
}
