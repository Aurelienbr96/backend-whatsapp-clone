// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"example.com/boiletplate/ent/contact"
	"example.com/boiletplate/ent/predicate"
	"example.com/boiletplate/ent/user"
	"github.com/google/uuid"
)

// ContactUpdate is the builder for updating Contact entities.
type ContactUpdate struct {
	config
	hooks    []Hook
	mutation *ContactMutation
}

// Where appends a list predicates to the ContactUpdate builder.
func (cu *ContactUpdate) Where(ps ...predicate.Contact) *ContactUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetOwnerID sets the "owner_id" field.
func (cu *ContactUpdate) SetOwnerID(u uuid.UUID) *ContactUpdate {
	cu.mutation.SetOwnerID(u)
	return cu
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (cu *ContactUpdate) SetNillableOwnerID(u *uuid.UUID) *ContactUpdate {
	if u != nil {
		cu.SetOwnerID(*u)
	}
	return cu
}

// SetContactUserID sets the "contact_user_id" field.
func (cu *ContactUpdate) SetContactUserID(u uuid.UUID) *ContactUpdate {
	cu.mutation.SetContactUserID(u)
	return cu
}

// SetNillableContactUserID sets the "contact_user_id" field if the given value is not nil.
func (cu *ContactUpdate) SetNillableContactUserID(u *uuid.UUID) *ContactUpdate {
	if u != nil {
		cu.SetContactUserID(*u)
	}
	return cu
}

// SetOwner sets the "owner" edge to the User entity.
func (cu *ContactUpdate) SetOwner(u *User) *ContactUpdate {
	return cu.SetOwnerID(u.ID)
}

// SetContactUser sets the "contact_user" edge to the User entity.
func (cu *ContactUpdate) SetContactUser(u *User) *ContactUpdate {
	return cu.SetContactUserID(u.ID)
}

// Mutation returns the ContactMutation object of the builder.
func (cu *ContactUpdate) Mutation() *ContactMutation {
	return cu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (cu *ContactUpdate) ClearOwner() *ContactUpdate {
	cu.mutation.ClearOwner()
	return cu
}

// ClearContactUser clears the "contact_user" edge to the User entity.
func (cu *ContactUpdate) ClearContactUser() *ContactUpdate {
	cu.mutation.ClearContactUser()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ContactUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ContactUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ContactUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ContactUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ContactUpdate) check() error {
	if cu.mutation.OwnerCleared() && len(cu.mutation.OwnerIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Contact.owner"`)
	}
	if cu.mutation.ContactUserCleared() && len(cu.mutation.ContactUserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Contact.contact_user"`)
	}
	return nil
}

func (cu *ContactUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(contact.Table, contact.Columns, sqlgraph.NewFieldSpec(contact.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contact.OwnerTable,
			Columns: []string{contact.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contact.OwnerTable,
			Columns: []string{contact.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ContactUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   contact.ContactUserTable,
			Columns: []string{contact.ContactUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ContactUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   contact.ContactUserTable,
			Columns: []string{contact.ContactUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contact.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ContactUpdateOne is the builder for updating a single Contact entity.
type ContactUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ContactMutation
}

// SetOwnerID sets the "owner_id" field.
func (cuo *ContactUpdateOne) SetOwnerID(u uuid.UUID) *ContactUpdateOne {
	cuo.mutation.SetOwnerID(u)
	return cuo
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (cuo *ContactUpdateOne) SetNillableOwnerID(u *uuid.UUID) *ContactUpdateOne {
	if u != nil {
		cuo.SetOwnerID(*u)
	}
	return cuo
}

// SetContactUserID sets the "contact_user_id" field.
func (cuo *ContactUpdateOne) SetContactUserID(u uuid.UUID) *ContactUpdateOne {
	cuo.mutation.SetContactUserID(u)
	return cuo
}

// SetNillableContactUserID sets the "contact_user_id" field if the given value is not nil.
func (cuo *ContactUpdateOne) SetNillableContactUserID(u *uuid.UUID) *ContactUpdateOne {
	if u != nil {
		cuo.SetContactUserID(*u)
	}
	return cuo
}

// SetOwner sets the "owner" edge to the User entity.
func (cuo *ContactUpdateOne) SetOwner(u *User) *ContactUpdateOne {
	return cuo.SetOwnerID(u.ID)
}

// SetContactUser sets the "contact_user" edge to the User entity.
func (cuo *ContactUpdateOne) SetContactUser(u *User) *ContactUpdateOne {
	return cuo.SetContactUserID(u.ID)
}

// Mutation returns the ContactMutation object of the builder.
func (cuo *ContactUpdateOne) Mutation() *ContactMutation {
	return cuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (cuo *ContactUpdateOne) ClearOwner() *ContactUpdateOne {
	cuo.mutation.ClearOwner()
	return cuo
}

// ClearContactUser clears the "contact_user" edge to the User entity.
func (cuo *ContactUpdateOne) ClearContactUser() *ContactUpdateOne {
	cuo.mutation.ClearContactUser()
	return cuo
}

// Where appends a list predicates to the ContactUpdate builder.
func (cuo *ContactUpdateOne) Where(ps ...predicate.Contact) *ContactUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ContactUpdateOne) Select(field string, fields ...string) *ContactUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Contact entity.
func (cuo *ContactUpdateOne) Save(ctx context.Context) (*Contact, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ContactUpdateOne) SaveX(ctx context.Context) *Contact {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ContactUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ContactUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ContactUpdateOne) check() error {
	if cuo.mutation.OwnerCleared() && len(cuo.mutation.OwnerIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Contact.owner"`)
	}
	if cuo.mutation.ContactUserCleared() && len(cuo.mutation.ContactUserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Contact.contact_user"`)
	}
	return nil
}

func (cuo *ContactUpdateOne) sqlSave(ctx context.Context) (_node *Contact, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(contact.Table, contact.Columns, sqlgraph.NewFieldSpec(contact.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Contact.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contact.FieldID)
		for _, f := range fields {
			if !contact.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != contact.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contact.OwnerTable,
			Columns: []string{contact.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contact.OwnerTable,
			Columns: []string{contact.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ContactUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   contact.ContactUserTable,
			Columns: []string{contact.ContactUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ContactUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   contact.ContactUserTable,
			Columns: []string{contact.ContactUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Contact{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contact.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
