// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tiagoposse/connect/ent/apikey"
	"github.com/tiagoposse/connect/ent/device"
	"github.com/tiagoposse/connect/ent/group"
	"github.com/tiagoposse/connect/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetFirstname sets the "firstname" field.
func (uc *UserCreate) SetFirstname(s string) *UserCreate {
	uc.mutation.SetFirstname(s)
	return uc
}

// SetLastname sets the "lastname" field.
func (uc *UserCreate) SetLastname(s string) *UserCreate {
	uc.mutation.SetLastname(s)
	return uc
}

// SetProvider sets the "provider" field.
func (uc *UserCreate) SetProvider(s string) *UserCreate {
	uc.mutation.SetProvider(s)
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uc *UserCreate) SetNillablePassword(s *string) *UserCreate {
	if s != nil {
		uc.SetPassword(*s)
	}
	return uc
}

// SetSalt sets the "salt" field.
func (uc *UserCreate) SetSalt(s string) *UserCreate {
	uc.mutation.SetSalt(s)
	return uc
}

// SetNillableSalt sets the "salt" field if the given value is not nil.
func (uc *UserCreate) SetNillableSalt(s *string) *UserCreate {
	if s != nil {
		uc.SetSalt(*s)
	}
	return uc
}

// SetPhotoURL sets the "photo_url" field.
func (uc *UserCreate) SetPhotoURL(s string) *UserCreate {
	uc.mutation.SetPhotoURL(s)
	return uc
}

// SetNillablePhotoURL sets the "photo_url" field if the given value is not nil.
func (uc *UserCreate) SetNillablePhotoURL(s *string) *UserCreate {
	if s != nil {
		uc.SetPhotoURL(*s)
	}
	return uc
}

// SetDisabled sets the "disabled" field.
func (uc *UserCreate) SetDisabled(b bool) *UserCreate {
	uc.mutation.SetDisabled(b)
	return uc
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (uc *UserCreate) SetNillableDisabled(b *bool) *UserCreate {
	if b != nil {
		uc.SetDisabled(*b)
	}
	return uc
}

// SetDisabledReason sets the "disabled_reason" field.
func (uc *UserCreate) SetDisabledReason(s string) *UserCreate {
	uc.mutation.SetDisabledReason(s)
	return uc
}

// SetNillableDisabledReason sets the "disabled_reason" field if the given value is not nil.
func (uc *UserCreate) SetNillableDisabledReason(s *string) *UserCreate {
	if s != nil {
		uc.SetDisabledReason(*s)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(s string) *UserCreate {
	uc.mutation.SetID(s)
	return uc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uc *UserCreate) SetNillableID(s *string) *UserCreate {
	if s != nil {
		uc.SetID(*s)
	}
	return uc
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (uc *UserCreate) SetGroupID(id string) *UserCreate {
	uc.mutation.SetGroupID(id)
	return uc
}

// SetGroup sets the "group" edge to the Group entity.
func (uc *UserCreate) SetGroup(g *Group) *UserCreate {
	return uc.SetGroupID(g.ID)
}

// AddDeviceIDs adds the "devices" edge to the Device entity by IDs.
func (uc *UserCreate) AddDeviceIDs(ids ...int) *UserCreate {
	uc.mutation.AddDeviceIDs(ids...)
	return uc
}

// AddDevices adds the "devices" edges to the Device entity.
func (uc *UserCreate) AddDevices(d ...*Device) *UserCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uc.AddDeviceIDs(ids...)
}

// AddKeyIDs adds the "keys" edge to the ApiKey entity by IDs.
func (uc *UserCreate) AddKeyIDs(ids ...int) *UserCreate {
	uc.mutation.AddKeyIDs(ids...)
	return uc
}

// AddKeys adds the "keys" edges to the ApiKey entity.
func (uc *UserCreate) AddKeys(a ...*ApiKey) *UserCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uc.AddKeyIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.Disabled(); !ok {
		v := user.DefaultDisabled
		uc.mutation.SetDisabled(v)
	}
	if _, ok := uc.mutation.ID(); !ok {
		v := user.DefaultID()
		uc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "User.email"`)}
	}
	if v, ok := uc.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Firstname(); !ok {
		return &ValidationError{Name: "firstname", err: errors.New(`ent: missing required field "User.firstname"`)}
	}
	if v, ok := uc.mutation.Firstname(); ok {
		if err := user.FirstnameValidator(v); err != nil {
			return &ValidationError{Name: "firstname", err: fmt.Errorf(`ent: validator failed for field "User.firstname": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Lastname(); !ok {
		return &ValidationError{Name: "lastname", err: errors.New(`ent: missing required field "User.lastname"`)}
	}
	if v, ok := uc.mutation.Lastname(); ok {
		if err := user.LastnameValidator(v); err != nil {
			return &ValidationError{Name: "lastname", err: fmt.Errorf(`ent: validator failed for field "User.lastname": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Provider(); !ok {
		return &ValidationError{Name: "provider", err: errors.New(`ent: missing required field "User.provider"`)}
	}
	if v, ok := uc.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if v, ok := uc.mutation.Salt(); ok {
		if err := user.SaltValidator(v); err != nil {
			return &ValidationError{Name: "salt", err: fmt.Errorf(`ent: validator failed for field "User.salt": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Disabled(); !ok {
		return &ValidationError{Name: "disabled", err: errors.New(`ent: missing required field "User.disabled"`)}
	}
	if _, ok := uc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group", err: errors.New(`ent: missing required edge "User.group"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected User.ID type: %T", _spec.ID.Value)
		}
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	)
	_spec.OnConflict = uc.conflict
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.Firstname(); ok {
		_spec.SetField(user.FieldFirstname, field.TypeString, value)
		_node.Firstname = value
	}
	if value, ok := uc.mutation.Lastname(); ok {
		_spec.SetField(user.FieldLastname, field.TypeString, value)
		_node.Lastname = value
	}
	if value, ok := uc.mutation.Provider(); ok {
		_spec.SetField(user.FieldProvider, field.TypeString, value)
		_node.Provider = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := uc.mutation.Salt(); ok {
		_spec.SetField(user.FieldSalt, field.TypeString, value)
		_node.Salt = value
	}
	if value, ok := uc.mutation.PhotoURL(); ok {
		_spec.SetField(user.FieldPhotoURL, field.TypeString, value)
		_node.PhotoURL = value
	}
	if value, ok := uc.mutation.Disabled(); ok {
		_spec.SetField(user.FieldDisabled, field.TypeBool, value)
		_node.Disabled = value
	}
	if value, ok := uc.mutation.DisabledReason(); ok {
		_spec.SetField(user.FieldDisabledReason, field.TypeString, value)
		_node.DisabledReason = value
	}
	if nodes := uc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.GroupTable,
			Columns: []string{user.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.group_users = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.DevicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DevicesTable,
			Columns: []string{user.DevicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.KeysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.KeysTable,
			Columns: []string{user.KeysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.Create().
//		SetEmail(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetEmail(v+v).
//		}).
//		Exec(ctx)
func (uc *UserCreate) OnConflict(opts ...sql.ConflictOption) *UserUpsertOne {
	uc.conflict = opts
	return &UserUpsertOne{
		create: uc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uc *UserCreate) OnConflictColumns(columns ...string) *UserUpsertOne {
	uc.conflict = append(uc.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertOne{
		create: uc,
	}
}

type (
	// UserUpsertOne is the builder for "upsert"-ing
	//  one User node.
	UserUpsertOne struct {
		create *UserCreate
	}

	// UserUpsert is the "OnConflict" setter.
	UserUpsert struct {
		*sql.UpdateSet
	}
)

// SetEmail sets the "email" field.
func (u *UserUpsert) SetEmail(v string) *UserUpsert {
	u.Set(user.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsert) UpdateEmail() *UserUpsert {
	u.SetExcluded(user.FieldEmail)
	return u
}

// SetFirstname sets the "firstname" field.
func (u *UserUpsert) SetFirstname(v string) *UserUpsert {
	u.Set(user.FieldFirstname, v)
	return u
}

// UpdateFirstname sets the "firstname" field to the value that was provided on create.
func (u *UserUpsert) UpdateFirstname() *UserUpsert {
	u.SetExcluded(user.FieldFirstname)
	return u
}

// SetLastname sets the "lastname" field.
func (u *UserUpsert) SetLastname(v string) *UserUpsert {
	u.Set(user.FieldLastname, v)
	return u
}

// UpdateLastname sets the "lastname" field to the value that was provided on create.
func (u *UserUpsert) UpdateLastname() *UserUpsert {
	u.SetExcluded(user.FieldLastname)
	return u
}

// SetPassword sets the "password" field.
func (u *UserUpsert) SetPassword(v string) *UserUpsert {
	u.Set(user.FieldPassword, v)
	return u
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *UserUpsert) UpdatePassword() *UserUpsert {
	u.SetExcluded(user.FieldPassword)
	return u
}

// ClearPassword clears the value of the "password" field.
func (u *UserUpsert) ClearPassword() *UserUpsert {
	u.SetNull(user.FieldPassword)
	return u
}

// SetSalt sets the "salt" field.
func (u *UserUpsert) SetSalt(v string) *UserUpsert {
	u.Set(user.FieldSalt, v)
	return u
}

// UpdateSalt sets the "salt" field to the value that was provided on create.
func (u *UserUpsert) UpdateSalt() *UserUpsert {
	u.SetExcluded(user.FieldSalt)
	return u
}

// ClearSalt clears the value of the "salt" field.
func (u *UserUpsert) ClearSalt() *UserUpsert {
	u.SetNull(user.FieldSalt)
	return u
}

// SetPhotoURL sets the "photo_url" field.
func (u *UserUpsert) SetPhotoURL(v string) *UserUpsert {
	u.Set(user.FieldPhotoURL, v)
	return u
}

// UpdatePhotoURL sets the "photo_url" field to the value that was provided on create.
func (u *UserUpsert) UpdatePhotoURL() *UserUpsert {
	u.SetExcluded(user.FieldPhotoURL)
	return u
}

// ClearPhotoURL clears the value of the "photo_url" field.
func (u *UserUpsert) ClearPhotoURL() *UserUpsert {
	u.SetNull(user.FieldPhotoURL)
	return u
}

// SetDisabled sets the "disabled" field.
func (u *UserUpsert) SetDisabled(v bool) *UserUpsert {
	u.Set(user.FieldDisabled, v)
	return u
}

// UpdateDisabled sets the "disabled" field to the value that was provided on create.
func (u *UserUpsert) UpdateDisabled() *UserUpsert {
	u.SetExcluded(user.FieldDisabled)
	return u
}

// SetDisabledReason sets the "disabled_reason" field.
func (u *UserUpsert) SetDisabledReason(v string) *UserUpsert {
	u.Set(user.FieldDisabledReason, v)
	return u
}

// UpdateDisabledReason sets the "disabled_reason" field to the value that was provided on create.
func (u *UserUpsert) UpdateDisabledReason() *UserUpsert {
	u.SetExcluded(user.FieldDisabledReason)
	return u
}

// ClearDisabledReason clears the value of the "disabled_reason" field.
func (u *UserUpsert) ClearDisabledReason() *UserUpsert {
	u.SetNull(user.FieldDisabledReason)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(user.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserUpsertOne) UpdateNewValues() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(user.FieldID)
		}
		if _, exists := u.create.mutation.Provider(); exists {
			s.SetIgnore(user.FieldProvider)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserUpsertOne) Ignore() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertOne) DoNothing() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreate.OnConflict
// documentation for more info.
func (u *UserUpsertOne) Update(set func(*UserUpsert)) *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetEmail sets the "email" field.
func (u *UserUpsertOne) SetEmail(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateEmail() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateEmail()
	})
}

// SetFirstname sets the "firstname" field.
func (u *UserUpsertOne) SetFirstname(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetFirstname(v)
	})
}

// UpdateFirstname sets the "firstname" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateFirstname() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateFirstname()
	})
}

// SetLastname sets the "lastname" field.
func (u *UserUpsertOne) SetLastname(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetLastname(v)
	})
}

// UpdateLastname sets the "lastname" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateLastname() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateLastname()
	})
}

// SetPassword sets the "password" field.
func (u *UserUpsertOne) SetPassword(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *UserUpsertOne) UpdatePassword() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdatePassword()
	})
}

// ClearPassword clears the value of the "password" field.
func (u *UserUpsertOne) ClearPassword() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearPassword()
	})
}

// SetSalt sets the "salt" field.
func (u *UserUpsertOne) SetSalt(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetSalt(v)
	})
}

// UpdateSalt sets the "salt" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateSalt() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateSalt()
	})
}

// ClearSalt clears the value of the "salt" field.
func (u *UserUpsertOne) ClearSalt() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearSalt()
	})
}

// SetPhotoURL sets the "photo_url" field.
func (u *UserUpsertOne) SetPhotoURL(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetPhotoURL(v)
	})
}

// UpdatePhotoURL sets the "photo_url" field to the value that was provided on create.
func (u *UserUpsertOne) UpdatePhotoURL() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdatePhotoURL()
	})
}

// ClearPhotoURL clears the value of the "photo_url" field.
func (u *UserUpsertOne) ClearPhotoURL() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearPhotoURL()
	})
}

// SetDisabled sets the "disabled" field.
func (u *UserUpsertOne) SetDisabled(v bool) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetDisabled(v)
	})
}

// UpdateDisabled sets the "disabled" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateDisabled() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateDisabled()
	})
}

// SetDisabledReason sets the "disabled_reason" field.
func (u *UserUpsertOne) SetDisabledReason(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetDisabledReason(v)
	})
}

// UpdateDisabledReason sets the "disabled_reason" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateDisabledReason() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateDisabledReason()
	})
}

// ClearDisabledReason clears the value of the "disabled_reason" field.
func (u *UserUpsertOne) ClearDisabledReason() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearDisabledReason()
	})
}

// Exec executes the query.
func (u *UserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: UserUpsertOne.ID is not supported by MySQL driver. Use UserUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
	conflict []sql.ConflictOption
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetEmail(v+v).
//		}).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserUpsertBulk {
	ucb.conflict = opts
	return &UserUpsertBulk{
		create: ucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflictColumns(columns ...string) *UserUpsertBulk {
	ucb.conflict = append(ucb.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertBulk{
		create: ucb,
	}
}

// UserUpsertBulk is the builder for "upsert"-ing
// a bulk of User nodes.
type UserUpsertBulk struct {
	create *UserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(user.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserUpsertBulk) UpdateNewValues() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(user.FieldID)
			}
			if _, exists := b.mutation.Provider(); exists {
				s.SetIgnore(user.FieldProvider)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserUpsertBulk) Ignore() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertBulk) DoNothing() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreateBulk.OnConflict
// documentation for more info.
func (u *UserUpsertBulk) Update(set func(*UserUpsert)) *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetEmail sets the "email" field.
func (u *UserUpsertBulk) SetEmail(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateEmail() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateEmail()
	})
}

// SetFirstname sets the "firstname" field.
func (u *UserUpsertBulk) SetFirstname(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetFirstname(v)
	})
}

// UpdateFirstname sets the "firstname" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateFirstname() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateFirstname()
	})
}

// SetLastname sets the "lastname" field.
func (u *UserUpsertBulk) SetLastname(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetLastname(v)
	})
}

// UpdateLastname sets the "lastname" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateLastname() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateLastname()
	})
}

// SetPassword sets the "password" field.
func (u *UserUpsertBulk) SetPassword(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdatePassword() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdatePassword()
	})
}

// ClearPassword clears the value of the "password" field.
func (u *UserUpsertBulk) ClearPassword() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearPassword()
	})
}

// SetSalt sets the "salt" field.
func (u *UserUpsertBulk) SetSalt(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetSalt(v)
	})
}

// UpdateSalt sets the "salt" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateSalt() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateSalt()
	})
}

// ClearSalt clears the value of the "salt" field.
func (u *UserUpsertBulk) ClearSalt() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearSalt()
	})
}

// SetPhotoURL sets the "photo_url" field.
func (u *UserUpsertBulk) SetPhotoURL(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetPhotoURL(v)
	})
}

// UpdatePhotoURL sets the "photo_url" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdatePhotoURL() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdatePhotoURL()
	})
}

// ClearPhotoURL clears the value of the "photo_url" field.
func (u *UserUpsertBulk) ClearPhotoURL() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearPhotoURL()
	})
}

// SetDisabled sets the "disabled" field.
func (u *UserUpsertBulk) SetDisabled(v bool) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetDisabled(v)
	})
}

// UpdateDisabled sets the "disabled" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateDisabled() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateDisabled()
	})
}

// SetDisabledReason sets the "disabled_reason" field.
func (u *UserUpsertBulk) SetDisabledReason(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetDisabledReason(v)
	})
}

// UpdateDisabledReason sets the "disabled_reason" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateDisabledReason() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateDisabledReason()
	})
}

// ClearDisabledReason clears the value of the "disabled_reason" field.
func (u *UserUpsertBulk) ClearDisabledReason() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearDisabledReason()
	})
}

// Exec executes the query.
func (u *UserUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
