// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tiagoposse/connect/ent/apikey"
	"github.com/tiagoposse/connect/ent/predicate"
)

// ApiKeyUpdate is the builder for updating ApiKey entities.
type ApiKeyUpdate struct {
	config
	hooks    []Hook
	mutation *ApiKeyMutation
}

// Where appends a list predicates to the ApiKeyUpdate builder.
func (aku *ApiKeyUpdate) Where(ps ...predicate.ApiKey) *ApiKeyUpdate {
	aku.mutation.Where(ps...)
	return aku
}

// Mutation returns the ApiKeyMutation object of the builder.
func (aku *ApiKeyUpdate) Mutation() *ApiKeyMutation {
	return aku.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aku *ApiKeyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, aku.sqlSave, aku.mutation, aku.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aku *ApiKeyUpdate) SaveX(ctx context.Context) int {
	affected, err := aku.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aku *ApiKeyUpdate) Exec(ctx context.Context) error {
	_, err := aku.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aku *ApiKeyUpdate) ExecX(ctx context.Context) {
	if err := aku.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aku *ApiKeyUpdate) check() error {
	if _, ok := aku.mutation.UserID(); aku.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ApiKey.user"`)
	}
	return nil
}

func (aku *ApiKeyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := aku.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(apikey.Table, apikey.Columns, sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeInt))
	if ps := aku.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aku.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apikey.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aku.mutation.done = true
	return n, nil
}

// ApiKeyUpdateOne is the builder for updating a single ApiKey entity.
type ApiKeyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ApiKeyMutation
}

// Mutation returns the ApiKeyMutation object of the builder.
func (akuo *ApiKeyUpdateOne) Mutation() *ApiKeyMutation {
	return akuo.mutation
}

// Where appends a list predicates to the ApiKeyUpdate builder.
func (akuo *ApiKeyUpdateOne) Where(ps ...predicate.ApiKey) *ApiKeyUpdateOne {
	akuo.mutation.Where(ps...)
	return akuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (akuo *ApiKeyUpdateOne) Select(field string, fields ...string) *ApiKeyUpdateOne {
	akuo.fields = append([]string{field}, fields...)
	return akuo
}

// Save executes the query and returns the updated ApiKey entity.
func (akuo *ApiKeyUpdateOne) Save(ctx context.Context) (*ApiKey, error) {
	return withHooks(ctx, akuo.sqlSave, akuo.mutation, akuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (akuo *ApiKeyUpdateOne) SaveX(ctx context.Context) *ApiKey {
	node, err := akuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (akuo *ApiKeyUpdateOne) Exec(ctx context.Context) error {
	_, err := akuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (akuo *ApiKeyUpdateOne) ExecX(ctx context.Context) {
	if err := akuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (akuo *ApiKeyUpdateOne) check() error {
	if _, ok := akuo.mutation.UserID(); akuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ApiKey.user"`)
	}
	return nil
}

func (akuo *ApiKeyUpdateOne) sqlSave(ctx context.Context) (_node *ApiKey, err error) {
	if err := akuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(apikey.Table, apikey.Columns, sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeInt))
	id, ok := akuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ApiKey.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := akuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apikey.FieldID)
		for _, f := range fields {
			if !apikey.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != apikey.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := akuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &ApiKey{config: akuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, akuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apikey.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	akuo.mutation.done = true
	return _node, nil
}
