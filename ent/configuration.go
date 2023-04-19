// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/vicanso/cybertect/ent/configuration"
	"github.com/vicanso/cybertect/schema"
)

// Configuration is the model entity for the Configuration schema.
type Configuration struct {
	config `json:"-" sql:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 创建时间，添加记录时由程序自动生成
	CreatedAt time.Time `json:"createdAt" sql:"created_at"`
	// 更新时间，更新记录时由程序自动生成
	UpdatedAt time.Time `json:"updatedAt" sql:"updated_at"`
	// 状态，默认为启用状态
	Status schema.Status `json:"status,omitempty"`
	// 配置名称
	Name string `json:"name,omitempty"`
	// 配置分类
	Category configuration.Category `json:"category,omitempty"`
	// 创建者
	Owner string `json:"owner,omitempty"`
	// 配置信息
	Data string `json:"data,omitempty"`
	// 配置启用时间
	StartedAt time.Time `json:"startedAt"`
	// 配置停用时间
	EndedAt time.Time `json:"endedAt"`
	// 配置说明
	Description  string `json:"description,omitempty"`
	selectValues sql.SelectValues

	// 状态描述
	StatusDesc string `json:"statusDesc,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Configuration) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case configuration.FieldID, configuration.FieldStatus:
			values[i] = new(sql.NullInt64)
		case configuration.FieldName, configuration.FieldCategory, configuration.FieldOwner, configuration.FieldData, configuration.FieldDescription:
			values[i] = new(sql.NullString)
		case configuration.FieldCreatedAt, configuration.FieldUpdatedAt, configuration.FieldStartedAt, configuration.FieldEndedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Configuration fields.
func (c *Configuration) assignValues(columns []string, values []any) error {
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
		case configuration.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Configuration.
// This includes values selected through modifiers, order, etc.
func (c *Configuration) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// Update returns a builder for updating this Configuration.
// Note that you need to call Configuration.Unwrap() before calling this method if this Configuration
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Configuration) Update() *ConfigurationUpdateOne {
	return NewConfigurationClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Configuration entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Configuration) Unwrap() *Configuration {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Configuration is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Configuration) String() string {
	var builder strings.Builder
	builder.WriteString("Configuration(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", c.Status))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(fmt.Sprintf("%v", c.Category))
	builder.WriteString(", ")
	builder.WriteString("owner=")
	builder.WriteString(c.Owner)
	builder.WriteString(", ")
	builder.WriteString("data=")
	builder.WriteString(c.Data)
	builder.WriteString(", ")
	builder.WriteString("started_at=")
	builder.WriteString(c.StartedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ended_at=")
	builder.WriteString(c.EndedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(c.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Configurations is a parsable slice of Configuration.
type Configurations []*Configuration