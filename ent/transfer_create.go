// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/s14t284/simplebank/ent/account"
	"github.com/s14t284/simplebank/ent/transfer"
)

// TransferCreate is the builder for creating a Transfer entity.
type TransferCreate struct {
	config
	mutation *TransferMutation
	hooks    []Hook
}

// SetFromAccountID sets the "from_account_id" field.
func (tc *TransferCreate) SetFromAccountID(i int) *TransferCreate {
	tc.mutation.SetFromAccountID(i)
	return tc
}

// SetToAccountID sets the "to_account_id" field.
func (tc *TransferCreate) SetToAccountID(i int) *TransferCreate {
	tc.mutation.SetToAccountID(i)
	return tc
}

// SetAmount sets the "amount" field.
func (tc *TransferCreate) SetAmount(i int) *TransferCreate {
	tc.mutation.SetAmount(i)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TransferCreate) SetCreatedAt(t time.Time) *TransferCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TransferCreate) SetNillableCreatedAt(t *time.Time) *TransferCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetFromAccountsID sets the "from_accounts" edge to the Account entity by ID.
func (tc *TransferCreate) SetFromAccountsID(id int) *TransferCreate {
	tc.mutation.SetFromAccountsID(id)
	return tc
}

// SetFromAccounts sets the "from_accounts" edge to the Account entity.
func (tc *TransferCreate) SetFromAccounts(a *Account) *TransferCreate {
	return tc.SetFromAccountsID(a.ID)
}

// SetToAccountsID sets the "to_accounts" edge to the Account entity by ID.
func (tc *TransferCreate) SetToAccountsID(id int) *TransferCreate {
	tc.mutation.SetToAccountsID(id)
	return tc
}

// SetToAccounts sets the "to_accounts" edge to the Account entity.
func (tc *TransferCreate) SetToAccounts(a *Account) *TransferCreate {
	return tc.SetToAccountsID(a.ID)
}

// Mutation returns the TransferMutation object of the builder.
func (tc *TransferCreate) Mutation() *TransferMutation {
	return tc.mutation
}

// Save creates the Transfer in the database.
func (tc *TransferCreate) Save(ctx context.Context) (*Transfer, error) {
	var (
		err  error
		node *Transfer
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransferMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TransferCreate) SaveX(ctx context.Context) *Transfer {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TransferCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TransferCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TransferCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := transfer.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TransferCreate) check() error {
	if _, ok := tc.mutation.FromAccountID(); !ok {
		return &ValidationError{Name: "from_account_id", err: errors.New(`ent: missing required field "Transfer.from_account_id"`)}
	}
	if _, ok := tc.mutation.ToAccountID(); !ok {
		return &ValidationError{Name: "to_account_id", err: errors.New(`ent: missing required field "Transfer.to_account_id"`)}
	}
	if _, ok := tc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Transfer.amount"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Transfer.created_at"`)}
	}
	if _, ok := tc.mutation.FromAccountsID(); !ok {
		return &ValidationError{Name: "from_accounts", err: errors.New(`ent: missing required edge "Transfer.from_accounts"`)}
	}
	if _, ok := tc.mutation.ToAccountsID(); !ok {
		return &ValidationError{Name: "to_accounts", err: errors.New(`ent: missing required edge "Transfer.to_accounts"`)}
	}
	return nil
}

func (tc *TransferCreate) sqlSave(ctx context.Context) (*Transfer, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *TransferCreate) createSpec() (*Transfer, *sqlgraph.CreateSpec) {
	var (
		_node = &Transfer{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: transfer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: transfer.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: transfer.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: transfer.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := tc.mutation.FromAccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transfer.FromAccountsTable,
			Columns: []string{transfer.FromAccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.FromAccountID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ToAccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transfer.ToAccountsTable,
			Columns: []string{transfer.ToAccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ToAccountID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TransferCreateBulk is the builder for creating many Transfer entities in bulk.
type TransferCreateBulk struct {
	config
	builders []*TransferCreate
}

// Save creates the Transfer entities in the database.
func (tcb *TransferCreateBulk) Save(ctx context.Context) ([]*Transfer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Transfer, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TransferMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TransferCreateBulk) SaveX(ctx context.Context) []*Transfer {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TransferCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TransferCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
