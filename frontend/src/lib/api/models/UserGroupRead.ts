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

import { exists, mapValues } from '../runtime';
import type { CreateGroupRequestRulesInner } from './CreateGroupRequestRulesInner';
import {
    CreateGroupRequestRulesInnerFromJSON,
    CreateGroupRequestRulesInnerFromJSONTyped,
    CreateGroupRequestRulesInnerToJSON,
} from './CreateGroupRequestRulesInner';

/**
 * 
 * @export
 * @interface UserGroupRead
 */
export interface UserGroupRead {
    /**
     * 
     * @type {string}
     * @memberof UserGroupRead
     */
    id: string;
    /**
     * 
     * @type {string}
     * @memberof UserGroupRead
     */
    name: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof UserGroupRead
     */
    scopes: Array<UserGroupReadScopesEnum>;
    /**
     * 
     * @type {string}
     * @memberof UserGroupRead
     */
    cidr: string;
    /**
     * 
     * @type {Array<CreateGroupRequestRulesInner>}
     * @memberof UserGroupRead
     */
    rules: Array<CreateGroupRequestRulesInner>;
}


/**
 * @export
 */
export const UserGroupReadScopesEnum = {
    AdminUsersWrite: 'admin.users.write',
    AdminUsersReadonly: 'admin.users.readonly',
    AdminGroupsWrite: 'admin.groups.write',
    AdminGroupsReadonly: 'admin.groups.readonly',
    AdminDevicesWrite: 'admin.devices.write',
    AdminDevicesReadonly: 'admin.devices.readonly',
    AdminSettingsWrite: 'admin.settings.write',
    AdminSettingsReadonly: 'admin.settings.readonly',
    Admin: 'admin.*',
    UserDevicesWrite: 'user.devices.write',
    UserApikeyWrite: 'user.apikey.write',
    UserDevicesReadonly: 'user.devices.readonly',
    User: 'user.*'
} as const;
export type UserGroupReadScopesEnum = typeof UserGroupReadScopesEnum[keyof typeof UserGroupReadScopesEnum];


/**
 * Check if a given object implements the UserGroupRead interface.
 */
export function instanceOfUserGroupRead(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "id" in value;
    isInstance = isInstance && "name" in value;
    isInstance = isInstance && "scopes" in value;
    isInstance = isInstance && "cidr" in value;
    isInstance = isInstance && "rules" in value;

    return isInstance;
}

export function UserGroupReadFromJSON(json: any): UserGroupRead {
    return UserGroupReadFromJSONTyped(json, false);
}

export function UserGroupReadFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserGroupRead {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'name': json['name'],
        'scopes': json['scopes'],
        'cidr': json['cidr'],
        'rules': ((json['rules'] as Array<any>).map(CreateGroupRequestRulesInnerFromJSON)),
    };
}

export function UserGroupReadToJSON(value?: UserGroupRead | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'name': value.name,
        'scopes': value.scopes,
        'cidr': value.cidr,
        'rules': ((value.rules as Array<any>).map(CreateGroupRequestRulesInnerToJSON)),
    };
}

