// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/vicanso/cybertect/ent/httpdetectorresult"
	"github.com/vicanso/cybertect/schema"
)

// HTTPDetectorResult is the model entity for the HTTPDetectorResult schema.
type HTTPDetectorResult struct {
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
	// 检测URL
	URL string `json:"url,omitempty"`
	// 检测结果列表
	Results      schema.HTTPDetectorSubResults `json:"results,omitempty"`
	selectValues sql.SelectValues

	// 状态描述
	TaskName string `json:"taskName,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*HTTPDetectorResult) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case httpdetectorresult.FieldMessages, httpdetectorresult.FieldResults:
			values[i] = new([]byte)
		case httpdetectorresult.FieldID, httpdetectorresult.FieldTask, httpdetectorresult.FieldResult, httpdetectorresult.FieldMaxDuration:
			values[i] = new(sql.NullInt64)
		case httpdetectorresult.FieldURL:
			values[i] = new(sql.NullString)
		case httpdetectorresult.FieldCreatedAt, httpdetectorresult.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the HTTPDetectorResult fields.
func (hdr *HTTPDetectorResult) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case httpdetectorresult.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			hdr.ID = int(value.Int64)
		case httpdetectorresult.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				hdr.CreatedAt = value.Time
			}
		case httpdetectorresult.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				hdr.UpdatedAt = value.Time
			}
		case httpdetectorresult.FieldTask:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field task", values[i])
			} else if value.Valid {
				hdr.Task = int(value.Int64)
			}
		case httpdetectorresult.FieldResult:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field result", values[i])
			} else if value.Valid {
				hdr.Result = schema.DetectorResult(value.Int64)
			}
		case httpdetectorresult.FieldMaxDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field maxDuration", values[i])
			} else if value.Valid {
				hdr.MaxDuration = int(value.Int64)
			}
		case httpdetectorresult.FieldMessages:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field messages", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &hdr.Messages); err != nil {
					return fmt.Errorf("unmarshal field messages: %w", err)
				}
			}
		case httpdetectorresult.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				hdr.URL = value.String
			}
		case httpdetectorresult.FieldResults:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field results", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &hdr.Results); err != nil {
					return fmt.Errorf("unmarshal field results: %w", err)
				}
			}
		default:
			hdr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the HTTPDetectorResult.
// This includes values selected through modifiers, order, etc.
func (hdr *HTTPDetectorResult) Value(name string) (ent.Value, error) {
	return hdr.selectValues.Get(name)
}

// Update returns a builder for updating this HTTPDetectorResult.
// Note that you need to call HTTPDetectorResult.Unwrap() before calling this method if this HTTPDetectorResult
// was returned from a transaction, and the transaction was committed or rolled back.
func (hdr *HTTPDetectorResult) Update() *HTTPDetectorResultUpdateOne {
	return NewHTTPDetectorResultClient(hdr.config).UpdateOne(hdr)
}

// Unwrap unwraps the HTTPDetectorResult entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (hdr *HTTPDetectorResult) Unwrap() *HTTPDetectorResult {
	_tx, ok := hdr.config.driver.(*txDriver)
	if !ok {
		panic("ent: HTTPDetectorResult is not a transactional entity")
	}
	hdr.config.driver = _tx.drv
	return hdr
}

// String implements the fmt.Stringer.
func (hdr *HTTPDetectorResult) String() string {
	var builder strings.Builder
	builder.WriteString("HTTPDetectorResult(")
	builder.WriteString(fmt.Sprintf("id=%v, ", hdr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(hdr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(hdr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("task=")
	builder.WriteString(fmt.Sprintf("%v", hdr.Task))
	builder.WriteString(", ")
	builder.WriteString("result=")
	builder.WriteString(fmt.Sprintf("%v", hdr.Result))
	builder.WriteString(", ")
	builder.WriteString("maxDuration=")
	builder.WriteString(fmt.Sprintf("%v", hdr.MaxDuration))
	builder.WriteString(", ")
	builder.WriteString("messages=")
	builder.WriteString(fmt.Sprintf("%v", hdr.Messages))
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(hdr.URL)
	builder.WriteString(", ")
	builder.WriteString("results=")
	builder.WriteString(fmt.Sprintf("%v", hdr.Results))
	builder.WriteByte(')')
	return builder.String()
}

// HTTPDetectorResults is a parsable slice of HTTPDetectorResult.
type HTTPDetectorResults []*HTTPDetectorResult
