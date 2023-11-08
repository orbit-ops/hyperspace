// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/tiagoposse/connect/ent/device"
	"github.com/tiagoposse/connect/ent/user"
	"github.com/tiagoposse/connect/internal/types"
)

// Device is the model entity for the Device schema.
type Device struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// DNS holds the value of the "dns" field.
	DNS types.InetSlice `json:"dns,omitempty"`
	// PublicKey holds the value of the "public_key" field.
	PublicKey string `json:"public_key,omitempty"`
	// PresharedKey holds the value of the "preshared_key" field.
	PresharedKey string `json:"preshared_key,omitempty"`
	// KeepAlive holds the value of the "keep_alive" field.
	KeepAlive bool `json:"keep_alive,omitempty"`
	// Endpoint holds the value of the "endpoint" field.
	Endpoint types.Inet `json:"endpoint,omitempty"`
	// AllowedIps holds the value of the "allowed_ips" field.
	AllowedIps types.CidrSlice `json:"allowed_ips,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeviceQuery when eager-loading is set.
	Edges        DeviceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DeviceEdges holds the relations/edges for other nodes in the graph.
type DeviceEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeviceEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Device) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case device.FieldKeepAlive:
			values[i] = new(sql.NullBool)
		case device.FieldName, device.FieldDescription, device.FieldType, device.FieldPublicKey, device.FieldPresharedKey, device.FieldUserID:
			values[i] = new(sql.NullString)
		case device.FieldAllowedIps:
			values[i] = new(types.CidrSlice)
		case device.FieldEndpoint:
			values[i] = new(types.Inet)
		case device.FieldDNS:
			values[i] = new(types.InetSlice)
		case device.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Device fields.
func (d *Device) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case device.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				d.ID = *value
			}
		case device.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		case device.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				d.Description = value.String
			}
		case device.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				d.Type = value.String
			}
		case device.FieldDNS:
			if value, ok := values[i].(*types.InetSlice); !ok {
				return fmt.Errorf("unexpected type %T for field dns", values[i])
			} else if value != nil {
				d.DNS = *value
			}
		case device.FieldPublicKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field public_key", values[i])
			} else if value.Valid {
				d.PublicKey = value.String
			}
		case device.FieldPresharedKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field preshared_key", values[i])
			} else if value.Valid {
				d.PresharedKey = value.String
			}
		case device.FieldKeepAlive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field keep_alive", values[i])
			} else if value.Valid {
				d.KeepAlive = value.Bool
			}
		case device.FieldEndpoint:
			if value, ok := values[i].(*types.Inet); !ok {
				return fmt.Errorf("unexpected type %T for field endpoint", values[i])
			} else if value != nil {
				d.Endpoint = *value
			}
		case device.FieldAllowedIps:
			if value, ok := values[i].(*types.CidrSlice); !ok {
				return fmt.Errorf("unexpected type %T for field allowed_ips", values[i])
			} else if value != nil {
				d.AllowedIps = *value
			}
		case device.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				d.UserID = value.String
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Device.
// This includes values selected through modifiers, order, etc.
func (d *Device) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Device entity.
func (d *Device) QueryUser() *UserQuery {
	return NewDeviceClient(d.config).QueryUser(d)
}

// Update returns a builder for updating this Device.
// Note that you need to call Device.Unwrap() before calling this method if this Device
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Device) Update() *DeviceUpdateOne {
	return NewDeviceClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Device entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Device) Unwrap() *Device {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Device is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Device) String() string {
	var builder strings.Builder
	builder.WriteString("Device(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("name=")
	builder.WriteString(d.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(d.Description)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(d.Type)
	builder.WriteString(", ")
	builder.WriteString("dns=")
	builder.WriteString(fmt.Sprintf("%v", d.DNS))
	builder.WriteString(", ")
	builder.WriteString("public_key=")
	builder.WriteString(d.PublicKey)
	builder.WriteString(", ")
	builder.WriteString("preshared_key=")
	builder.WriteString(d.PresharedKey)
	builder.WriteString(", ")
	builder.WriteString("keep_alive=")
	builder.WriteString(fmt.Sprintf("%v", d.KeepAlive))
	builder.WriteString(", ")
	builder.WriteString("endpoint=")
	builder.WriteString(fmt.Sprintf("%v", d.Endpoint))
	builder.WriteString(", ")
	builder.WriteString("allowed_ips=")
	builder.WriteString(fmt.Sprintf("%v", d.AllowedIps))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(d.UserID)
	builder.WriteByte(')')
	return builder.String()
}

// Devices is a parsable slice of Device.
type Devices []*Device
