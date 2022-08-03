// Copyright 2021-present The Atlas Authors. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"ariga.io/atlas/cmd/atlas/internal/migrate/ent/internal"
	"ariga.io/atlas/cmd/atlas/internal/migrate/ent/predicate"
	"ariga.io/atlas/cmd/atlas/internal/migrate/ent/revision"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RevisionUpdate is the builder for updating Revision entities.
type RevisionUpdate struct {
	config
	hooks    []Hook
	mutation *RevisionMutation
}

// Where appends a list predicates to the RevisionUpdate builder.
func (ru *RevisionUpdate) Where(ps ...predicate.Revision) *RevisionUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetApplied sets the "applied" field.
func (ru *RevisionUpdate) SetApplied(i int) *RevisionUpdate {
	ru.mutation.ResetApplied()
	ru.mutation.SetApplied(i)
	return ru
}

// AddApplied adds i to the "applied" field.
func (ru *RevisionUpdate) AddApplied(i int) *RevisionUpdate {
	ru.mutation.AddApplied(i)
	return ru
}

// SetTotal sets the "total" field.
func (ru *RevisionUpdate) SetTotal(i int) *RevisionUpdate {
	ru.mutation.ResetTotal()
	ru.mutation.SetTotal(i)
	return ru
}

// AddTotal adds i to the "total" field.
func (ru *RevisionUpdate) AddTotal(i int) *RevisionUpdate {
	ru.mutation.AddTotal(i)
	return ru
}

// SetExecutionTime sets the "execution_time" field.
func (ru *RevisionUpdate) SetExecutionTime(t time.Duration) *RevisionUpdate {
	ru.mutation.ResetExecutionTime()
	ru.mutation.SetExecutionTime(t)
	return ru
}

// AddExecutionTime adds t to the "execution_time" field.
func (ru *RevisionUpdate) AddExecutionTime(t time.Duration) *RevisionUpdate {
	ru.mutation.AddExecutionTime(t)
	return ru
}

// SetError sets the "error" field.
func (ru *RevisionUpdate) SetError(s string) *RevisionUpdate {
	ru.mutation.SetError(s)
	return ru
}

// SetNillableError sets the "error" field if the given value is not nil.
func (ru *RevisionUpdate) SetNillableError(s *string) *RevisionUpdate {
	if s != nil {
		ru.SetError(*s)
	}
	return ru
}

// ClearError clears the value of the "error" field.
func (ru *RevisionUpdate) ClearError() *RevisionUpdate {
	ru.mutation.ClearError()
	return ru
}

// SetHash sets the "hash" field.
func (ru *RevisionUpdate) SetHash(s string) *RevisionUpdate {
	ru.mutation.SetHash(s)
	return ru
}

// SetPartialHashes sets the "partial_hashes" field.
func (ru *RevisionUpdate) SetPartialHashes(s []string) *RevisionUpdate {
	ru.mutation.SetPartialHashes(s)
	return ru
}

// ClearPartialHashes clears the value of the "partial_hashes" field.
func (ru *RevisionUpdate) ClearPartialHashes() *RevisionUpdate {
	ru.mutation.ClearPartialHashes()
	return ru
}

// SetOperatorVersion sets the "operator_version" field.
func (ru *RevisionUpdate) SetOperatorVersion(s string) *RevisionUpdate {
	ru.mutation.SetOperatorVersion(s)
	return ru
}

// SetMeta sets the "meta" field.
func (ru *RevisionUpdate) SetMeta(m map[string]string) *RevisionUpdate {
	ru.mutation.SetMeta(m)
	return ru
}

// Mutation returns the RevisionMutation object of the builder.
func (ru *RevisionUpdate) Mutation() *RevisionMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RevisionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RevisionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RevisionUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RevisionUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RevisionUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RevisionUpdate) check() error {
	if v, ok := ru.mutation.Total(); ok {
		if err := revision.TotalValidator(v); err != nil {
			return &ValidationError{Name: "total", err: fmt.Errorf(`ent: validator failed for field "Revision.total": %w`, err)}
		}
	}
	return nil
}

func (ru *RevisionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   revision.Table,
			Columns: revision.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: revision.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Applied(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: revision.FieldApplied,
		})
	}
	if value, ok := ru.mutation.AddedApplied(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: revision.FieldApplied,
		})
	}
	if value, ok := ru.mutation.Total(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: revision.FieldTotal,
		})
	}
	if value, ok := ru.mutation.AddedTotal(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: revision.FieldTotal,
		})
	}
	if value, ok := ru.mutation.ExecutionTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: revision.FieldExecutionTime,
		})
	}
	if value, ok := ru.mutation.AddedExecutionTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: revision.FieldExecutionTime,
		})
	}
	if value, ok := ru.mutation.Error(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: revision.FieldError,
		})
	}
	if ru.mutation.ErrorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: revision.FieldError,
		})
	}
	if value, ok := ru.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: revision.FieldHash,
		})
	}
	if value, ok := ru.mutation.PartialHashes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: revision.FieldPartialHashes,
		})
	}
	if ru.mutation.PartialHashesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: revision.FieldPartialHashes,
		})
	}
	if value, ok := ru.mutation.OperatorVersion(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: revision.FieldOperatorVersion,
		})
	}
	if value, ok := ru.mutation.Meta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: revision.FieldMeta,
		})
	}
	_spec.Node.Schema = ru.schemaConfig.Revision
	ctx = internal.NewSchemaConfigContext(ctx, ru.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{revision.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// RevisionUpdateOne is the builder for updating a single Revision entity.
type RevisionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RevisionMutation
}

// SetApplied sets the "applied" field.
func (ruo *RevisionUpdateOne) SetApplied(i int) *RevisionUpdateOne {
	ruo.mutation.ResetApplied()
	ruo.mutation.SetApplied(i)
	return ruo
}

// AddApplied adds i to the "applied" field.
func (ruo *RevisionUpdateOne) AddApplied(i int) *RevisionUpdateOne {
	ruo.mutation.AddApplied(i)
	return ruo
}

// SetTotal sets the "total" field.
func (ruo *RevisionUpdateOne) SetTotal(i int) *RevisionUpdateOne {
	ruo.mutation.ResetTotal()
	ruo.mutation.SetTotal(i)
	return ruo
}

// AddTotal adds i to the "total" field.
func (ruo *RevisionUpdateOne) AddTotal(i int) *RevisionUpdateOne {
	ruo.mutation.AddTotal(i)
	return ruo
}

// SetExecutionTime sets the "execution_time" field.
func (ruo *RevisionUpdateOne) SetExecutionTime(t time.Duration) *RevisionUpdateOne {
	ruo.mutation.ResetExecutionTime()
	ruo.mutation.SetExecutionTime(t)
	return ruo
}

// AddExecutionTime adds t to the "execution_time" field.
func (ruo *RevisionUpdateOne) AddExecutionTime(t time.Duration) *RevisionUpdateOne {
	ruo.mutation.AddExecutionTime(t)
	return ruo
}

// SetError sets the "error" field.
func (ruo *RevisionUpdateOne) SetError(s string) *RevisionUpdateOne {
	ruo.mutation.SetError(s)
	return ruo
}

// SetNillableError sets the "error" field if the given value is not nil.
func (ruo *RevisionUpdateOne) SetNillableError(s *string) *RevisionUpdateOne {
	if s != nil {
		ruo.SetError(*s)
	}
	return ruo
}

// ClearError clears the value of the "error" field.
func (ruo *RevisionUpdateOne) ClearError() *RevisionUpdateOne {
	ruo.mutation.ClearError()
	return ruo
}

// SetHash sets the "hash" field.
func (ruo *RevisionUpdateOne) SetHash(s string) *RevisionUpdateOne {
	ruo.mutation.SetHash(s)
	return ruo
}

// SetPartialHashes sets the "partial_hashes" field.
func (ruo *RevisionUpdateOne) SetPartialHashes(s []string) *RevisionUpdateOne {
	ruo.mutation.SetPartialHashes(s)
	return ruo
}

// ClearPartialHashes clears the value of the "partial_hashes" field.
func (ruo *RevisionUpdateOne) ClearPartialHashes() *RevisionUpdateOne {
	ruo.mutation.ClearPartialHashes()
	return ruo
}

// SetOperatorVersion sets the "operator_version" field.
func (ruo *RevisionUpdateOne) SetOperatorVersion(s string) *RevisionUpdateOne {
	ruo.mutation.SetOperatorVersion(s)
	return ruo
}

// SetMeta sets the "meta" field.
func (ruo *RevisionUpdateOne) SetMeta(m map[string]string) *RevisionUpdateOne {
	ruo.mutation.SetMeta(m)
	return ruo
}

// Mutation returns the RevisionMutation object of the builder.
func (ruo *RevisionUpdateOne) Mutation() *RevisionMutation {
	return ruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RevisionUpdateOne) Select(field string, fields ...string) *RevisionUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Revision entity.
func (ruo *RevisionUpdateOne) Save(ctx context.Context) (*Revision, error) {
	var (
		err  error
		node *Revision
	)
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RevisionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Revision)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RevisionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RevisionUpdateOne) SaveX(ctx context.Context) *Revision {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RevisionUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RevisionUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RevisionUpdateOne) check() error {
	if v, ok := ruo.mutation.Total(); ok {
		if err := revision.TotalValidator(v); err != nil {
			return &ValidationError{Name: "total", err: fmt.Errorf(`ent: validator failed for field "Revision.total": %w`, err)}
		}
	}
	return nil
}

func (ruo *RevisionUpdateOne) sqlSave(ctx context.Context) (_node *Revision, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   revision.Table,
			Columns: revision.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: revision.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Revision.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, revision.FieldID)
		for _, f := range fields {
			if !revision.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != revision.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Applied(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: revision.FieldApplied,
		})
	}
	if value, ok := ruo.mutation.AddedApplied(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: revision.FieldApplied,
		})
	}
	if value, ok := ruo.mutation.Total(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: revision.FieldTotal,
		})
	}
	if value, ok := ruo.mutation.AddedTotal(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: revision.FieldTotal,
		})
	}
	if value, ok := ruo.mutation.ExecutionTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: revision.FieldExecutionTime,
		})
	}
	if value, ok := ruo.mutation.AddedExecutionTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: revision.FieldExecutionTime,
		})
	}
	if value, ok := ruo.mutation.Error(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: revision.FieldError,
		})
	}
	if ruo.mutation.ErrorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: revision.FieldError,
		})
	}
	if value, ok := ruo.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: revision.FieldHash,
		})
	}
	if value, ok := ruo.mutation.PartialHashes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: revision.FieldPartialHashes,
		})
	}
	if ruo.mutation.PartialHashesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: revision.FieldPartialHashes,
		})
	}
	if value, ok := ruo.mutation.OperatorVersion(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: revision.FieldOperatorVersion,
		})
	}
	if value, ok := ruo.mutation.Meta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: revision.FieldMeta,
		})
	}
	_spec.Node.Schema = ruo.schemaConfig.Revision
	ctx = internal.NewSchemaConfigContext(ctx, ruo.schemaConfig)
	_node = &Revision{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{revision.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}