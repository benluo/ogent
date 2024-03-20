// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"ariga.io/ogent/internal/integration/ogent/ent/hat"
	"ariga.io/ogent/internal/integration/ogent/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HatDelete is the builder for deleting a Hat entity.
type HatDelete struct {
	config
	hooks    []Hook
	mutation *HatMutation
}

// Where appends a list predicates to the HatDelete builder.
func (hd *HatDelete) Where(ps ...predicate.Hat) *HatDelete {
	hd.mutation.Where(ps...)
	return hd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hd *HatDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, hd.sqlExec, hd.mutation, hd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (hd *HatDelete) ExecX(ctx context.Context) int {
	n, err := hd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hd *HatDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(hat.Table, sqlgraph.NewFieldSpec(hat.FieldID, field.TypeInt))
	if ps := hd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, hd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	hd.mutation.done = true
	return affected, err
}

// HatDeleteOne is the builder for deleting a single Hat entity.
type HatDeleteOne struct {
	hd *HatDelete
}

// Where appends a list predicates to the HatDelete builder.
func (hdo *HatDeleteOne) Where(ps ...predicate.Hat) *HatDeleteOne {
	hdo.hd.mutation.Where(ps...)
	return hdo
}

// Exec executes the deletion query.
func (hdo *HatDeleteOne) Exec(ctx context.Context) error {
	n, err := hdo.hd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{hat.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hdo *HatDeleteOne) ExecX(ctx context.Context) {
	if err := hdo.Exec(ctx); err != nil {
		panic(err)
	}
}