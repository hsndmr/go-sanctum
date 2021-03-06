// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/hsndmr/go-sanctum/ent/personalaccesstoken"
	"github.com/hsndmr/go-sanctum/ent/user"
)

// PersonalAccessToken is the model entity for the PersonalAccessToken schema.
type PersonalAccessToken struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// Token holds the value of the "token" field.
	Token string `json:"token,omitempty"`
	// Abilities holds the value of the "abilities" field.
	Abilities []string `json:"abilities,omitempty"`
	// ExpirationAt holds the value of the "expiration_at" field.
	ExpirationAt time.Time `json:"expiration_at,omitempty"`
	// LastUsedAt holds the value of the "last_used_at" field.
	LastUsedAt *time.Time `json:"last_used_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PersonalAccessTokenQuery when eager-loading is set.
	Edges PersonalAccessTokenEdges `json:"edges"`
}

// PersonalAccessTokenEdges holds the relations/edges for other nodes in the graph.
type PersonalAccessTokenEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PersonalAccessTokenEdges) UserOrErr() (*User, error) {
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
func (*PersonalAccessToken) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case personalaccesstoken.FieldAbilities:
			values[i] = new([]byte)
		case personalaccesstoken.FieldID, personalaccesstoken.FieldUserID:
			values[i] = new(sql.NullInt64)
		case personalaccesstoken.FieldName, personalaccesstoken.FieldToken:
			values[i] = new(sql.NullString)
		case personalaccesstoken.FieldExpirationAt, personalaccesstoken.FieldLastUsedAt, personalaccesstoken.FieldCreatedAt, personalaccesstoken.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PersonalAccessToken", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PersonalAccessToken fields.
func (pat *PersonalAccessToken) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case personalaccesstoken.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pat.ID = int(value.Int64)
		case personalaccesstoken.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pat.Name = value.String
			}
		case personalaccesstoken.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				pat.UserID = int(value.Int64)
			}
		case personalaccesstoken.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				pat.Token = value.String
			}
		case personalaccesstoken.FieldAbilities:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field abilities", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pat.Abilities); err != nil {
					return fmt.Errorf("unmarshal field abilities: %w", err)
				}
			}
		case personalaccesstoken.FieldExpirationAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expiration_at", values[i])
			} else if value.Valid {
				pat.ExpirationAt = value.Time
			}
		case personalaccesstoken.FieldLastUsedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_used_at", values[i])
			} else if value.Valid {
				pat.LastUsedAt = new(time.Time)
				*pat.LastUsedAt = value.Time
			}
		case personalaccesstoken.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pat.CreatedAt = value.Time
			}
		case personalaccesstoken.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pat.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the PersonalAccessToken entity.
func (pat *PersonalAccessToken) QueryUser() *UserQuery {
	return (&PersonalAccessTokenClient{config: pat.config}).QueryUser(pat)
}

// Update returns a builder for updating this PersonalAccessToken.
// Note that you need to call PersonalAccessToken.Unwrap() before calling this method if this PersonalAccessToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (pat *PersonalAccessToken) Update() *PersonalAccessTokenUpdateOne {
	return (&PersonalAccessTokenClient{config: pat.config}).UpdateOne(pat)
}

// Unwrap unwraps the PersonalAccessToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pat *PersonalAccessToken) Unwrap() *PersonalAccessToken {
	tx, ok := pat.config.driver.(*txDriver)
	if !ok {
		panic("ent: PersonalAccessToken is not a transactional entity")
	}
	pat.config.driver = tx.drv
	return pat
}

// String implements the fmt.Stringer.
func (pat *PersonalAccessToken) String() string {
	var builder strings.Builder
	builder.WriteString("PersonalAccessToken(")
	builder.WriteString(fmt.Sprintf("id=%v", pat.ID))
	builder.WriteString(", name=")
	builder.WriteString(pat.Name)
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", pat.UserID))
	builder.WriteString(", token=")
	builder.WriteString(pat.Token)
	builder.WriteString(", abilities=")
	builder.WriteString(fmt.Sprintf("%v", pat.Abilities))
	builder.WriteString(", expiration_at=")
	builder.WriteString(pat.ExpirationAt.Format(time.ANSIC))
	if v := pat.LastUsedAt; v != nil {
		builder.WriteString(", last_used_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", created_at=")
	builder.WriteString(pat.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(pat.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// PersonalAccessTokens is a parsable slice of PersonalAccessToken.
type PersonalAccessTokens []*PersonalAccessToken

func (pat PersonalAccessTokens) config(cfg config) {
	for _i := range pat {
		pat[_i].config = cfg
	}
}
