// Code generated by ent, DO NOT EDIT.

package pingdetector

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/vicanso/cybertect/schema"
)

const (
	// Label holds the string label denoting the pingdetector type in the database.
	Label = "ping_detector"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldOwners holds the string denoting the owners field in the database.
	FieldOwners = "owners"
	// FieldReceivers holds the string denoting the receivers field in the database.
	FieldReceivers = "receivers"
	// FieldTimeout holds the string denoting the timeout field in the database.
	FieldTimeout = "timeout"
	// FieldInterval holds the string denoting the interval field in the database.
	FieldInterval = "interval"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldIps holds the string denoting the ips field in the database.
	FieldIps = "ips"
	// Table holds the table name of the pingdetector in the database.
	Table = "ping_detectors"
)

// Columns holds all SQL columns for pingdetector fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldName,
	FieldOwners,
	FieldReceivers,
	FieldTimeout,
	FieldInterval,
	FieldDescription,
	FieldIps,
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

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus schema.Status
	// StatusValidator is a validator for the "status" field. It is called by the builders before save.
	StatusValidator func(int8) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// TimeoutValidator is a validator for the "timeout" field. It is called by the builders before save.
	TimeoutValidator func(string) error
)

// OrderOption defines the ordering options for the PingDetector queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByTimeout orders the results by the timeout field.
func ByTimeout(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimeout, opts...).ToFunc()
}

// ByInterval orders the results by the interval field.
func ByInterval(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInterval, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}