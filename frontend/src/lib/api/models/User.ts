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
import type { ApiKey } from './ApiKey';
import {
    ApiKeyFromJSON,
    ApiKeyFromJSONTyped,
    ApiKeyToJSON,
} from './ApiKey';
import type { Audit } from './Audit';
import {
    AuditFromJSON,
    AuditFromJSONTyped,
    AuditToJSON,
} from './Audit';
import type { Device } from './Device';
import {
    DeviceFromJSON,
    DeviceFromJSONTyped,
    DeviceToJSON,
} from './Device';
import type { Group } from './Group';
import {
    GroupFromJSON,
    GroupFromJSONTyped,
    GroupToJSON,
} from './Group';

/**
 * 
 * @export
 * @interface User
 */
export interface User {
    /**
     * 
     * @type {string}
     * @memberof User
     */
    id: string;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    email: string;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    firstname: string;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    lastname: string;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    provider: string;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    password?: string;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    salt?: string;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    photoUrl?: string;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    disabled: boolean;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    disabledReason?: string;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    groupId: string;
    /**
     * 
     * @type {Array<Device>}
     * @memberof User
     */
    devices?: Array<Device>;
    /**
     * 
     * @type {Array<ApiKey>}
     * @memberof User
     */
    keys?: Array<ApiKey>;
    /**
     * 
     * @type {Array<Audit>}
     * @memberof User
     */
    audit?: Array<Audit>;
    /**
     * 
     * @type {Group}
     * @memberof User
     */
    group: Group;
}

/**
 * Check if a given object implements the User interface.
 */
export function instanceOfUser(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "id" in value;
    isInstance = isInstance && "email" in value;
    isInstance = isInstance && "firstname" in value;
    isInstance = isInstance && "lastname" in value;
    isInstance = isInstance && "provider" in value;
    isInstance = isInstance && "disabled" in value;
    isInstance = isInstance && "groupId" in value;
    isInstance = isInstance && "group" in value;

    return isInstance;
}

export function UserFromJSON(json: any): User {
    return UserFromJSONTyped(json, false);
}

export function UserFromJSONTyped(json: any, ignoreDiscriminator: boolean): User {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'email': json['email'],
        'firstname': json['firstname'],
        'lastname': json['lastname'],
        'provider': json['provider'],
        'password': !exists(json, 'password') ? undefined : json['password'],
        'salt': !exists(json, 'salt') ? undefined : json['salt'],
        'photoUrl': !exists(json, 'photo_url') ? undefined : json['photo_url'],
        'disabled': json['disabled'],
        'disabledReason': !exists(json, 'disabled_reason') ? undefined : json['disabled_reason'],
        'groupId': json['group_id'],
        'devices': !exists(json, 'devices') ? undefined : ((json['devices'] as Array<any>).map(DeviceFromJSON)),
        'keys': !exists(json, 'keys') ? undefined : ((json['keys'] as Array<any>).map(ApiKeyFromJSON)),
        'audit': !exists(json, 'audit') ? undefined : ((json['audit'] as Array<any>).map(AuditFromJSON)),
        'group': GroupFromJSON(json['group']),
    };
}

export function UserToJSON(value?: User | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'email': value.email,
        'firstname': value.firstname,
        'lastname': value.lastname,
        'provider': value.provider,
        'password': value.password,
        'salt': value.salt,
        'photo_url': value.photoUrl,
        'disabled': value.disabled,
        'disabled_reason': value.disabledReason,
        'group_id': value.groupId,
        'devices': value.devices === undefined ? undefined : ((value.devices as Array<any>).map(DeviceToJSON)),
        'keys': value.keys === undefined ? undefined : ((value.keys as Array<any>).map(ApiKeyToJSON)),
        'audit': value.audit === undefined ? undefined : ((value.audit as Array<any>).map(AuditToJSON)),
        'group': GroupToJSON(value.group),
    };
}

