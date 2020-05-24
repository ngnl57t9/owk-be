// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"owknight-be/ent/profile"
	"owknight-be/ent/session"
	"owknight-be/ent/user"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// Salt holds the value of the "salt" field.
	Salt string `json:"-"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Session holds the value of the session edge.
	Session *Session `json:"-"`
	// Profile holds the value of the profile edge.
	Profile *Profile
	// Roles holds the value of the roles edge.
	Roles []*Role
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// SessionOrErr returns the Session value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) SessionOrErr() (*Session, error) {
	if e.loadedTypes[0] {
		if e.Session == nil {
			// The edge session was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: session.Label}
		}
		return e.Session, nil
	}
	return nil, &NotLoadedError{edge: "session"}
}

// ProfileOrErr returns the Profile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ProfileOrErr() (*Profile, error) {
	if e.loadedTypes[1] {
		if e.Profile == nil {
			// The edge profile was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: profile.Label}
		}
		return e.Profile, nil
	}
	return nil, &NotLoadedError{edge: "profile"}
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RolesOrErr() ([]*Role, error) {
	if e.loadedTypes[2] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // username
		&sql.NullString{}, // password
		&sql.NullString{}, // salt
		&sql.NullTime{},   // updatedAt
		&sql.NullTime{},   // createdAt
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		u.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field username", values[1])
	} else if value.Valid {
		u.Username = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field password", values[2])
	} else if value.Valid {
		u.Password = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field salt", values[3])
	} else if value.Valid {
		u.Salt = value.String
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updatedAt", values[4])
	} else if value.Valid {
		u.UpdatedAt = value.Time
	}
	if value, ok := values[5].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field createdAt", values[5])
	} else if value.Valid {
		u.CreatedAt = value.Time
	}
	return nil
}

// QuerySession queries the session edge of the User.
func (u *User) QuerySession() *SessionQuery {
	return (&UserClient{config: u.config}).QuerySession(u)
}

// QueryProfile queries the profile edge of the User.
func (u *User) QueryProfile() *ProfileQuery {
	return (&UserClient{config: u.config}).QueryProfile(u)
}

// QueryRoles queries the roles edge of the User.
func (u *User) QueryRoles() *RoleQuery {
	return (&UserClient{config: u.config}).QueryRoles(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", username=")
	builder.WriteString(u.Username)
	builder.WriteString(", password=")
	builder.WriteString(u.Password)
	builder.WriteString(", salt=")
	builder.WriteString(u.Salt)
	builder.WriteString(", updatedAt=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", createdAt=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
