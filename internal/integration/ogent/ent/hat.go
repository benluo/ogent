// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"ariga.io/ogent/internal/integration/ogent/ent/hat"
	"ariga.io/ogent/internal/integration/ogent/ent/user"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Hat is the model entity for the Hat schema.
type Hat struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type hat.Type `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HatQuery when eager-loading is set.
	Edges             HatEdges `json:"edges"`
	user_favorite_hat *int
	selectValues      sql.SelectValues
}

// HatEdges holds the relations/edges for other nodes in the graph.
type HatEdges struct {
	// Wearer holds the value of the wearer edge.
	Wearer *User `json:"wearer,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// WearerOrErr returns the Wearer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HatEdges) WearerOrErr() (*User, error) {
	if e.Wearer != nil {
		return e.Wearer, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "wearer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Hat) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case hat.FieldID:
			values[i] = new(sql.NullInt64)
		case hat.FieldName, hat.FieldType:
			values[i] = new(sql.NullString)
		case hat.ForeignKeys[0]: // user_favorite_hat
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Hat fields.
func (h *Hat) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hat.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			h.ID = int(value.Int64)
		case hat.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				h.Name = value.String
			}
		case hat.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				h.Type = hat.Type(value.String)
			}
		case hat.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_favorite_hat", value)
			} else if value.Valid {
				h.user_favorite_hat = new(int)
				*h.user_favorite_hat = int(value.Int64)
			}
		default:
			h.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Hat.
// This includes values selected through modifiers, order, etc.
func (h *Hat) Value(name string) (ent.Value, error) {
	return h.selectValues.Get(name)
}

// QueryWearer queries the "wearer" edge of the Hat entity.
func (h *Hat) QueryWearer() *UserQuery {
	return NewHatClient(h.config).QueryWearer(h)
}

// Update returns a builder for updating this Hat.
// Note that you need to call Hat.Unwrap() before calling this method if this Hat
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *Hat) Update() *HatUpdateOne {
	return NewHatClient(h.config).UpdateOne(h)
}

// Unwrap unwraps the Hat entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *Hat) Unwrap() *Hat {
	_tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: Hat is not a transactional entity")
	}
	h.config.driver = _tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *Hat) String() string {
	var builder strings.Builder
	builder.WriteString("Hat(")
	builder.WriteString(fmt.Sprintf("id=%v, ", h.ID))
	builder.WriteString("name=")
	builder.WriteString(h.Name)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", h.Type))
	builder.WriteByte(')')
	return builder.String()
}

// Hats is a parsable slice of Hat.
type Hats []*Hat