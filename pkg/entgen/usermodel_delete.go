// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/entgen/predicate"
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/entgen/usermodel"
)

// UserModelDelete is the builder for deleting a UserModel entity.
type UserModelDelete struct {
	config
	hooks    []Hook
	mutation *UserModelMutation
}

// Where appends a list predicates to the UserModelDelete builder.
func (umd *UserModelDelete) Where(ps ...predicate.UserModel) *UserModelDelete {
	umd.mutation.Where(ps...)
	return umd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (umd *UserModelDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, umd.sqlExec, umd.mutation, umd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (umd *UserModelDelete) ExecX(ctx context.Context) int {
	n, err := umd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (umd *UserModelDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(usermodel.Table, sqlgraph.NewFieldSpec(usermodel.FieldID, field.TypeInt))
	if ps := umd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, umd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	umd.mutation.done = true
	return affected, err
}

// UserModelDeleteOne is the builder for deleting a single UserModel entity.
type UserModelDeleteOne struct {
	umd *UserModelDelete
}

// Where appends a list predicates to the UserModelDelete builder.
func (umdo *UserModelDeleteOne) Where(ps ...predicate.UserModel) *UserModelDeleteOne {
	umdo.umd.mutation.Where(ps...)
	return umdo
}

// Exec executes the deletion query.
func (umdo *UserModelDeleteOne) Exec(ctx context.Context) error {
	n, err := umdo.umd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{usermodel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (umdo *UserModelDeleteOne) ExecX(ctx context.Context) {
	if err := umdo.Exec(ctx); err != nil {
		panic(err)
	}
}
