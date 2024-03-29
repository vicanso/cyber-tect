// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/vicanso/cybertect/ent/pingdetectorresult"
	"github.com/vicanso/cybertect/schema"
)

// PingDetectorResult is the model entity for the PingDetectorResult schema.
type PingDetectorResult struct {
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
	// 检测IP
	Ips []string `json:"ips,omitempty"`
	// 检测结果列表
	Results      schema.PingDetectorSubResults `json:"results,omitempty"`
	selectValues sql.SelectValues

	// 状态描述
	TaskName string `json:"taskName,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PingDetectorResult) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case pingdetectorresult.FieldMessages, pingdetectorresult.FieldIps, pingdetectorresult.FieldResults:
			values[i] = new([]byte)
		case pingdetectorresult.FieldID, pingdetectorresult.FieldTask, pingdetectorresult.FieldResult, pingdetectorresult.FieldMaxDuration:
			values[i] = new(sql.NullInt64)
		case pingdetectorresult.FieldCreatedAt, pingdetectorresult.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PingDetectorResult fields.
func (pdr *PingDetectorResult) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pingdetectorresult.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pdr.ID = int(value.Int64)
		case pingdetectorresult.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pdr.CreatedAt = value.Time
			}
		case pingdetectorresult.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pdr.UpdatedAt = value.Time
			}
		case pingdetectorresult.FieldTask:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field task", values[i])
			} else if value.Valid {
				pdr.Task = int(value.Int64)
			}
		case pingdetectorresult.FieldResult:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field result", values[i])
			} else if value.Valid {
				pdr.Result = schema.DetectorResult(value.Int64)
			}
		case pingdetectorresult.FieldMaxDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field maxDuration", values[i])
			} else if value.Valid {
				pdr.MaxDuration = int(value.Int64)
			}
		case pingdetectorresult.FieldMessages:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field messages", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pdr.Messages); err != nil {
					return fmt.Errorf("unmarshal field messages: %w", err)
				}
			}
		case pingdetectorresult.FieldIps:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field ips", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pdr.Ips); err != nil {
					return fmt.Errorf("unmarshal field ips: %w", err)
				}
			}
		case pingdetectorresult.FieldResults:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field results", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pdr.Results); err != nil {
					return fmt.Errorf("unmarshal field results: %w", err)
				}
			}
		default:
			pdr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PingDetectorResult.
// This includes values selected through modifiers, order, etc.
func (pdr *PingDetectorResult) Value(name string) (ent.Value, error) {
	return pdr.selectValues.Get(name)
}

// Update returns a builder for updating this PingDetectorResult.
// Note that you need to call PingDetectorResult.Unwrap() before calling this method if this PingDetectorResult
// was returned from a transaction, and the transaction was committed or rolled back.
func (pdr *PingDetectorResult) Update() *PingDetectorResultUpdateOne {
	return NewPingDetectorResultClient(pdr.config).UpdateOne(pdr)
}

// Unwrap unwraps the PingDetectorResult entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pdr *PingDetectorResult) Unwrap() *PingDetectorResult {
	_tx, ok := pdr.config.driver.(*txDriver)
	if !ok {
		panic("ent: PingDetectorResult is not a transactional entity")
	}
	pdr.config.driver = _tx.drv
	return pdr
}

// String implements the fmt.Stringer.
func (pdr *PingDetectorResult) String() string {
	var builder strings.Builder
	builder.WriteString("PingDetectorResult(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pdr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pdr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pdr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("task=")
	builder.WriteString(fmt.Sprintf("%v", pdr.Task))
	builder.WriteString(", ")
	builder.WriteString("result=")
	builder.WriteString(fmt.Sprintf("%v", pdr.Result))
	builder.WriteString(", ")
	builder.WriteString("maxDuration=")
	builder.WriteString(fmt.Sprintf("%v", pdr.MaxDuration))
	builder.WriteString(", ")
	builder.WriteString("messages=")
	builder.WriteString(fmt.Sprintf("%v", pdr.Messages))
	builder.WriteString(", ")
	builder.WriteString("ips=")
	builder.WriteString(fmt.Sprintf("%v", pdr.Ips))
	builder.WriteString(", ")
	builder.WriteString("results=")
	builder.WriteString(fmt.Sprintf("%v", pdr.Results))
	builder.WriteByte(')')
	return builder.String()
}

// PingDetectorResults is a parsable slice of PingDetectorResult.
type PingDetectorResults []*PingDetectorResult
