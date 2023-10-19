// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CreateApiKey implements createApiKey operation.
	//
	// Creates a new ApiKey and persists it to storage.
	//
	// POST /api-keys
	CreateApiKey(ctx context.Context, req *CreateApiKeyReq) (CreateApiKeyRes, error)
	// CreateDevice implements createDevice operation.
	//
	// Creates a new Device and persists it to storage.
	//
	// POST /devices
	CreateDevice(ctx context.Context, req *CreateDeviceReq) (CreateDeviceRes, error)
	// CreateGroup implements createGroup operation.
	//
	// Creates a new Group and persists it to storage.
	//
	// POST /groups
	CreateGroup(ctx context.Context, req *CreateGroupReq) (CreateGroupRes, error)
	// CreateUser implements createUser operation.
	//
	// Creates a new User and persists it to storage.
	//
	// POST /users
	CreateUser(ctx context.Context, req *CreateUserReq) (CreateUserRes, error)
	// DeleteApiKey implements deleteApiKey operation.
	//
	// Deletes the ApiKey with the requested ID.
	//
	// DELETE /api-keys/{id}
	DeleteApiKey(ctx context.Context, params DeleteApiKeyParams) (DeleteApiKeyRes, error)
	// DeleteDevice implements deleteDevice operation.
	//
	// Deletes the Device with the requested ID.
	//
	// DELETE /devices/{id}
	DeleteDevice(ctx context.Context, params DeleteDeviceParams) (DeleteDeviceRes, error)
	// DeleteGroup implements deleteGroup operation.
	//
	// Deletes the Group with the requested ID.
	//
	// DELETE /groups/{id}
	DeleteGroup(ctx context.Context, params DeleteGroupParams) (DeleteGroupRes, error)
	// DeleteUser implements deleteUser operation.
	//
	// Deletes the User with the requested ID.
	//
	// DELETE /users/{id}
	DeleteUser(ctx context.Context, params DeleteUserParams) (DeleteUserRes, error)
	// GoogleAuthCallback implements googleAuthCallback operation.
	//
	// POST /auth/google/callback
	GoogleAuthCallback(ctx context.Context, req OptGoogleAuthCallbackReq) (GoogleAuthCallbackRes, error)
	// GoogleAuthStart implements googleAuthStart operation.
	//
	// GET /auth/google/start
	GoogleAuthStart(ctx context.Context, params GoogleAuthStartParams) (GoogleAuthStartRes, error)
	// GoogleAuthSync implements googleAuthSync operation.
	//
	// Synchronize users for the google provider.
	//
	// GET /auth/google/sync
	GoogleAuthSync(ctx context.Context) (GoogleAuthSyncRes, error)
	// ListApiKey implements listApiKey operation.
	//
	// List ApiKeys.
	//
	// GET /api-keys
	ListApiKey(ctx context.Context, params ListApiKeyParams) (ListApiKeyRes, error)
	// ListDevice implements listDevice operation.
	//
	// List Devices.
	//
	// GET /devices
	ListDevice(ctx context.Context, params ListDeviceParams) (ListDeviceRes, error)
	// ListGroup implements listGroup operation.
	//
	// List Groups.
	//
	// GET /groups
	ListGroup(ctx context.Context, params ListGroupParams) (ListGroupRes, error)
	// ListGroupUsers implements listGroupUsers operation.
	//
	// List attached Users.
	//
	// GET /groups/{id}/users
	ListGroupUsers(ctx context.Context, params ListGroupUsersParams) (ListGroupUsersRes, error)
	// ListUser implements listUser operation.
	//
	// List Users.
	//
	// GET /users
	ListUser(ctx context.Context, params ListUserParams) (ListUserRes, error)
	// ListUserDevices implements listUserDevices operation.
	//
	// List attached Devices.
	//
	// GET /users/{id}/devices
	ListUserDevices(ctx context.Context, params ListUserDevicesParams) (ListUserDevicesRes, error)
	// ListUserKeys implements listUserKeys operation.
	//
	// List attached Keys.
	//
	// GET /users/{id}/keys
	ListUserKeys(ctx context.Context, params ListUserKeysParams) (ListUserKeysRes, error)
	// ReadApiKey implements readApiKey operation.
	//
	// Finds the ApiKey with the requested ID and returns it.
	//
	// GET /api-keys/{id}
	ReadApiKey(ctx context.Context, params ReadApiKeyParams) (ReadApiKeyRes, error)
	// ReadApiKeyUser implements readApiKeyUser operation.
	//
	// Find the attached User of the ApiKey with the given ID.
	//
	// GET /api-keys/{id}/user
	ReadApiKeyUser(ctx context.Context, params ReadApiKeyUserParams) (ReadApiKeyUserRes, error)
	// ReadDevice implements readDevice operation.
	//
	// Finds the Device with the requested ID and returns it.
	//
	// GET /devices/{id}
	ReadDevice(ctx context.Context, params ReadDeviceParams) (ReadDeviceRes, error)
	// ReadDeviceUser implements readDeviceUser operation.
	//
	// Find the attached User of the Device with the given ID.
	//
	// GET /devices/{id}/user
	ReadDeviceUser(ctx context.Context, params ReadDeviceUserParams) (ReadDeviceUserRes, error)
	// ReadGroup implements readGroup operation.
	//
	// Finds the Group with the requested ID and returns it.
	//
	// GET /groups/{id}
	ReadGroup(ctx context.Context, params ReadGroupParams) (ReadGroupRes, error)
	// ReadUser implements readUser operation.
	//
	// Finds the User with the requested ID and returns it.
	//
	// GET /users/{id}
	ReadUser(ctx context.Context, params ReadUserParams) (ReadUserRes, error)
	// ReadUserGroup implements readUserGroup operation.
	//
	// Find the attached Group of the User with the given ID.
	//
	// GET /users/{id}/group
	ReadUserGroup(ctx context.Context, params ReadUserGroupParams) (ReadUserGroupRes, error)
	// Status implements status operation.
	//
	// Ping the database and report.
	//
	// GET /status
	Status(ctx context.Context) (StatusRes, error)
	// UpdateDevice implements updateDevice operation.
	//
	// Updates a Device and persists changes to storage.
	//
	// PATCH /devices/{id}
	UpdateDevice(ctx context.Context, req *UpdateDeviceReq, params UpdateDeviceParams) (UpdateDeviceRes, error)
	// UpdateUser implements updateUser operation.
	//
	// Updates a User and persists changes to storage.
	//
	// PATCH /users/{id}
	UpdateUser(ctx context.Context, req *UpdateUserReq, params UpdateUserParams) (UpdateUserRes, error)
	// UserpassLogin implements userpassLogin operation.
	//
	// Login with a user and password.
	//
	// POST /auth/userpass/login
	UserpassLogin(ctx context.Context, req OptUserpassLoginReq) (UserpassLoginRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
