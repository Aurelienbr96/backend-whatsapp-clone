// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ContactsColumns holds the columns for the "contacts" table.
	ContactsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "contact_user_id", Type: field.TypeUUID},
		{Name: "owner_id", Type: field.TypeUUID},
	}
	// ContactsTable holds the schema information for the "contacts" table.
	ContactsTable = &schema.Table{
		Name:       "contacts",
		Columns:    ContactsColumns,
		PrimaryKey: []*schema.Column{ContactsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "contacts_users_contact_user",
				Columns:    []*schema.Column{ContactsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "contacts_users_contacts",
				Columns:    []*schema.Column{ContactsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "contact_owner_id_contact_user_id",
				Unique:  true,
				Columns: []*schema.Column{ContactsColumns[3], ContactsColumns[2]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "phone_number", Type: field.TypeString, Unique: true},
		{Name: "avatar", Type: field.TypeString, Nullable: true},
		{Name: "username", Type: field.TypeString, Nullable: true},
		{Name: "is_verified", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ContactsTable,
		UsersTable,
	}
)

func init() {
	ContactsTable.ForeignKeys[0].RefTable = UsersTable
	ContactsTable.ForeignKeys[1].RefTable = UsersTable
}
