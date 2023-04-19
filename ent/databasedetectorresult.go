// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/vicanso/cybertect/ent/databasedetectorresult"
	"github.com/vicanso/cybertect/schema"
)

// DatabaseDetectorResult is the model entity for the DatabaseDetectorResult schema.
type DatabaseDetectorResult struct {
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
	// 检测的redis连接地址
	Uris []string `json:"uris,omitempty"`
	// 检测结果列表
	Results      schema.DatabaseDetectorSubResults `json:"results,omitempty"`
	selectValues sql.SelectValues

	// 状态描述
	TaskName string `json:"taskName,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DatabaseDetectorResult) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case databasedetectorresult.FieldMessages, databasedetectorresult.FieldUris, databasedetectorresult.FieldResults:
			values[i] = new([]byte)
		case databasedetectorresult.FieldID, databasedetectorresult.FieldTask, databasedetectorresult.FieldResult, databasedetectorresult.FieldMaxDuration:
			values[i] = new(sql.NullInt64)
		case databasedetectorresult.FieldCreatedAt, databasedetectorresult.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DatabaseDetectorResult fields.
func (ddr *DatabaseDetectorResult) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case databasedetectorresult.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ddr.ID = int(value.Int64)
		case databasedetectorresult.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ddr.CreatedAt = value.Time
			}
		case databasedetectorresult.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ddr.UpdatedAt = value.Time
			}
		case databasedetectorresult.FieldTask:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field task", values[i])
			} else if value.Valid {
				ddr.Task = int(value.Int64)
			}
		case databasedetectorresult.FieldResult:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field result", values[i])
			} else if value.Valid {
				ddr.Result = schema.DetectorResult(value.Int64)
			}
		case databasedetectorresult.FieldMaxDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field maxDuration", values[i])
			} else if value.Valid {
				ddr.MaxDuration = int(value.Int64)
			}
		case databasedetectorresult.FieldMessages:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field messages", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ddr.Messages); err != nil {
					return fmt.Errorf("unmarshal field messages: %w", err)
				}
			}
		case databasedetectorresult.FieldUris:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field uris", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ddr.Uris); err != nil {
					return fmt.Errorf("unmarshal field uris: %w", err)
				}
			}
		case databasedetectorresult.FieldResults:
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

// Value returns the ent.Value that was dynamically selected and assigned to the DatabaseDetectorResult.
// This includes values selected through modifiers, order, etc.
func (ddr *DatabaseDetectorResult) Value(name string) (ent.Value, error) {
	return ddr.selectValues.Get(name)
}

// Update returns a builder for updating this DatabaseDetectorResult.
// Note that you need to call DatabaseDetectorResult.Unwrap() before calling this method if this DatabaseDetectorResult
// was returned from a transaction, and the transaction was committed or rolled back.
func (ddr *DatabaseDetectorResult) Update() *DatabaseDetectorResultUpdateOne {
	return NewDatabaseDetectorResultClient(ddr.config).UpdateOne(ddr)
}

// Unwrap unwraps the DatabaseDetectorResult entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ddr *DatabaseDetectorResult) Unwrap() *DatabaseDetectorResult {
	_tx, ok := ddr.config.driver.(*txDriver)
	if !ok {
		panic("ent: DatabaseDetectorResult is not a transactional entity")
	}
	ddr.config.driver = _tx.drv
	return ddr
}

// String implements the fmt.Stringer.
func (ddr *DatabaseDetectorResult) String() string {
	var builder strings.Builder
	builder.WriteString("DatabaseDetectorResult(")
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
	builder.WriteString("uris=")
	builder.WriteString(fmt.Sprintf("%v", ddr.Uris))
	builder.WriteString(", ")
	builder.WriteString("results=")
	builder.WriteString(fmt.Sprintf("%v", ddr.Results))
	builder.WriteByte(')')
	return builder.String()
}

// DatabaseDetectorResults is a parsable slice of DatabaseDetectorResult.
type DatabaseDetectorResults []*DatabaseDetectorResult
