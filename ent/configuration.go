// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/vicanso/cybertect/ent/configuration"
	"github.com/vicanso/cybertect/ent/schema"
)

// Configuration is the model entity for the Configuration schema.
type Configuration struct {
	config `json:"-" sql:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt,omitempty" sql:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty" sql:"updated_at"`
	// Status holds the value of the "status" field.
	Status schema.Status `json:"status,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Category holds the value of the "category" field.
	Category configuration.Category `json:"category,omitempty"`
	// Owner holds the value of the "owner" field.
	Owner string `json:"owner,omitempty"`
	// Data holds the value of the "data" field.
	Data string `json:"data,omitempty"`
	// StartedAt holds the value of the "started_at" field.
	StartedAt time.Time `json:"startedAt,omitempty"`
	// EndedAt holds the value of the "ended_at" field.
	EndedAt time.Time `json:"endedAt,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Configuration) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case configuration.FieldID, configuration.FieldStatus:
			values[i] = &sql.NullInt64{}
		case configuration.FieldName, configuration.FieldCategory, configuration.FieldOwner, configuration.FieldData:
			values[i] = &sql.NullString{}
		case configuration.FieldCreatedAt, configuration.FieldUpdatedAt, configuration.FieldStartedAt, configuration.FieldEndedAt:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Configuration", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Configuration fields.
func (c *Configuration) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case configuration.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case configuration.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case configuration.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case configuration.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				c.Status = schema.Status(value.Int64)
			}
		case configuration.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case configuration.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				c.Category = configuration.Category(value.String)
			}
		case configuration.FieldOwner:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner", values[i])
			} else if value.Valid {
				c.Owner = value.String
			}
		case configuration.FieldData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value.Valid {
				c.Data = value.String
			}
		case configuration.FieldStartedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field started_at", values[i])
			} else if value.Valid {
				c.StartedAt = value.Time
			}
		case configuration.FieldEndedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ended_at", values[i])
			} else if value.Valid {
				c.EndedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Configuration.
// Note that you need to call Configuration.Unwrap() before calling this method if this Configuration
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Configuration) Update() *ConfigurationUpdateOne {
	return (&ConfigurationClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Configuration entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Configuration) Unwrap() *Configuration {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Configuration is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Configuration) String() string {
	var builder strings.Builder
	builder.WriteString("Configuration(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", c.Status))
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteString(", category=")
	builder.WriteString(fmt.Sprintf("%v", c.Category))
	builder.WriteString(", owner=")
	builder.WriteString(c.Owner)
	builder.WriteString(", data=")
	builder.WriteString(c.Data)
	builder.WriteString(", started_at=")
	builder.WriteString(c.StartedAt.Format(time.ANSIC))
	builder.WriteString(", ended_at=")
	builder.WriteString(c.EndedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Configurations is a parsable slice of Configuration.
type Configurations []*Configuration

func (c Configurations) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
