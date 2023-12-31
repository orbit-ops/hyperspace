/* tslint:disable */
/* eslint-disable */
/**
 * Ent Schema API
 * This is an auto generated API description made out of an Ent schema definition
 *
 * The version of the OpenAPI document: 0.1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  CreateUserRequest,
  ListApiKey400Response,
  UpdateUserRequest,
  UserAuditList,
  UserCreate,
  UserDevicesList,
  UserGroupRead,
  UserKeysList,
  UserList,
  UserRead,
  UserUpdate,
} from '../models/index';
import {
    CreateUserRequestFromJSON,
    CreateUserRequestToJSON,
    ListApiKey400ResponseFromJSON,
    ListApiKey400ResponseToJSON,
    UpdateUserRequestFromJSON,
    UpdateUserRequestToJSON,
    UserAuditListFromJSON,
    UserAuditListToJSON,
    UserCreateFromJSON,
    UserCreateToJSON,
    UserDevicesListFromJSON,
    UserDevicesListToJSON,
    UserGroupReadFromJSON,
    UserGroupReadToJSON,
    UserKeysListFromJSON,
    UserKeysListToJSON,
    UserListFromJSON,
    UserListToJSON,
    UserReadFromJSON,
    UserReadToJSON,
    UserUpdateFromJSON,
    UserUpdateToJSON,
} from '../models/index';

export interface CreateUserOperationRequest {
    createUserRequest: CreateUserRequest;
}

export interface DeleteUserRequest {
    id: string;
}

export interface ListUserRequest {
    page?: number;
    itemsPerPage?: number;
    sort?: string;
    filter?: string;
}

export interface ListUserAuditRequest {
    id: string;
    page?: number;
    itemsPerPage?: number;
    sort?: string;
    filter?: string;
}

export interface ListUserDevicesRequest {
    id: string;
    page?: number;
    itemsPerPage?: number;
    sort?: string;
    filter?: string;
}

export interface ListUserKeysRequest {
    id: string;
    page?: number;
    itemsPerPage?: number;
    sort?: string;
    filter?: string;
}

export interface ReadUserRequest {
    id: string;
}

export interface ReadUserGroupRequest {
    id: string;
}

export interface UpdateUserOperationRequest {
    id: string;
    updateUserRequest: UpdateUserRequest;
}

/**
 * 
 */
export class UserApi extends runtime.BaseAPI {

    /**
     * Creates a new User and persists it to storage.
     * Create a new User
     */
    async createUserRaw(requestParameters: CreateUserOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserCreate>> {
        if (requestParameters.createUserRequest === null || requestParameters.createUserRequest === undefined) {
            throw new runtime.RequiredError('createUserRequest','Required parameter requestParameters.createUserRequest was null or undefined when calling createUser.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["x-api-key"] = this.configuration.apiKey("x-api-key"); // ApiKeyAuth authentication
        }

        const response = await this.request({
            path: `/users`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: CreateUserRequestToJSON(requestParameters.createUserRequest),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UserCreateFromJSON(jsonValue));
    }

    /**
     * Creates a new User and persists it to storage.
     * Create a new User
     */
    async createUser(requestParameters: CreateUserOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserCreate> {
        const response = await this.createUserRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Deletes the User with the requested ID.
     * Deletes a User by ID
     */
    async deleteUserRaw(requestParameters: DeleteUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteUser.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["x-api-key"] = this.configuration.apiKey("x-api-key"); // ApiKeyAuth authentication
        }

        const response = await this.request({
            path: `/users/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Deletes the User with the requested ID.
     * Deletes a User by ID
     */
    async deleteUser(requestParameters: DeleteUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deleteUserRaw(requestParameters, initOverrides);
    }

    /**
     * List Users.
     * List Users
     */
    async listUserRaw(requestParameters: ListUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<UserList>>> {
        const queryParameters: any = {};

        if (requestParameters.sort !== undefined) {
            queryParameters['sort'] = requestParameters.sort;
        }

        if (requestParameters.filter !== undefined) {
            queryParameters['filter'] = requestParameters.filter;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (requestParameters.page !== undefined && requestParameters.page !== null) {
            headerParameters['x-page'] = String(requestParameters.page);
        }

        if (requestParameters.itemsPerPage !== undefined && requestParameters.itemsPerPage !== null) {
            headerParameters['x-items-per-page'] = String(requestParameters.itemsPerPage);
        }

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["x-api-key"] = this.configuration.apiKey("x-api-key"); // ApiKeyAuth authentication
        }

        const response = await this.request({
            path: `/users`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(UserListFromJSON));
    }

    /**
     * List Users.
     * List Users
     */
    async listUser(requestParameters: ListUserRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<UserList>> {
        const response = await this.listUserRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List attached Audits.
     * List attached Audits
     */
    async listUserAuditRaw(requestParameters: ListUserAuditRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<UserAuditList>>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling listUserAudit.');
        }

        const queryParameters: any = {};

        if (requestParameters.sort !== undefined) {
            queryParameters['sort'] = requestParameters.sort;
        }

        if (requestParameters.filter !== undefined) {
            queryParameters['filter'] = requestParameters.filter;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (requestParameters.page !== undefined && requestParameters.page !== null) {
            headerParameters['x-page'] = String(requestParameters.page);
        }

        if (requestParameters.itemsPerPage !== undefined && requestParameters.itemsPerPage !== null) {
            headerParameters['x-items-per-page'] = String(requestParameters.itemsPerPage);
        }

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["x-api-key"] = this.configuration.apiKey("x-api-key"); // ApiKeyAuth authentication
        }

        const response = await this.request({
            path: `/users/{id}/audit`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(UserAuditListFromJSON));
    }

    /**
     * List attached Audits.
     * List attached Audits
     */
    async listUserAudit(requestParameters: ListUserAuditRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<UserAuditList>> {
        const response = await this.listUserAuditRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List attached Devices.
     * List attached Devices
     */
    async listUserDevicesRaw(requestParameters: ListUserDevicesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<UserDevicesList>>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling listUserDevices.');
        }

        const queryParameters: any = {};

        if (requestParameters.sort !== undefined) {
            queryParameters['sort'] = requestParameters.sort;
        }

        if (requestParameters.filter !== undefined) {
            queryParameters['filter'] = requestParameters.filter;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (requestParameters.page !== undefined && requestParameters.page !== null) {
            headerParameters['x-page'] = String(requestParameters.page);
        }

        if (requestParameters.itemsPerPage !== undefined && requestParameters.itemsPerPage !== null) {
            headerParameters['x-items-per-page'] = String(requestParameters.itemsPerPage);
        }

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["x-api-key"] = this.configuration.apiKey("x-api-key"); // ApiKeyAuth authentication
        }

        const response = await this.request({
            path: `/users/{id}/devices`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(UserDevicesListFromJSON));
    }

    /**
     * List attached Devices.
     * List attached Devices
     */
    async listUserDevices(requestParameters: ListUserDevicesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<UserDevicesList>> {
        const response = await this.listUserDevicesRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List attached Keys.
     * List attached Keys
     */
    async listUserKeysRaw(requestParameters: ListUserKeysRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<UserKeysList>>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling listUserKeys.');
        }

        const queryParameters: any = {};

        if (requestParameters.sort !== undefined) {
            queryParameters['sort'] = requestParameters.sort;
        }

        if (requestParameters.filter !== undefined) {
            queryParameters['filter'] = requestParameters.filter;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (requestParameters.page !== undefined && requestParameters.page !== null) {
            headerParameters['x-page'] = String(requestParameters.page);
        }

        if (requestParameters.itemsPerPage !== undefined && requestParameters.itemsPerPage !== null) {
            headerParameters['x-items-per-page'] = String(requestParameters.itemsPerPage);
        }

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["x-api-key"] = this.configuration.apiKey("x-api-key"); // ApiKeyAuth authentication
        }

        const response = await this.request({
            path: `/users/{id}/keys`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(UserKeysListFromJSON));
    }

    /**
     * List attached Keys.
     * List attached Keys
     */
    async listUserKeys(requestParameters: ListUserKeysRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<UserKeysList>> {
        const response = await this.listUserKeysRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Finds the User with the requested ID and returns it.
     * Find a User by ID
     */
    async readUserRaw(requestParameters: ReadUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserRead>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling readUser.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["x-api-key"] = this.configuration.apiKey("x-api-key"); // ApiKeyAuth authentication
        }

        const response = await this.request({
            path: `/users/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UserReadFromJSON(jsonValue));
    }

    /**
     * Finds the User with the requested ID and returns it.
     * Find a User by ID
     */
    async readUser(requestParameters: ReadUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserRead> {
        const response = await this.readUserRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Find the attached Group of the User with the given ID
     * Find the attached Group
     */
    async readUserGroupRaw(requestParameters: ReadUserGroupRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserGroupRead>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling readUserGroup.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["x-api-key"] = this.configuration.apiKey("x-api-key"); // ApiKeyAuth authentication
        }

        const response = await this.request({
            path: `/users/{id}/group`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UserGroupReadFromJSON(jsonValue));
    }

    /**
     * Find the attached Group of the User with the given ID
     * Find the attached Group
     */
    async readUserGroup(requestParameters: ReadUserGroupRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserGroupRead> {
        const response = await this.readUserGroupRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Updates a User and persists changes to storage.
     * Updates a User
     */
    async updateUserRaw(requestParameters: UpdateUserOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserUpdate>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateUser.');
        }

        if (requestParameters.updateUserRequest === null || requestParameters.updateUserRequest === undefined) {
            throw new runtime.RequiredError('updateUserRequest','Required parameter requestParameters.updateUserRequest was null or undefined when calling updateUser.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/users/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: UpdateUserRequestToJSON(requestParameters.updateUserRequest),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UserUpdateFromJSON(jsonValue));
    }

    /**
     * Updates a User and persists changes to storage.
     * Updates a User
     */
    async updateUser(requestParameters: UpdateUserOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserUpdate> {
        const response = await this.updateUserRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
