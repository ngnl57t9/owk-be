// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"owknight-be/ent/profile"
	"owknight-be/ent/user"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ProfileCreate is the builder for creating a Profile entity.
type ProfileCreate struct {
	config
	mutation *ProfileMutation
	hooks    []Hook
}

// SetEmail sets the email field.
func (pc *ProfileCreate) SetEmail(s string) *ProfileCreate {
	pc.mutation.SetEmail(s)
	return pc
}

// SetRemark sets the remark field.
func (pc *ProfileCreate) SetRemark(s string) *ProfileCreate {
	pc.mutation.SetRemark(s)
	return pc
}

// SetUpdatedAt sets the updatedAt field.
func (pc *ProfileCreate) SetUpdatedAt(t time.Time) *ProfileCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the updatedAt field if the given value is not nil.
func (pc *ProfileCreate) SetNillableUpdatedAt(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetCreatedAt sets the createdAt field.
func (pc *ProfileCreate) SetCreatedAt(t time.Time) *ProfileCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetUserID sets the user edge to User by id.
func (pc *ProfileCreate) SetUserID(id int) *ProfileCreate {
	pc.mutation.SetUserID(id)
	return pc
}

// SetUser sets the user edge to User.
func (pc *ProfileCreate) SetUser(u *User) *ProfileCreate {
	return pc.SetUserID(u.ID)
}

// Save creates the Profile in the database.
func (pc *ProfileCreate) Save(ctx context.Context) (*Profile, error) {
	if _, ok := pc.mutation.Email(); !ok {
		return nil, errors.New("ent: missing required field \"email\"")
	}
	if _, ok := pc.mutation.Remark(); !ok {
		return nil, errors.New("ent: missing required field \"remark\"")
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := profile.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return nil, errors.New("ent: missing required field \"createdAt\"")
	}
	if _, ok := pc.mutation.UserID(); !ok {
		return nil, errors.New("ent: missing required edge \"user\"")
	}
	var (
		err  error
		node *Profile
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProfileCreate) SaveX(ctx context.Context) *Profile {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *ProfileCreate) sqlSave(ctx context.Context) (*Profile, error) {
	var (
		pr    = &Profile{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: profile.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: profile.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldEmail,
		})
		pr.Email = value
	}
	if value, ok := pc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldRemark,
		})
		pr.Remark = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldUpdatedAt,
		})
		pr.UpdatedAt = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldCreatedAt,
		})
		pr.CreatedAt = value
	}
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   profile.UserTable,
			Columns: []string{profile.UserColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pr.ID = int(id)
	return pr, nil
}