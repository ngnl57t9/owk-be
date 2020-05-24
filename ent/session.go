// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"owknight-be/ent/session"
	"owknight-be/ent/user"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// Session is the model entity for the Session schema.
type Session struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Token holds the value of the "token" field.
	Token string `json:"token,omitempty"`
	// ExpiredAt holds the value of the "expiredAt" field.
	ExpiredAt time.Time `json:"expiredAt,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SessionQuery when eager-loading is set.
	Edges        SessionEdges `json:"edges"`
	user_session *int
}

// SessionEdges holds the relations/edges for other nodes in the graph.
type SessionEdges struct {
	// User holds the value of the user edge.
	User *User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SessionEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Session) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // token
		&sql.NullTime{},   // expiredAt
		&sql.NullTime{},   // createdAt
		&sql.NullTime{},   // updatedAt
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Session) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // user_session
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Session fields.
func (s *Session) assignValues(values ...interface{}) error {
	if m, n := len(values), len(session.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field token", values[0])
	} else if value.Valid {
		s.Token = value.String
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field expiredAt", values[1])
	} else if value.Valid {
		s.ExpiredAt = value.Time
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field createdAt", values[2])
	} else if value.Valid {
		s.CreatedAt = value.Time
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updatedAt", values[3])
	} else if value.Valid {
		s.UpdatedAt = value.Time
	}
	values = values[4:]
	if len(values) == len(session.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_session", value)
		} else if value.Valid {
			s.user_session = new(int)
			*s.user_session = int(value.Int64)
		}
	}
	return nil
}

// QueryUser queries the user edge of the Session.
func (s *Session) QueryUser() *UserQuery {
	return (&SessionClient{config: s.config}).QueryUser(s)
}

// Update returns a builder for updating this Session.
// Note that, you need to call Session.Unwrap() before calling this method, if this Session
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Session) Update() *SessionUpdateOne {
	return (&SessionClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Session) Unwrap() *Session {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Session is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Session) String() string {
	var builder strings.Builder
	builder.WriteString("Session(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", token=")
	builder.WriteString(s.Token)
	builder.WriteString(", expiredAt=")
	builder.WriteString(s.ExpiredAt.Format(time.ANSIC))
	builder.WriteString(", createdAt=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Sessions is a parsable slice of Session.
type Sessions []*Session

func (s Sessions) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}