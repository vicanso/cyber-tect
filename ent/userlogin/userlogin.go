// Code generated by ent, DO NOT EDIT.

package userlogin

import (
	"time"
)

const (
	// Label holds the string label denoting the userlogin type in the database.
	Label = "user_login"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldAccount holds the string denoting the account field in the database.
	FieldAccount = "account"
	// FieldUserAgent holds the string denoting the user_agent field in the database.
	FieldUserAgent = "user_agent"
	// FieldIP holds the string denoting the ip field in the database.
	FieldIP = "ip"
	// FieldTrackID holds the string denoting the track_id field in the database.
	FieldTrackID = "track_id"
	// FieldSessionID holds the string denoting the session_id field in the database.
	FieldSessionID = "session_id"
	// FieldXForwardedFor holds the string denoting the x_forwarded_for field in the database.
	FieldXForwardedFor = "x_forwarded_for"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldProvince holds the string denoting the province field in the database.
	FieldProvince = "province"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldIsp holds the string denoting the isp field in the database.
	FieldIsp = "isp"
	// Table holds the table name of the userlogin in the database.
	Table = "user_logins"
)

// Columns holds all SQL columns for userlogin fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldAccount,
	FieldUserAgent,
	FieldIP,
	FieldTrackID,
	FieldSessionID,
	FieldXForwardedFor,
	FieldCountry,
	FieldProvince,
	FieldCity,
	FieldIsp,
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
	// AccountValidator is a validator for the "account" field. It is called by the builders before save.
	AccountValidator func(string) error
)
