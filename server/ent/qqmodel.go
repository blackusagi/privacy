// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/kallydev/privacy/ent/qqmodel"
)

// QQModel is the model entity for the QQModel schema.
type QQModel struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// QqNumber holds the value of the "qq_number" field.
	QqNumber int64 `json:"qq_number,omitempty"`
	// PhoneNumber holds the value of the "phone_number" field.
	PhoneNumber int64 `json:"phone_number,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*QQModel) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullInt64{}, // qq_number
		&sql.NullInt64{}, // phone_number
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the QQModel fields.
func (qm *QQModel) assignValues(values ...interface{}) error {
	if m, n := len(values), len(qqmodel.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	qm.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field qq_number", values[0])
	} else if value.Valid {
		qm.QqNumber = value.Int64
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field phone_number", values[1])
	} else if value.Valid {
		qm.PhoneNumber = value.Int64
	}
	return nil
}

// Update returns a builder for updating this QQModel.
// Note that, you need to call QQModel.Unwrap() before calling this method, if this QQModel
// was returned from a transaction, and the transaction was committed or rolled back.
func (qm *QQModel) Update() *QQModelUpdateOne {
	return (&QQModelClient{config: qm.config}).UpdateOne(qm)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (qm *QQModel) Unwrap() *QQModel {
	tx, ok := qm.config.driver.(*txDriver)
	if !ok {
		panic("ent: QQModel is not a transactional entity")
	}
	qm.config.driver = tx.drv
	return qm
}

// String implements the fmt.Stringer.
func (qm *QQModel) String() string {
	var builder strings.Builder
	builder.WriteString("QQModel(")
	builder.WriteString(fmt.Sprintf("id=%v", qm.ID))
	builder.WriteString(", qq_number=")
	builder.WriteString(fmt.Sprintf("%v", qm.QqNumber))
	builder.WriteString(", phone_number=")
	builder.WriteString(fmt.Sprintf("%v", qm.PhoneNumber))
	builder.WriteByte(')')
	return builder.String()
}

// QQModels is a parsable slice of QQModel.
type QQModels []*QQModel

func (qm QQModels) config(cfg config) {
	for _i := range qm {
		qm[_i].config = cfg
	}
}
