// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"owknight-be/ent/predicate"
	"owknight-be/ent/role"
	"owknight-be/ent/user"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// RoleUpdate is the builder for updating Role entities.
type RoleUpdate struct {
	config
	hooks      []Hook
	mutation   *RoleMutation
	predicates []predicate.Role
}

// Where adds a new predicate for the builder.
func (ru *RoleUpdate) Where(ps ...predicate.Role) *RoleUpdate {
	ru.predicates = append(ru.predicates, ps...)
	return ru
}

// SetName sets the name field.
func (ru *RoleUpdate) SetName(s string) *RoleUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetDescription sets the description field.
func (ru *RoleUpdate) SetDescription(s string) *RoleUpdate {
	ru.mutation.SetDescription(s)
	return ru
}

// SetUpdatedAt sets the updatedAt field.
func (ru *RoleUpdate) SetUpdatedAt(t time.Time) *RoleUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetNillableUpdatedAt sets the updatedAt field if the given value is not nil.
func (ru *RoleUpdate) SetNillableUpdatedAt(t *time.Time) *RoleUpdate {
	if t != nil {
		ru.SetUpdatedAt(*t)
	}
	return ru
}

// SetCreatedAt sets the createdAt field.
func (ru *RoleUpdate) SetCreatedAt(t time.Time) *RoleUpdate {
	ru.mutation.SetCreatedAt(t)
	return ru
}

// AddUserIDs adds the users edge to User by ids.
func (ru *RoleUpdate) AddUserIDs(ids ...int) *RoleUpdate {
	ru.mutation.AddUserIDs(ids...)
	return ru
}

// AddUsers adds the users edges to User.
func (ru *RoleUpdate) AddUsers(u ...*User) *RoleUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ru.AddUserIDs(ids...)
}

// RemoveUserIDs removes the users edge to User by ids.
func (ru *RoleUpdate) RemoveUserIDs(ids ...int) *RoleUpdate {
	ru.mutation.RemoveUserIDs(ids...)
	return ru
}

// RemoveUsers removes users edges to User.
func (ru *RoleUpdate) RemoveUsers(u ...*User) *RoleUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ru.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ru *RoleUpdate) Save(ctx context.Context) (int, error) {
	if _, ok := ru.mutation.CreatedAt(); !ok {
		v := role.UpdateDefaultCreatedAt()
		ru.mutation.SetCreatedAt(v)
	}

	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoleUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoleUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoleUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   role.Table,
			Columns: role.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: role.FieldID,
			},
		},
	}
	if ps := ru.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldName,
		})
	}
	if value, ok := ru.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldDescription,
		})
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: role.FieldUpdatedAt,
		})
	}
	if value, ok := ru.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: role.FieldCreatedAt,
		})
	}
	if nodes := ru.mutation.RemovedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.UsersTable,
			Columns: role.UsersPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.UsersTable,
			Columns: role.UsersPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RoleUpdateOne is the builder for updating a single Role entity.
type RoleUpdateOne struct {
	config
	hooks    []Hook
	mutation *RoleMutation
}

// SetName sets the name field.
func (ruo *RoleUpdateOne) SetName(s string) *RoleUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetDescription sets the description field.
func (ruo *RoleUpdateOne) SetDescription(s string) *RoleUpdateOne {
	ruo.mutation.SetDescription(s)
	return ruo
}

// SetUpdatedAt sets the updatedAt field.
func (ruo *RoleUpdateOne) SetUpdatedAt(t time.Time) *RoleUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetNillableUpdatedAt sets the updatedAt field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableUpdatedAt(t *time.Time) *RoleUpdateOne {
	if t != nil {
		ruo.SetUpdatedAt(*t)
	}
	return ruo
}

// SetCreatedAt sets the createdAt field.
func (ruo *RoleUpdateOne) SetCreatedAt(t time.Time) *RoleUpdateOne {
	ruo.mutation.SetCreatedAt(t)
	return ruo
}

// AddUserIDs adds the users edge to User by ids.
func (ruo *RoleUpdateOne) AddUserIDs(ids ...int) *RoleUpdateOne {
	ruo.mutation.AddUserIDs(ids...)
	return ruo
}

// AddUsers adds the users edges to User.
func (ruo *RoleUpdateOne) AddUsers(u ...*User) *RoleUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ruo.AddUserIDs(ids...)
}

// RemoveUserIDs removes the users edge to User by ids.
func (ruo *RoleUpdateOne) RemoveUserIDs(ids ...int) *RoleUpdateOne {
	ruo.mutation.RemoveUserIDs(ids...)
	return ruo
}

// RemoveUsers removes users edges to User.
func (ruo *RoleUpdateOne) RemoveUsers(u ...*User) *RoleUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ruo.RemoveUserIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (ruo *RoleUpdateOne) Save(ctx context.Context) (*Role, error) {
	if _, ok := ruo.mutation.CreatedAt(); !ok {
		v := role.UpdateDefaultCreatedAt()
		ruo.mutation.SetCreatedAt(v)
	}

	var (
		err  error
		node *Role
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoleUpdateOne) SaveX(ctx context.Context) *Role {
	r, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return r
}

// Exec executes the query on the entity.
func (ruo *RoleUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoleUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RoleUpdateOne) sqlSave(ctx context.Context) (r *Role, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   role.Table,
			Columns: role.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: role.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Role.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := ruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldName,
		})
	}
	if value, ok := ruo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldDescription,
		})
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: role.FieldUpdatedAt,
		})
	}
	if value, ok := ruo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: role.FieldCreatedAt,
		})
	}
	if nodes := ruo.mutation.RemovedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.UsersTable,
			Columns: role.UsersPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.UsersTable,
			Columns: role.UsersPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	r = &Role{config: ruo.config}
	_spec.Assign = r.assignValues
	_spec.ScanValues = r.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return r, nil
}
