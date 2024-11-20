// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"example.com/boiletplate/ent/contact"
	"example.com/boiletplate/ent/user"
	"github.com/google/uuid"
)

// Contact is the model entity for the Contact schema.
type Contact struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID uuid.UUID `json:"owner_id,omitempty"`
	// ContactUserID holds the value of the "contact_user_id" field.
	ContactUserID uuid.UUID `json:"contact_user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ContactQuery when eager-loading is set.
	Edges        ContactEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ContactEdges holds the relations/edges for other nodes in the graph.
type ContactEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// ContactUser holds the value of the contact_user edge.
	ContactUser *User `json:"contact_user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ContactEdges) OwnerOrErr() (*User, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// ContactUserOrErr returns the ContactUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ContactEdges) ContactUserOrErr() (*User, error) {
	if e.ContactUser != nil {
		return e.ContactUser, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "contact_user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Contact) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case contact.FieldID, contact.FieldOwnerID, contact.FieldContactUserID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Contact fields.
func (c *Contact) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case contact.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case contact.FieldOwnerID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value != nil {
				c.OwnerID = *value
			}
		case contact.FieldContactUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field contact_user_id", values[i])
			} else if value != nil {
				c.ContactUserID = *value
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Contact.
// This includes values selected through modifiers, order, etc.
func (c *Contact) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Contact entity.
func (c *Contact) QueryOwner() *UserQuery {
	return NewContactClient(c.config).QueryOwner(c)
}

// QueryContactUser queries the "contact_user" edge of the Contact entity.
func (c *Contact) QueryContactUser() *UserQuery {
	return NewContactClient(c.config).QueryContactUser(c)
}

// Update returns a builder for updating this Contact.
// Note that you need to call Contact.Unwrap() before calling this method if this Contact
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Contact) Update() *ContactUpdateOne {
	return NewContactClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Contact entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Contact) Unwrap() *Contact {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Contact is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Contact) String() string {
	var builder strings.Builder
	builder.WriteString("Contact(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("owner_id=")
	builder.WriteString(fmt.Sprintf("%v", c.OwnerID))
	builder.WriteString(", ")
	builder.WriteString("contact_user_id=")
	builder.WriteString(fmt.Sprintf("%v", c.ContactUserID))
	builder.WriteByte(')')
	return builder.String()
}

// Contacts is a parsable slice of Contact.
type Contacts []*Contact
