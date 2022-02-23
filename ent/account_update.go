// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/s14t284/simplebank/ent/account"
	"github.com/s14t284/simplebank/ent/entry"
	"github.com/s14t284/simplebank/ent/predicate"
	"github.com/s14t284/simplebank/ent/transfer"
	"github.com/s14t284/simplebank/ent/user"
)

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetOwner sets the "owner" field.
func (au *AccountUpdate) SetOwner(s string) *AccountUpdate {
	au.mutation.SetOwner(s)
	return au
}

// SetNillableOwner sets the "owner" field if the given value is not nil.
func (au *AccountUpdate) SetNillableOwner(s *string) *AccountUpdate {
	if s != nil {
		au.SetOwner(*s)
	}
	return au
}

// ClearOwner clears the value of the "owner" field.
func (au *AccountUpdate) ClearOwner() *AccountUpdate {
	au.mutation.ClearOwner()
	return au
}

// SetBalance sets the "balance" field.
func (au *AccountUpdate) SetBalance(i int) *AccountUpdate {
	au.mutation.ResetBalance()
	au.mutation.SetBalance(i)
	return au
}

// AddBalance adds i to the "balance" field.
func (au *AccountUpdate) AddBalance(i int) *AccountUpdate {
	au.mutation.AddBalance(i)
	return au
}

// SetCurrency sets the "currency" field.
func (au *AccountUpdate) SetCurrency(s string) *AccountUpdate {
	au.mutation.SetCurrency(s)
	return au
}

// AddEntryIDs adds the "entries" edge to the Entry entity by IDs.
func (au *AccountUpdate) AddEntryIDs(ids ...int) *AccountUpdate {
	au.mutation.AddEntryIDs(ids...)
	return au
}

// AddEntries adds the "entries" edges to the Entry entity.
func (au *AccountUpdate) AddEntries(e ...*Entry) *AccountUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return au.AddEntryIDs(ids...)
}

// AddFromTransferIDs adds the "from_transfers" edge to the Transfer entity by IDs.
func (au *AccountUpdate) AddFromTransferIDs(ids ...int) *AccountUpdate {
	au.mutation.AddFromTransferIDs(ids...)
	return au
}

// AddFromTransfers adds the "from_transfers" edges to the Transfer entity.
func (au *AccountUpdate) AddFromTransfers(t ...*Transfer) *AccountUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.AddFromTransferIDs(ids...)
}

// AddToTransferIDs adds the "to_transfers" edge to the Transfer entity by IDs.
func (au *AccountUpdate) AddToTransferIDs(ids ...int) *AccountUpdate {
	au.mutation.AddToTransferIDs(ids...)
	return au
}

// AddToTransfers adds the "to_transfers" edges to the Transfer entity.
func (au *AccountUpdate) AddToTransfers(t ...*Transfer) *AccountUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.AddToTransferIDs(ids...)
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (au *AccountUpdate) SetUsersID(id string) *AccountUpdate {
	au.mutation.SetUsersID(id)
	return au
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (au *AccountUpdate) SetNillableUsersID(id *string) *AccountUpdate {
	if id != nil {
		au = au.SetUsersID(*id)
	}
	return au
}

// SetUsers sets the "users" edge to the User entity.
func (au *AccountUpdate) SetUsers(u *User) *AccountUpdate {
	return au.SetUsersID(u.ID)
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// ClearEntries clears all "entries" edges to the Entry entity.
func (au *AccountUpdate) ClearEntries() *AccountUpdate {
	au.mutation.ClearEntries()
	return au
}

// RemoveEntryIDs removes the "entries" edge to Entry entities by IDs.
func (au *AccountUpdate) RemoveEntryIDs(ids ...int) *AccountUpdate {
	au.mutation.RemoveEntryIDs(ids...)
	return au
}

// RemoveEntries removes "entries" edges to Entry entities.
func (au *AccountUpdate) RemoveEntries(e ...*Entry) *AccountUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return au.RemoveEntryIDs(ids...)
}

// ClearFromTransfers clears all "from_transfers" edges to the Transfer entity.
func (au *AccountUpdate) ClearFromTransfers() *AccountUpdate {
	au.mutation.ClearFromTransfers()
	return au
}

// RemoveFromTransferIDs removes the "from_transfers" edge to Transfer entities by IDs.
func (au *AccountUpdate) RemoveFromTransferIDs(ids ...int) *AccountUpdate {
	au.mutation.RemoveFromTransferIDs(ids...)
	return au
}

// RemoveFromTransfers removes "from_transfers" edges to Transfer entities.
func (au *AccountUpdate) RemoveFromTransfers(t ...*Transfer) *AccountUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.RemoveFromTransferIDs(ids...)
}

// ClearToTransfers clears all "to_transfers" edges to the Transfer entity.
func (au *AccountUpdate) ClearToTransfers() *AccountUpdate {
	au.mutation.ClearToTransfers()
	return au
}

// RemoveToTransferIDs removes the "to_transfers" edge to Transfer entities by IDs.
func (au *AccountUpdate) RemoveToTransferIDs(ids ...int) *AccountUpdate {
	au.mutation.RemoveToTransferIDs(ids...)
	return au
}

// RemoveToTransfers removes "to_transfers" edges to Transfer entities.
func (au *AccountUpdate) RemoveToTransfers(t ...*Transfer) *AccountUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.RemoveToTransferIDs(ids...)
}

// ClearUsers clears the "users" edge to the User entity.
func (au *AccountUpdate) ClearUsers() *AccountUpdate {
	au.mutation.ClearUsers()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: account.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Balance(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldBalance,
		})
	}
	if value, ok := au.mutation.AddedBalance(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldBalance,
		})
	}
	if value, ok := au.mutation.Currency(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldCurrency,
		})
	}
	if au.mutation.EntriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.EntriesTable,
			Columns: []string{account.EntriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: entry.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedEntriesIDs(); len(nodes) > 0 && !au.mutation.EntriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.EntriesTable,
			Columns: []string{account.EntriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: entry.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.EntriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.EntriesTable,
			Columns: []string{account.EntriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: entry.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.FromTransfersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.FromTransfersTable,
			Columns: []string{account.FromTransfersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transfer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedFromTransfersIDs(); len(nodes) > 0 && !au.mutation.FromTransfersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.FromTransfersTable,
			Columns: []string{account.FromTransfersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transfer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.FromTransfersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.FromTransfersTable,
			Columns: []string{account.FromTransfersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transfer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.ToTransfersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.ToTransfersTable,
			Columns: []string{account.ToTransfersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transfer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedToTransfersIDs(); len(nodes) > 0 && !au.mutation.ToTransfersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.ToTransfersTable,
			Columns: []string{account.ToTransfersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transfer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.ToTransfersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.ToTransfersTable,
			Columns: []string{account.ToTransfersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transfer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   account.UsersTable,
			Columns: []string{account.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   account.UsersTable,
			Columns: []string{account.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccountMutation
}

// SetOwner sets the "owner" field.
func (auo *AccountUpdateOne) SetOwner(s string) *AccountUpdateOne {
	auo.mutation.SetOwner(s)
	return auo
}

// SetNillableOwner sets the "owner" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableOwner(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetOwner(*s)
	}
	return auo
}

// ClearOwner clears the value of the "owner" field.
func (auo *AccountUpdateOne) ClearOwner() *AccountUpdateOne {
	auo.mutation.ClearOwner()
	return auo
}

// SetBalance sets the "balance" field.
func (auo *AccountUpdateOne) SetBalance(i int) *AccountUpdateOne {
	auo.mutation.ResetBalance()
	auo.mutation.SetBalance(i)
	return auo
}

// AddBalance adds i to the "balance" field.
func (auo *AccountUpdateOne) AddBalance(i int) *AccountUpdateOne {
	auo.mutation.AddBalance(i)
	return auo
}

// SetCurrency sets the "currency" field.
func (auo *AccountUpdateOne) SetCurrency(s string) *AccountUpdateOne {
	auo.mutation.SetCurrency(s)
	return auo
}

// AddEntryIDs adds the "entries" edge to the Entry entity by IDs.
func (auo *AccountUpdateOne) AddEntryIDs(ids ...int) *AccountUpdateOne {
	auo.mutation.AddEntryIDs(ids...)
	return auo
}

// AddEntries adds the "entries" edges to the Entry entity.
func (auo *AccountUpdateOne) AddEntries(e ...*Entry) *AccountUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return auo.AddEntryIDs(ids...)
}

// AddFromTransferIDs adds the "from_transfers" edge to the Transfer entity by IDs.
func (auo *AccountUpdateOne) AddFromTransferIDs(ids ...int) *AccountUpdateOne {
	auo.mutation.AddFromTransferIDs(ids...)
	return auo
}

// AddFromTransfers adds the "from_transfers" edges to the Transfer entity.
func (auo *AccountUpdateOne) AddFromTransfers(t ...*Transfer) *AccountUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.AddFromTransferIDs(ids...)
}

// AddToTransferIDs adds the "to_transfers" edge to the Transfer entity by IDs.
func (auo *AccountUpdateOne) AddToTransferIDs(ids ...int) *AccountUpdateOne {
	auo.mutation.AddToTransferIDs(ids...)
	return auo
}

// AddToTransfers adds the "to_transfers" edges to the Transfer entity.
func (auo *AccountUpdateOne) AddToTransfers(t ...*Transfer) *AccountUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.AddToTransferIDs(ids...)
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (auo *AccountUpdateOne) SetUsersID(id string) *AccountUpdateOne {
	auo.mutation.SetUsersID(id)
	return auo
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableUsersID(id *string) *AccountUpdateOne {
	if id != nil {
		auo = auo.SetUsersID(*id)
	}
	return auo
}

// SetUsers sets the "users" edge to the User entity.
func (auo *AccountUpdateOne) SetUsers(u *User) *AccountUpdateOne {
	return auo.SetUsersID(u.ID)
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// ClearEntries clears all "entries" edges to the Entry entity.
func (auo *AccountUpdateOne) ClearEntries() *AccountUpdateOne {
	auo.mutation.ClearEntries()
	return auo
}

// RemoveEntryIDs removes the "entries" edge to Entry entities by IDs.
func (auo *AccountUpdateOne) RemoveEntryIDs(ids ...int) *AccountUpdateOne {
	auo.mutation.RemoveEntryIDs(ids...)
	return auo
}

// RemoveEntries removes "entries" edges to Entry entities.
func (auo *AccountUpdateOne) RemoveEntries(e ...*Entry) *AccountUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return auo.RemoveEntryIDs(ids...)
}

// ClearFromTransfers clears all "from_transfers" edges to the Transfer entity.
func (auo *AccountUpdateOne) ClearFromTransfers() *AccountUpdateOne {
	auo.mutation.ClearFromTransfers()
	return auo
}

// RemoveFromTransferIDs removes the "from_transfers" edge to Transfer entities by IDs.
func (auo *AccountUpdateOne) RemoveFromTransferIDs(ids ...int) *AccountUpdateOne {
	auo.mutation.RemoveFromTransferIDs(ids...)
	return auo
}

// RemoveFromTransfers removes "from_transfers" edges to Transfer entities.
func (auo *AccountUpdateOne) RemoveFromTransfers(t ...*Transfer) *AccountUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.RemoveFromTransferIDs(ids...)
}

// ClearToTransfers clears all "to_transfers" edges to the Transfer entity.
func (auo *AccountUpdateOne) ClearToTransfers() *AccountUpdateOne {
	auo.mutation.ClearToTransfers()
	return auo
}

// RemoveToTransferIDs removes the "to_transfers" edge to Transfer entities by IDs.
func (auo *AccountUpdateOne) RemoveToTransferIDs(ids ...int) *AccountUpdateOne {
	auo.mutation.RemoveToTransferIDs(ids...)
	return auo
}

// RemoveToTransfers removes "to_transfers" edges to Transfer entities.
func (auo *AccountUpdateOne) RemoveToTransfers(t ...*Transfer) *AccountUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.RemoveToTransferIDs(ids...)
}

// ClearUsers clears the "users" edge to the User entity.
func (auo *AccountUpdateOne) ClearUsers() *AccountUpdateOne {
	auo.mutation.ClearUsers()
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccountUpdateOne) Select(field string, fields ...string) *AccountUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Account entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	var (
		err  error
		node *Account
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: account.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Account.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for _, f := range fields {
			if !account.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Balance(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldBalance,
		})
	}
	if value, ok := auo.mutation.AddedBalance(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldBalance,
		})
	}
	if value, ok := auo.mutation.Currency(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldCurrency,
		})
	}
	if auo.mutation.EntriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.EntriesTable,
			Columns: []string{account.EntriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: entry.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedEntriesIDs(); len(nodes) > 0 && !auo.mutation.EntriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.EntriesTable,
			Columns: []string{account.EntriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: entry.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.EntriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.EntriesTable,
			Columns: []string{account.EntriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: entry.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.FromTransfersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.FromTransfersTable,
			Columns: []string{account.FromTransfersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transfer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedFromTransfersIDs(); len(nodes) > 0 && !auo.mutation.FromTransfersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.FromTransfersTable,
			Columns: []string{account.FromTransfersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transfer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.FromTransfersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.FromTransfersTable,
			Columns: []string{account.FromTransfersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transfer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.ToTransfersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.ToTransfersTable,
			Columns: []string{account.ToTransfersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transfer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedToTransfersIDs(); len(nodes) > 0 && !auo.mutation.ToTransfersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.ToTransfersTable,
			Columns: []string{account.ToTransfersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transfer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.ToTransfersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.ToTransfersTable,
			Columns: []string{account.ToTransfersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transfer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   account.UsersTable,
			Columns: []string{account.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   account.UsersTable,
			Columns: []string{account.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
