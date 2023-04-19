// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/vicanso/cybertect/ent/dnsdetectorresult"
	"github.com/vicanso/cybertect/schema"
)

// DNSDetectorResult is the model entity for the DNSDetectorResult schema.
type DNSDetectorResult struct {
	config `json:"-" sql:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 创建时间，添加记录时由程序自动生成
	CreatedAt time.Time `json:"createdAt" sql:"created_at"`
	// 更新时间，更新记录时由程序自动生成
	UpdatedAt time.Time `json:"updatedAt" sql:"updated_at"`
	// 任务ID
	Task int `json:"task,omitempty"`
	// 检测结果
	Result schema.DetectorResult `json:"result,omitempty"`
	// 最长时长
	MaxDuration int `json:"maxDuration,omitempty" sql:"max_duration"`
	// 出错信息汇总
	Messages []string `json:"messages,omitempty"`
	// 检测Host
	Host string `json:"host,omitempty"`
	// 检测结果列表
	Results      schema.DNSDetectorSubResults `json:"results,omitempty"`
	selectValues sql.SelectValues

	// 状态描述
	TaskName string `json:"taskName,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DNSDetectorResult) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case dnsdetectorresult.FieldMessages, dnsdetectorresult.FieldResults:
			values[i] = new([]byte)
		case dnsdetectorresult.FieldID, dnsdetectorresult.FieldTask, dnsdetectorresult.FieldResult, dnsdetectorresult.FieldMaxDuration:
			values[i] = new(sql.NullInt64)
		case dnsdetectorresult.FieldHost:
			values[i] = new(sql.NullString)
		case dnsdetectorresult.FieldCreatedAt, dnsdetectorresult.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DNSDetectorResult fields.
func (ddr *DNSDetectorResult) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dnsdetectorresult.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ddr.ID = int(value.Int64)
		case dnsdetectorresult.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ddr.CreatedAt = value.Time
			}
		case dnsdetectorresult.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ddr.UpdatedAt = value.Time
			}
		case dnsdetectorresult.FieldTask:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field task", values[i])
			} else if value.Valid {
				ddr.Task = int(value.Int64)
			}
		case dnsdetectorresult.FieldResult:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field result", values[i])
			} else if value.Valid {
				ddr.Result = schema.DetectorResult(value.Int64)
			}
		case dnsdetectorresult.FieldMaxDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field maxDuration", values[i])
			} else if value.Valid {
				ddr.MaxDuration = int(value.Int64)
			}
		case dnsdetectorresult.FieldMessages:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field messages", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ddr.Messages); err != nil {
					return fmt.Errorf("unmarshal field messages: %w", err)
				}
			}
		case dnsdetectorresult.FieldHost:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field host", values[i])
			} else if value.Valid {
				ddr.Host = value.String
			}
		case dnsdetectorresult.FieldResults:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field results", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ddr.Results); err != nil {
					return fmt.Errorf("unmarshal field results: %w", err)
				}
			}
		default:
			ddr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DNSDetectorResult.
// This includes values selected through modifiers, order, etc.
func (ddr *DNSDetectorResult) Value(name string) (ent.Value, error) {
	return ddr.selectValues.Get(name)
}

// Update returns a builder for updating this DNSDetectorResult.
// Note that you need to call DNSDetectorResult.Unwrap() before calling this method if this DNSDetectorResult
// was returned from a transaction, and the transaction was committed or rolled back.
func (ddr *DNSDetectorResult) Update() *DNSDetectorResultUpdateOne {
	return NewDNSDetectorResultClient(ddr.config).UpdateOne(ddr)
}

// Unwrap unwraps the DNSDetectorResult entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ddr *DNSDetectorResult) Unwrap() *DNSDetectorResult {
	_tx, ok := ddr.config.driver.(*txDriver)
	if !ok {
		panic("ent: DNSDetectorResult is not a transactional entity")
	}
	ddr.config.driver = _tx.drv
	return ddr
}

// String implements the fmt.Stringer.
func (ddr *DNSDetectorResult) String() string {
	var builder strings.Builder
	builder.WriteString("DNSDetectorResult(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ddr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ddr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ddr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("task=")
	builder.WriteString(fmt.Sprintf("%v", ddr.Task))
	builder.WriteString(", ")
	builder.WriteString("result=")
	builder.WriteString(fmt.Sprintf("%v", ddr.Result))
	builder.WriteString(", ")
	builder.WriteString("maxDuration=")
	builder.WriteString(fmt.Sprintf("%v", ddr.MaxDuration))
	builder.WriteString(", ")
	builder.WriteString("messages=")
	builder.WriteString(fmt.Sprintf("%v", ddr.Messages))
	builder.WriteString(", ")
	builder.WriteString("host=")
	builder.WriteString(ddr.Host)
	builder.WriteString(", ")
	builder.WriteString("results=")
	builder.WriteString(fmt.Sprintf("%v", ddr.Results))
	builder.WriteByte(')')
	return builder.String()
}

// DNSDetectorResults is a parsable slice of DNSDetectorResult.
type DNSDetectorResults []*DNSDetectorResult