// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/vicanso/cybertect/ent/pingdetector"
	"github.com/vicanso/cybertect/schema"
)

// PingDetector is the model entity for the PingDetector schema.
type PingDetector struct {
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
	// 配置拥有者
	Owners []string `json:"owners,omitempty"`
	// 接收者列表
	Receivers []string `json:"receivers,omitempty"`
	// 超时设置
	Timeout string `json:"timeout,omitempty"`
	// 检测间隔
	Interval string `json:"interval,omitempty"`
	// 配置描述
	Description string `json:"description,omitempty"`
	// 检测IP列表
	Ips          []string `json:"ips,omitempty"`
	selectValues sql.SelectValues

	// 状态描述
	StatusDesc string `json:"statusDesc,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PingDetector) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case pingdetector.FieldOwners, pingdetector.FieldReceivers, pingdetector.FieldIps:
			values[i] = new([]byte)
		case pingdetector.FieldID, pingdetector.FieldStatus:
			values[i] = new(sql.NullInt64)
		case pingdetector.FieldName, pingdetector.FieldTimeout, pingdetector.FieldInterval, pingdetector.FieldDescription:
			values[i] = new(sql.NullString)
		case pingdetector.FieldCreatedAt, pingdetector.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PingDetector fields.
func (pd *PingDetector) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pingdetector.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pd.ID = int(value.Int64)
		case pingdetector.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pd.CreatedAt = value.Time
			}
		case pingdetector.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pd.UpdatedAt = value.Time
			}
		case pingdetector.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pd.Status = schema.Status(value.Int64)
			}
		case pingdetector.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pd.Name = value.String
			}
		case pingdetector.FieldOwners:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field owners", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pd.Owners); err != nil {
					return fmt.Errorf("unmarshal field owners: %w", err)
				}
			}
		case pingdetector.FieldReceivers:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field receivers", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pd.Receivers); err != nil {
					return fmt.Errorf("unmarshal field receivers: %w", err)
				}
			}
		case pingdetector.FieldTimeout:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field timeout", values[i])
			} else if value.Valid {
				pd.Timeout = value.String
			}
		case pingdetector.FieldInterval:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field interval", values[i])
			} else if value.Valid {
				pd.Interval = value.String
			}
		case pingdetector.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pd.Description = value.String
			}
		case pingdetector.FieldIps:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field ips", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pd.Ips); err != nil {
					return fmt.Errorf("unmarshal field ips: %w", err)
				}
			}
		default:
			pd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PingDetector.
// This includes values selected through modifiers, order, etc.
func (pd *PingDetector) Value(name string) (ent.Value, error) {
	return pd.selectValues.Get(name)
}

// Update returns a builder for updating this PingDetector.
// Note that you need to call PingDetector.Unwrap() before calling this method if this PingDetector
// was returned from a transaction, and the transaction was committed or rolled back.
func (pd *PingDetector) Update() *PingDetectorUpdateOne {
	return NewPingDetectorClient(pd.config).UpdateOne(pd)
}

// Unwrap unwraps the PingDetector entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pd *PingDetector) Unwrap() *PingDetector {
	_tx, ok := pd.config.driver.(*txDriver)
	if !ok {
		panic("ent: PingDetector is not a transactional entity")
	}
	pd.config.driver = _tx.drv
	return pd
}

// String implements the fmt.Stringer.
func (pd *PingDetector) String() string {
	var builder strings.Builder
	builder.WriteString("PingDetector(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pd.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pd.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pd.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pd.Status))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pd.Name)
	builder.WriteString(", ")
	builder.WriteString("owners=")
	builder.WriteString(fmt.Sprintf("%v", pd.Owners))
	builder.WriteString(", ")
	builder.WriteString("receivers=")
	builder.WriteString(fmt.Sprintf("%v", pd.Receivers))
	builder.WriteString(", ")
	builder.WriteString("timeout=")
	builder.WriteString(pd.Timeout)
	builder.WriteString(", ")
	builder.WriteString("interval=")
	builder.WriteString(pd.Interval)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pd.Description)
	builder.WriteString(", ")
	builder.WriteString("ips=")
	builder.WriteString(fmt.Sprintf("%v", pd.Ips))
	builder.WriteByte(')')
	return builder.String()
}

// PingDetectors is a parsable slice of PingDetector.
type PingDetectors []*PingDetector