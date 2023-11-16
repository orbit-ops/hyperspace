// Code generated by ent, DO NOT EDIT.

package apikey

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/tiagoposse/go-auth/controller"
)

const (
	// Label holds the string label denoting the apikey type in the database.
	Label = "api_key"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldScopes holds the string denoting the scopes field in the database.
	FieldScopes = "scopes"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the apikey in the database.
	Table = "api_keys"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "api_keys"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for apikey fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldKey,
	FieldScopes,
	FieldUserID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/tiagoposse/connect/ent/runtime"
var (
	Hooks [1]ent.Hook
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultScopes holds the default value on creation for the "scopes" field.
	DefaultScopes controller.Scopes
)

// OrderOption defines the ordering options for the ApiKey queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByKey orders the results by the key field.
func ByKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKey, opts...).ToFunc()
}

// ByScopes orders the results by the scopes field.
func ByScopes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScopes, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
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