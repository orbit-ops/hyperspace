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
	"github.com/google/uuid"
	"github.com/tiagoposse/connect/ent/device"
	"github.com/tiagoposse/connect/ent/user"
	"github.com/tiagoposse/connect/internal/types"
)

// DeviceCreate is the builder for creating a Device entity.
type DeviceCreate struct {
	config
	mutation *DeviceMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (dc *DeviceCreate) SetName(s string) *DeviceCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetDescription sets the "description" field.
func (dc *DeviceCreate) SetDescription(s string) *DeviceCreate {
	dc.mutation.SetDescription(s)
	return dc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableDescription(s *string) *DeviceCreate {
	if s != nil {
		dc.SetDescription(*s)
	}
	return dc
}

// SetType sets the "type" field.
func (dc *DeviceCreate) SetType(s string) *DeviceCreate {
	dc.mutation.SetType(s)
	return dc
}

// SetDNS sets the "dns" field.
func (dc *DeviceCreate) SetDNS(ts types.InetSlice) *DeviceCreate {
	dc.mutation.SetDNS(ts)
	return dc
}

// SetPublicKey sets the "public_key" field.
func (dc *DeviceCreate) SetPublicKey(s string) *DeviceCreate {
	dc.mutation.SetPublicKey(s)
	return dc
}

// SetPresharedKey sets the "preshared_key" field.
func (dc *DeviceCreate) SetPresharedKey(s string) *DeviceCreate {
	dc.mutation.SetPresharedKey(s)
	return dc
}

// SetKeepAlive sets the "keep_alive" field.
func (dc *DeviceCreate) SetKeepAlive(b bool) *DeviceCreate {
	dc.mutation.SetKeepAlive(b)
	return dc
}

// SetEndpoint sets the "endpoint" field.
func (dc *DeviceCreate) SetEndpoint(t types.Inet) *DeviceCreate {
	dc.mutation.SetEndpoint(t)
	return dc
}

// SetAllowedIps sets the "allowed_ips" field.
func (dc *DeviceCreate) SetAllowedIps(ts types.CidrSlice) *DeviceCreate {
	dc.mutation.SetAllowedIps(ts)
	return dc
}

// SetUserID sets the "user_id" field.
func (dc *DeviceCreate) SetUserID(s string) *DeviceCreate {
	dc.mutation.SetUserID(s)
	return dc
}

// SetID sets the "id" field.
func (dc *DeviceCreate) SetID(u uuid.UUID) *DeviceCreate {
	dc.mutation.SetID(u)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableID(u *uuid.UUID) *DeviceCreate {
	if u != nil {
		dc.SetID(*u)
	}
	return dc
}

// SetUser sets the "user" edge to the User entity.
func (dc *DeviceCreate) SetUser(u *User) *DeviceCreate {
	return dc.SetUserID(u.ID)
}

// Mutation returns the DeviceMutation object of the builder.
func (dc *DeviceCreate) Mutation() *DeviceMutation {
	return dc.mutation
}

// Save creates the Device in the database.
func (dc *DeviceCreate) Save(ctx context.Context) (*Device, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DeviceCreate) SaveX(ctx context.Context) *Device {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DeviceCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DeviceCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DeviceCreate) defaults() {
	if _, ok := dc.mutation.ID(); !ok {
		v := device.DefaultID()
		dc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DeviceCreate) check() error {
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Device.name"`)}
	}
	if _, ok := dc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Device.type"`)}
	}
	if _, ok := dc.mutation.DNS(); !ok {
		return &ValidationError{Name: "dns", err: errors.New(`ent: missing required field "Device.dns"`)}
	}
	if _, ok := dc.mutation.PublicKey(); !ok {
		return &ValidationError{Name: "public_key", err: errors.New(`ent: missing required field "Device.public_key"`)}
	}
	if _, ok := dc.mutation.PresharedKey(); !ok {
		return &ValidationError{Name: "preshared_key", err: errors.New(`ent: missing required field "Device.preshared_key"`)}
	}
	if _, ok := dc.mutation.KeepAlive(); !ok {
		return &ValidationError{Name: "keep_alive", err: errors.New(`ent: missing required field "Device.keep_alive"`)}
	}
	if _, ok := dc.mutation.Endpoint(); !ok {
		return &ValidationError{Name: "endpoint", err: errors.New(`ent: missing required field "Device.endpoint"`)}
	}
	if _, ok := dc.mutation.AllowedIps(); !ok {
		return &ValidationError{Name: "allowed_ips", err: errors.New(`ent: missing required field "Device.allowed_ips"`)}
	}
	if _, ok := dc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Device.user_id"`)}
	}
	if _, ok := dc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Device.user"`)}
	}
	return nil
}

func (dc *DeviceCreate) sqlSave(ctx context.Context) (*Device, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DeviceCreate) createSpec() (*Device, *sqlgraph.CreateSpec) {
	var (
		_node = &Device{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(device.Table, sqlgraph.NewFieldSpec(device.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = dc.conflict
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(device.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dc.mutation.Description(); ok {
		_spec.SetField(device.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := dc.mutation.GetType(); ok {
		_spec.SetField(device.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := dc.mutation.DNS(); ok {
		_spec.SetField(device.FieldDNS, field.TypeOther, value)
		_node.DNS = value
	}
	if value, ok := dc.mutation.PublicKey(); ok {
		_spec.SetField(device.FieldPublicKey, field.TypeString, value)
		_node.PublicKey = value
	}
	if value, ok := dc.mutation.PresharedKey(); ok {
		_spec.SetField(device.FieldPresharedKey, field.TypeString, value)
		_node.PresharedKey = value
	}
	if value, ok := dc.mutation.KeepAlive(); ok {
		_spec.SetField(device.FieldKeepAlive, field.TypeBool, value)
		_node.KeepAlive = value
	}
	if value, ok := dc.mutation.Endpoint(); ok {
		_spec.SetField(device.FieldEndpoint, field.TypeString, value)
		_node.Endpoint = value
	}
	if value, ok := dc.mutation.AllowedIps(); ok {
		_spec.SetField(device.FieldAllowedIps, field.TypeOther, value)
		_node.AllowedIps = value
	}
	if nodes := dc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   device.UserTable,
			Columns: []string{device.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Device.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeviceUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (dc *DeviceCreate) OnConflict(opts ...sql.ConflictOption) *DeviceUpsertOne {
	dc.conflict = opts
	return &DeviceUpsertOne{
		create: dc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Device.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dc *DeviceCreate) OnConflictColumns(columns ...string) *DeviceUpsertOne {
	dc.conflict = append(dc.conflict, sql.ConflictColumns(columns...))
	return &DeviceUpsertOne{
		create: dc,
	}
}

type (
	// DeviceUpsertOne is the builder for "upsert"-ing
	//  one Device node.
	DeviceUpsertOne struct {
		create *DeviceCreate
	}

	// DeviceUpsert is the "OnConflict" setter.
	DeviceUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *DeviceUpsert) SetName(v string) *DeviceUpsert {
	u.Set(device.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DeviceUpsert) UpdateName() *DeviceUpsert {
	u.SetExcluded(device.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *DeviceUpsert) SetDescription(v string) *DeviceUpsert {
	u.Set(device.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *DeviceUpsert) UpdateDescription() *DeviceUpsert {
	u.SetExcluded(device.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *DeviceUpsert) ClearDescription() *DeviceUpsert {
	u.SetNull(device.FieldDescription)
	return u
}

// SetType sets the "type" field.
func (u *DeviceUpsert) SetType(v string) *DeviceUpsert {
	u.Set(device.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *DeviceUpsert) UpdateType() *DeviceUpsert {
	u.SetExcluded(device.FieldType)
	return u
}

// SetDNS sets the "dns" field.
func (u *DeviceUpsert) SetDNS(v types.InetSlice) *DeviceUpsert {
	u.Set(device.FieldDNS, v)
	return u
}

// UpdateDNS sets the "dns" field to the value that was provided on create.
func (u *DeviceUpsert) UpdateDNS() *DeviceUpsert {
	u.SetExcluded(device.FieldDNS)
	return u
}

// SetKeepAlive sets the "keep_alive" field.
func (u *DeviceUpsert) SetKeepAlive(v bool) *DeviceUpsert {
	u.Set(device.FieldKeepAlive, v)
	return u
}

// UpdateKeepAlive sets the "keep_alive" field to the value that was provided on create.
func (u *DeviceUpsert) UpdateKeepAlive() *DeviceUpsert {
	u.SetExcluded(device.FieldKeepAlive)
	return u
}

// SetEndpoint sets the "endpoint" field.
func (u *DeviceUpsert) SetEndpoint(v types.Inet) *DeviceUpsert {
	u.Set(device.FieldEndpoint, v)
	return u
}

// UpdateEndpoint sets the "endpoint" field to the value that was provided on create.
func (u *DeviceUpsert) UpdateEndpoint() *DeviceUpsert {
	u.SetExcluded(device.FieldEndpoint)
	return u
}

// SetAllowedIps sets the "allowed_ips" field.
func (u *DeviceUpsert) SetAllowedIps(v types.CidrSlice) *DeviceUpsert {
	u.Set(device.FieldAllowedIps, v)
	return u
}

// UpdateAllowedIps sets the "allowed_ips" field to the value that was provided on create.
func (u *DeviceUpsert) UpdateAllowedIps() *DeviceUpsert {
	u.SetExcluded(device.FieldAllowedIps)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Device.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(device.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DeviceUpsertOne) UpdateNewValues() *DeviceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(device.FieldID)
		}
		if _, exists := u.create.mutation.PublicKey(); exists {
			s.SetIgnore(device.FieldPublicKey)
		}
		if _, exists := u.create.mutation.PresharedKey(); exists {
			s.SetIgnore(device.FieldPresharedKey)
		}
		if _, exists := u.create.mutation.UserID(); exists {
			s.SetIgnore(device.FieldUserID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Device.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DeviceUpsertOne) Ignore() *DeviceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeviceUpsertOne) DoNothing() *DeviceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeviceCreate.OnConflict
// documentation for more info.
func (u *DeviceUpsertOne) Update(set func(*DeviceUpsert)) *DeviceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeviceUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *DeviceUpsertOne) SetName(v string) *DeviceUpsertOne {
	return u.Update(func(s *DeviceUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DeviceUpsertOne) UpdateName() *DeviceUpsertOne {
	return u.Update(func(s *DeviceUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *DeviceUpsertOne) SetDescription(v string) *DeviceUpsertOne {
	return u.Update(func(s *DeviceUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *DeviceUpsertOne) UpdateDescription() *DeviceUpsertOne {
	return u.Update(func(s *DeviceUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *DeviceUpsertOne) ClearDescription() *DeviceUpsertOne {
	return u.Update(func(s *DeviceUpsert) {
		s.ClearDescription()
	})
}

// SetType sets the "type" field.
func (u *DeviceUpsertOne) SetType(v string) *DeviceUpsertOne {
	return u.Update(func(s *DeviceUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *DeviceUpsertOne) UpdateType() *DeviceUpsertOne {
	return u.Update(func(s *DeviceUpsert) {
		s.UpdateType()
	})
}

// SetDNS sets the "dns" field.
func (u *DeviceUpsertOne) SetDNS(v types.InetSlice) *DeviceUpsertOne {
	return u.Update(func(s *DeviceUpsert) {
		s.SetDNS(v)
	})
}

// UpdateDNS sets the "dns" field to the value that was provided on create.
func (u *DeviceUpsertOne) UpdateDNS() *DeviceUpsertOne {
	return u.Update(func(s *DeviceUpsert) {
		s.UpdateDNS()
	})
}

// SetKeepAlive sets the "keep_alive" field.
func (u *DeviceUpsertOne) SetKeepAlive(v bool) *DeviceUpsertOne {
	return u.Update(func(s *DeviceUpsert) {
		s.SetKeepAlive(v)
	})
}

// UpdateKeepAlive sets the "keep_alive" field to the value that was provided on create.
func (u *DeviceUpsertOne) UpdateKeepAlive() *DeviceUpsertOne {
	return u.Update(func(s *DeviceUpsert) {
		s.UpdateKeepAlive()
	})
}

// SetEndpoint sets the "endpoint" field.
func (u *DeviceUpsertOne) SetEndpoint(v types.Inet) *DeviceUpsertOne {
	return u.Update(func(s *DeviceUpsert) {
		s.SetEndpoint(v)
	})
}

// UpdateEndpoint sets the "endpoint" field to the value that was provided on create.
func (u *DeviceUpsertOne) UpdateEndpoint() *DeviceUpsertOne {
	return u.Update(func(s *DeviceUpsert) {
		s.UpdateEndpoint()
	})
}

// SetAllowedIps sets the "allowed_ips" field.
func (u *DeviceUpsertOne) SetAllowedIps(v types.CidrSlice) *DeviceUpsertOne {
	return u.Update(func(s *DeviceUpsert) {
		s.SetAllowedIps(v)
	})
}

// UpdateAllowedIps sets the "allowed_ips" field to the value that was provided on create.
func (u *DeviceUpsertOne) UpdateAllowedIps() *DeviceUpsertOne {
	return u.Update(func(s *DeviceUpsert) {
		s.UpdateAllowedIps()
	})
}

// Exec executes the query.
func (u *DeviceUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DeviceCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeviceUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DeviceUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: DeviceUpsertOne.ID is not supported by MySQL driver. Use DeviceUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DeviceUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DeviceCreateBulk is the builder for creating many Device entities in bulk.
type DeviceCreateBulk struct {
	config
	err      error
	builders []*DeviceCreate
	conflict []sql.ConflictOption
}

// Save creates the Device entities in the database.
func (dcb *DeviceCreateBulk) Save(ctx context.Context) ([]*Device, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Device, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeviceMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DeviceCreateBulk) SaveX(ctx context.Context) []*Device {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DeviceCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DeviceCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Device.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeviceUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (dcb *DeviceCreateBulk) OnConflict(opts ...sql.ConflictOption) *DeviceUpsertBulk {
	dcb.conflict = opts
	return &DeviceUpsertBulk{
		create: dcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Device.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dcb *DeviceCreateBulk) OnConflictColumns(columns ...string) *DeviceUpsertBulk {
	dcb.conflict = append(dcb.conflict, sql.ConflictColumns(columns...))
	return &DeviceUpsertBulk{
		create: dcb,
	}
}

// DeviceUpsertBulk is the builder for "upsert"-ing
// a bulk of Device nodes.
type DeviceUpsertBulk struct {
	create *DeviceCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Device.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(device.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DeviceUpsertBulk) UpdateNewValues() *DeviceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(device.FieldID)
			}
			if _, exists := b.mutation.PublicKey(); exists {
				s.SetIgnore(device.FieldPublicKey)
			}
			if _, exists := b.mutation.PresharedKey(); exists {
				s.SetIgnore(device.FieldPresharedKey)
			}
			if _, exists := b.mutation.UserID(); exists {
				s.SetIgnore(device.FieldUserID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Device.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DeviceUpsertBulk) Ignore() *DeviceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeviceUpsertBulk) DoNothing() *DeviceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeviceCreateBulk.OnConflict
// documentation for more info.
func (u *DeviceUpsertBulk) Update(set func(*DeviceUpsert)) *DeviceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeviceUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *DeviceUpsertBulk) SetName(v string) *DeviceUpsertBulk {
	return u.Update(func(s *DeviceUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DeviceUpsertBulk) UpdateName() *DeviceUpsertBulk {
	return u.Update(func(s *DeviceUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *DeviceUpsertBulk) SetDescription(v string) *DeviceUpsertBulk {
	return u.Update(func(s *DeviceUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *DeviceUpsertBulk) UpdateDescription() *DeviceUpsertBulk {
	return u.Update(func(s *DeviceUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *DeviceUpsertBulk) ClearDescription() *DeviceUpsertBulk {
	return u.Update(func(s *DeviceUpsert) {
		s.ClearDescription()
	})
}

// SetType sets the "type" field.
func (u *DeviceUpsertBulk) SetType(v string) *DeviceUpsertBulk {
	return u.Update(func(s *DeviceUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *DeviceUpsertBulk) UpdateType() *DeviceUpsertBulk {
	return u.Update(func(s *DeviceUpsert) {
		s.UpdateType()
	})
}

// SetDNS sets the "dns" field.
func (u *DeviceUpsertBulk) SetDNS(v types.InetSlice) *DeviceUpsertBulk {
	return u.Update(func(s *DeviceUpsert) {
		s.SetDNS(v)
	})
}

// UpdateDNS sets the "dns" field to the value that was provided on create.
func (u *DeviceUpsertBulk) UpdateDNS() *DeviceUpsertBulk {
	return u.Update(func(s *DeviceUpsert) {
		s.UpdateDNS()
	})
}

// SetKeepAlive sets the "keep_alive" field.
func (u *DeviceUpsertBulk) SetKeepAlive(v bool) *DeviceUpsertBulk {
	return u.Update(func(s *DeviceUpsert) {
		s.SetKeepAlive(v)
	})
}

// UpdateKeepAlive sets the "keep_alive" field to the value that was provided on create.
func (u *DeviceUpsertBulk) UpdateKeepAlive() *DeviceUpsertBulk {
	return u.Update(func(s *DeviceUpsert) {
		s.UpdateKeepAlive()
	})
}

// SetEndpoint sets the "endpoint" field.
func (u *DeviceUpsertBulk) SetEndpoint(v types.Inet) *DeviceUpsertBulk {
	return u.Update(func(s *DeviceUpsert) {
		s.SetEndpoint(v)
	})
}

// UpdateEndpoint sets the "endpoint" field to the value that was provided on create.
func (u *DeviceUpsertBulk) UpdateEndpoint() *DeviceUpsertBulk {
	return u.Update(func(s *DeviceUpsert) {
		s.UpdateEndpoint()
	})
}

// SetAllowedIps sets the "allowed_ips" field.
func (u *DeviceUpsertBulk) SetAllowedIps(v types.CidrSlice) *DeviceUpsertBulk {
	return u.Update(func(s *DeviceUpsert) {
		s.SetAllowedIps(v)
	})
}

// UpdateAllowedIps sets the "allowed_ips" field to the value that was provided on create.
func (u *DeviceUpsertBulk) UpdateAllowedIps() *DeviceUpsertBulk {
	return u.Update(func(s *DeviceUpsert) {
		s.UpdateAllowedIps()
	})
}

// Exec executes the query.
func (u *DeviceUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DeviceCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DeviceCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeviceUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
