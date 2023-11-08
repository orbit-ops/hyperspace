// Code generated by ent, DO NOT EDIT.

package audit

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the audit type in the database.
	Label = "audit"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAction holds the string denoting the action field in the database.
	FieldAction = "action"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the audit in the database.
	Table = "audits"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "audits"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_audit"
)

// Columns holds all SQL columns for audit fields.
var Columns = []string{
	FieldID,
	FieldAction,
	FieldAuthor,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "audits"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_audit",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// ActionValidator is a validator for the "action" field. It is called by the builders before save.
	ActionValidator func(string) error
	// AuthorValidator is a validator for the "author" field. It is called by the builders before save.
	AuthorValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the Audit queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAction orders the results by the action field.
func ByAction(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAction, opts...).ToFunc()
}

// ByAuthor orders the results by the author field.
func ByAuthor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthor, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
