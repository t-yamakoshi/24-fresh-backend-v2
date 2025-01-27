// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/adapter/entgen/followsmodel"
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/adapter/entgen/usermodel"
)

// FollowsModelCreate is the builder for creating a FollowsModel entity.
type FollowsModelCreate struct {
	config
	mutation *FollowsModelMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (fmc *FollowsModelCreate) SetCreatedAt(t time.Time) *FollowsModelCreate {
	fmc.mutation.SetCreatedAt(t)
	return fmc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fmc *FollowsModelCreate) SetNillableCreatedAt(t *time.Time) *FollowsModelCreate {
	if t != nil {
		fmc.SetCreatedAt(*t)
	}
	return fmc
}

// SetUpdatedAt sets the "updated_at" field.
func (fmc *FollowsModelCreate) SetUpdatedAt(t time.Time) *FollowsModelCreate {
	fmc.mutation.SetUpdatedAt(t)
	return fmc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fmc *FollowsModelCreate) SetNillableUpdatedAt(t *time.Time) *FollowsModelCreate {
	if t != nil {
		fmc.SetUpdatedAt(*t)
	}
	return fmc
}

// SetFollowerUserID sets the "follower_user_id" field.
func (fmc *FollowsModelCreate) SetFollowerUserID(i int64) *FollowsModelCreate {
	fmc.mutation.SetFollowerUserID(i)
	return fmc
}

// SetFolloweeUserID sets the "followee_user_id" field.
func (fmc *FollowsModelCreate) SetFolloweeUserID(i int64) *FollowsModelCreate {
	fmc.mutation.SetFolloweeUserID(i)
	return fmc
}

// SetFollowerID sets the "follower" edge to the UserModel entity by ID.
func (fmc *FollowsModelCreate) SetFollowerID(id int64) *FollowsModelCreate {
	fmc.mutation.SetFollowerID(id)
	return fmc
}

// SetFollower sets the "follower" edge to the UserModel entity.
func (fmc *FollowsModelCreate) SetFollower(u *UserModel) *FollowsModelCreate {
	return fmc.SetFollowerID(u.ID)
}

// SetFolloweeID sets the "followee" edge to the UserModel entity by ID.
func (fmc *FollowsModelCreate) SetFolloweeID(id int64) *FollowsModelCreate {
	fmc.mutation.SetFolloweeID(id)
	return fmc
}

// SetFollowee sets the "followee" edge to the UserModel entity.
func (fmc *FollowsModelCreate) SetFollowee(u *UserModel) *FollowsModelCreate {
	return fmc.SetFolloweeID(u.ID)
}

// Mutation returns the FollowsModelMutation object of the builder.
func (fmc *FollowsModelCreate) Mutation() *FollowsModelMutation {
	return fmc.mutation
}

// Save creates the FollowsModel in the database.
func (fmc *FollowsModelCreate) Save(ctx context.Context) (*FollowsModel, error) {
	return withHooks(ctx, fmc.sqlSave, fmc.mutation, fmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fmc *FollowsModelCreate) SaveX(ctx context.Context) *FollowsModel {
	v, err := fmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fmc *FollowsModelCreate) Exec(ctx context.Context) error {
	_, err := fmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fmc *FollowsModelCreate) ExecX(ctx context.Context) {
	if err := fmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fmc *FollowsModelCreate) check() error {
	if _, ok := fmc.mutation.FollowerUserID(); !ok {
		return &ValidationError{Name: "follower_user_id", err: errors.New(`entgen: missing required field "FollowsModel.follower_user_id"`)}
	}
	if _, ok := fmc.mutation.FolloweeUserID(); !ok {
		return &ValidationError{Name: "followee_user_id", err: errors.New(`entgen: missing required field "FollowsModel.followee_user_id"`)}
	}
	if len(fmc.mutation.FollowerIDs()) == 0 {
		return &ValidationError{Name: "follower", err: errors.New(`entgen: missing required edge "FollowsModel.follower"`)}
	}
	if len(fmc.mutation.FolloweeIDs()) == 0 {
		return &ValidationError{Name: "followee", err: errors.New(`entgen: missing required edge "FollowsModel.followee"`)}
	}
	return nil
}

func (fmc *FollowsModelCreate) sqlSave(ctx context.Context) (*FollowsModel, error) {
	if err := fmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	fmc.mutation.id = &_node.ID
	fmc.mutation.done = true
	return _node, nil
}

func (fmc *FollowsModelCreate) createSpec() (*FollowsModel, *sqlgraph.CreateSpec) {
	var (
		_node = &FollowsModel{config: fmc.config}
		_spec = sqlgraph.NewCreateSpec(followsmodel.Table, sqlgraph.NewFieldSpec(followsmodel.FieldID, field.TypeInt))
	)
	if value, ok := fmc.mutation.CreatedAt(); ok {
		_spec.SetField(followsmodel.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := fmc.mutation.UpdatedAt(); ok {
		_spec.SetField(followsmodel.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := fmc.mutation.FollowerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   followsmodel.FollowerTable,
			Columns: []string{followsmodel.FollowerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usermodel.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.FollowerUserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fmc.mutation.FolloweeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   followsmodel.FolloweeTable,
			Columns: []string{followsmodel.FolloweeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usermodel.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.FolloweeUserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FollowsModelCreateBulk is the builder for creating many FollowsModel entities in bulk.
type FollowsModelCreateBulk struct {
	config
	err      error
	builders []*FollowsModelCreate
}

// Save creates the FollowsModel entities in the database.
func (fmcb *FollowsModelCreateBulk) Save(ctx context.Context) ([]*FollowsModel, error) {
	if fmcb.err != nil {
		return nil, fmcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fmcb.builders))
	nodes := make([]*FollowsModel, len(fmcb.builders))
	mutators := make([]Mutator, len(fmcb.builders))
	for i := range fmcb.builders {
		func(i int, root context.Context) {
			builder := fmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FollowsModelMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fmcb *FollowsModelCreateBulk) SaveX(ctx context.Context) []*FollowsModel {
	v, err := fmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fmcb *FollowsModelCreateBulk) Exec(ctx context.Context) error {
	_, err := fmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fmcb *FollowsModelCreateBulk) ExecX(ctx context.Context) {
	if err := fmcb.Exec(ctx); err != nil {
		panic(err)
	}
}