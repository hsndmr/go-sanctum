// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hsndmr/go-sanctum/ent/personalaccesstoken"
	"github.com/hsndmr/go-sanctum/ent/user"
)

// PersonalAccessTokenCreate is the builder for creating a PersonalAccessToken entity.
type PersonalAccessTokenCreate struct {
	config
	mutation *PersonalAccessTokenMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (patc *PersonalAccessTokenCreate) SetName(s string) *PersonalAccessTokenCreate {
	patc.mutation.SetName(s)
	return patc
}

// SetUserID sets the "user_id" field.
func (patc *PersonalAccessTokenCreate) SetUserID(i int) *PersonalAccessTokenCreate {
	patc.mutation.SetUserID(i)
	return patc
}

// SetToken sets the "token" field.
func (patc *PersonalAccessTokenCreate) SetToken(s string) *PersonalAccessTokenCreate {
	patc.mutation.SetToken(s)
	return patc
}

// SetAbilities sets the "abilities" field.
func (patc *PersonalAccessTokenCreate) SetAbilities(s []string) *PersonalAccessTokenCreate {
	patc.mutation.SetAbilities(s)
	return patc
}

// SetExpirationAt sets the "expiration_at" field.
func (patc *PersonalAccessTokenCreate) SetExpirationAt(t time.Time) *PersonalAccessTokenCreate {
	patc.mutation.SetExpirationAt(t)
	return patc
}

// SetLastUsedAt sets the "last_used_at" field.
func (patc *PersonalAccessTokenCreate) SetLastUsedAt(t time.Time) *PersonalAccessTokenCreate {
	patc.mutation.SetLastUsedAt(t)
	return patc
}

// SetNillableLastUsedAt sets the "last_used_at" field if the given value is not nil.
func (patc *PersonalAccessTokenCreate) SetNillableLastUsedAt(t *time.Time) *PersonalAccessTokenCreate {
	if t != nil {
		patc.SetLastUsedAt(*t)
	}
	return patc
}

// SetCreatedAt sets the "created_at" field.
func (patc *PersonalAccessTokenCreate) SetCreatedAt(t time.Time) *PersonalAccessTokenCreate {
	patc.mutation.SetCreatedAt(t)
	return patc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (patc *PersonalAccessTokenCreate) SetNillableCreatedAt(t *time.Time) *PersonalAccessTokenCreate {
	if t != nil {
		patc.SetCreatedAt(*t)
	}
	return patc
}

// SetUpdatedAt sets the "updated_at" field.
func (patc *PersonalAccessTokenCreate) SetUpdatedAt(t time.Time) *PersonalAccessTokenCreate {
	patc.mutation.SetUpdatedAt(t)
	return patc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (patc *PersonalAccessTokenCreate) SetNillableUpdatedAt(t *time.Time) *PersonalAccessTokenCreate {
	if t != nil {
		patc.SetUpdatedAt(*t)
	}
	return patc
}

// SetUser sets the "user" edge to the User entity.
func (patc *PersonalAccessTokenCreate) SetUser(u *User) *PersonalAccessTokenCreate {
	return patc.SetUserID(u.ID)
}

// Mutation returns the PersonalAccessTokenMutation object of the builder.
func (patc *PersonalAccessTokenCreate) Mutation() *PersonalAccessTokenMutation {
	return patc.mutation
}

// Save creates the PersonalAccessToken in the database.
func (patc *PersonalAccessTokenCreate) Save(ctx context.Context) (*PersonalAccessToken, error) {
	var (
		err  error
		node *PersonalAccessToken
	)
	patc.defaults()
	if len(patc.hooks) == 0 {
		if err = patc.check(); err != nil {
			return nil, err
		}
		node, err = patc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PersonalAccessTokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = patc.check(); err != nil {
				return nil, err
			}
			patc.mutation = mutation
			if node, err = patc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(patc.hooks) - 1; i >= 0; i-- {
			if patc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = patc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, patc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (patc *PersonalAccessTokenCreate) SaveX(ctx context.Context) *PersonalAccessToken {
	v, err := patc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (patc *PersonalAccessTokenCreate) Exec(ctx context.Context) error {
	_, err := patc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (patc *PersonalAccessTokenCreate) ExecX(ctx context.Context) {
	if err := patc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (patc *PersonalAccessTokenCreate) defaults() {
	if _, ok := patc.mutation.CreatedAt(); !ok {
		v := personalaccesstoken.DefaultCreatedAt()
		patc.mutation.SetCreatedAt(v)
	}
	if _, ok := patc.mutation.UpdatedAt(); !ok {
		v := personalaccesstoken.DefaultUpdatedAt()
		patc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (patc *PersonalAccessTokenCreate) check() error {
	if _, ok := patc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "PersonalAccessToken.name"`)}
	}
	if _, ok := patc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "PersonalAccessToken.user_id"`)}
	}
	if _, ok := patc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`ent: missing required field "PersonalAccessToken.token"`)}
	}
	if v, ok := patc.mutation.Token(); ok {
		if err := personalaccesstoken.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`ent: validator failed for field "PersonalAccessToken.token": %w`, err)}
		}
	}
	if _, ok := patc.mutation.ExpirationAt(); !ok {
		return &ValidationError{Name: "expiration_at", err: errors.New(`ent: missing required field "PersonalAccessToken.expiration_at"`)}
	}
	if _, ok := patc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "PersonalAccessToken.created_at"`)}
	}
	if _, ok := patc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "PersonalAccessToken.updated_at"`)}
	}
	if _, ok := patc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "PersonalAccessToken.user"`)}
	}
	return nil
}

func (patc *PersonalAccessTokenCreate) sqlSave(ctx context.Context) (*PersonalAccessToken, error) {
	_node, _spec := patc.createSpec()
	if err := sqlgraph.CreateNode(ctx, patc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (patc *PersonalAccessTokenCreate) createSpec() (*PersonalAccessToken, *sqlgraph.CreateSpec) {
	var (
		_node = &PersonalAccessToken{config: patc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: personalaccesstoken.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: personalaccesstoken.FieldID,
			},
		}
	)
	if value, ok := patc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personalaccesstoken.FieldName,
		})
		_node.Name = value
	}
	if value, ok := patc.mutation.Token(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personalaccesstoken.FieldToken,
		})
		_node.Token = value
	}
	if value, ok := patc.mutation.Abilities(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: personalaccesstoken.FieldAbilities,
		})
		_node.Abilities = value
	}
	if value, ok := patc.mutation.ExpirationAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: personalaccesstoken.FieldExpirationAt,
		})
		_node.ExpirationAt = value
	}
	if value, ok := patc.mutation.LastUsedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: personalaccesstoken.FieldLastUsedAt,
		})
		_node.LastUsedAt = &value
	}
	if value, ok := patc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: personalaccesstoken.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := patc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: personalaccesstoken.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := patc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   personalaccesstoken.UserTable,
			Columns: []string{personalaccesstoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PersonalAccessTokenCreateBulk is the builder for creating many PersonalAccessToken entities in bulk.
type PersonalAccessTokenCreateBulk struct {
	config
	builders []*PersonalAccessTokenCreate
}

// Save creates the PersonalAccessToken entities in the database.
func (patcb *PersonalAccessTokenCreateBulk) Save(ctx context.Context) ([]*PersonalAccessToken, error) {
	specs := make([]*sqlgraph.CreateSpec, len(patcb.builders))
	nodes := make([]*PersonalAccessToken, len(patcb.builders))
	mutators := make([]Mutator, len(patcb.builders))
	for i := range patcb.builders {
		func(i int, root context.Context) {
			builder := patcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PersonalAccessTokenMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, patcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, patcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, patcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (patcb *PersonalAccessTokenCreateBulk) SaveX(ctx context.Context) []*PersonalAccessToken {
	v, err := patcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (patcb *PersonalAccessTokenCreateBulk) Exec(ctx context.Context) error {
	_, err := patcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (patcb *PersonalAccessTokenCreateBulk) ExecX(ctx context.Context) {
	if err := patcb.Exec(ctx); err != nil {
		panic(err)
	}
}
