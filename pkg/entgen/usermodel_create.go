// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/entgen/usermodel"
)

// UserModelCreate is the builder for creating a UserModel entity.
type UserModelCreate struct {
	config
	mutation *UserModelMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (umc *UserModelCreate) SetName(s string) *UserModelCreate {
	umc.mutation.SetName(s)
	return umc
}

// SetUserName sets the "user_name" field.
func (umc *UserModelCreate) SetUserName(s string) *UserModelCreate {
	umc.mutation.SetUserName(s)
	return umc
}

// SetEmail sets the "email" field.
func (umc *UserModelCreate) SetEmail(s string) *UserModelCreate {
	umc.mutation.SetEmail(s)
	return umc
}

// Mutation returns the UserModelMutation object of the builder.
func (umc *UserModelCreate) Mutation() *UserModelMutation {
	return umc.mutation
}

// Save creates the UserModel in the database.
func (umc *UserModelCreate) Save(ctx context.Context) (*UserModel, error) {
	return withHooks(ctx, umc.sqlSave, umc.mutation, umc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (umc *UserModelCreate) SaveX(ctx context.Context) *UserModel {
	v, err := umc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (umc *UserModelCreate) Exec(ctx context.Context) error {
	_, err := umc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (umc *UserModelCreate) ExecX(ctx context.Context) {
	if err := umc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (umc *UserModelCreate) check() error {
	if _, ok := umc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`entgen: missing required field "UserModel.name"`)}
	}
	if _, ok := umc.mutation.UserName(); !ok {
		return &ValidationError{Name: "user_name", err: errors.New(`entgen: missing required field "UserModel.user_name"`)}
	}
	if _, ok := umc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`entgen: missing required field "UserModel.email"`)}
	}
	return nil
}

func (umc *UserModelCreate) sqlSave(ctx context.Context) (*UserModel, error) {
	if err := umc.check(); err != nil {
		return nil, err
	}
	_node, _spec := umc.createSpec()
	if err := sqlgraph.CreateNode(ctx, umc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	umc.mutation.id = &_node.ID
	umc.mutation.done = true
	return _node, nil
}

func (umc *UserModelCreate) createSpec() (*UserModel, *sqlgraph.CreateSpec) {
	var (
		_node = &UserModel{config: umc.config}
		_spec = sqlgraph.NewCreateSpec(usermodel.Table, sqlgraph.NewFieldSpec(usermodel.FieldID, field.TypeInt))
	)
	if value, ok := umc.mutation.Name(); ok {
		_spec.SetField(usermodel.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := umc.mutation.UserName(); ok {
		_spec.SetField(usermodel.FieldUserName, field.TypeString, value)
		_node.UserName = value
	}
	if value, ok := umc.mutation.Email(); ok {
		_spec.SetField(usermodel.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	return _node, _spec
}

// UserModelCreateBulk is the builder for creating many UserModel entities in bulk.
type UserModelCreateBulk struct {
	config
	err      error
	builders []*UserModelCreate
}

// Save creates the UserModel entities in the database.
func (umcb *UserModelCreateBulk) Save(ctx context.Context) ([]*UserModel, error) {
	if umcb.err != nil {
		return nil, umcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(umcb.builders))
	nodes := make([]*UserModel, len(umcb.builders))
	mutators := make([]Mutator, len(umcb.builders))
	for i := range umcb.builders {
		func(i int, root context.Context) {
			builder := umcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserModelMutation)
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
					_, err = mutators[i+1].Mutate(root, umcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, umcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, umcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (umcb *UserModelCreateBulk) SaveX(ctx context.Context) []*UserModel {
	v, err := umcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (umcb *UserModelCreateBulk) Exec(ctx context.Context) error {
	_, err := umcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (umcb *UserModelCreateBulk) ExecX(ctx context.Context) {
	if err := umcb.Exec(ctx); err != nil {
		panic(err)
	}
}