// Code generated by ent, DO NOT EDIT.

package databasedetectorresult

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the databasedetectorresult type in the database.
	Label = "database_detector_result"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldTask holds the string denoting the task field in the database.
	FieldTask = "task"
	// FieldResult holds the string denoting the result field in the database.
	FieldResult = "result"
	// FieldMaxDuration holds the string denoting the maxduration field in the database.
	FieldMaxDuration = "max_duration"
	// FieldMessages holds the string denoting the messages field in the database.
	FieldMessages = "messages"
	// FieldUris holds the string denoting the uris field in the database.
	FieldUris = "uris"
	// FieldResults holds the string denoting the results field in the database.
	FieldResults = "results"
	// Table holds the table name of the databasedetectorresult in the database.
	Table = "database_detector_results"
)

// Columns holds all SQL columns for databasedetectorresult fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldTask,
	FieldResult,
	FieldMaxDuration,
	FieldMessages,
	FieldUris,
	FieldResults,
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
	// ResultValidator is a validator for the "result" field. It is called by the builders before save.
	ResultValidator func(int8) error
)

// OrderOption defines the ordering options for the DatabaseDetectorResult queries.
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

// ByTask orders the results by the task field.
func ByTask(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTask, opts...).ToFunc()
}

// ByResult orders the results by the result field.
func ByResult(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldResult, opts...).ToFunc()
}

// ByMaxDuration orders the results by the maxDuration field.
func ByMaxDuration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMaxDuration, opts...).ToFunc()
}
