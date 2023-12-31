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
/**
 * 
 * @export
 * @interface Status200Response
 */
export interface Status200Response {
    /**
     * 
     * @type {string}
     * @memberof Status200Response
     */
    id: string;
    /**
     * 
     * @type {string}
     * @memberof Status200Response
     */
    photoUrl: string;
    /**
     * 
     * @type {string}
     * @memberof Status200Response
     */
    provider: string;
    /**
     * 
     * @type {string}
     * @memberof Status200Response
     */
    email: string;
    /**
     * 
     * @type {string}
     * @memberof Status200Response
     */
    lastname: string;
    /**
     * 
     * @type {string}
     * @memberof Status200Response
     */
    firstname: string;
    /**
     * 
     * @type {string}
     * @memberof Status200Response
     */
    group: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof Status200Response
     */
    scopes: Array<Status200ResponseScopesEnum>;
}


/**
 * @export
 */
export const Status200ResponseScopesEnum = {
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
export type Status200ResponseScopesEnum = typeof Status200ResponseScopesEnum[keyof typeof Status200ResponseScopesEnum];


/**
 * Check if a given object implements the Status200Response interface.
 */
export function instanceOfStatus200Response(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "id" in value;
    isInstance = isInstance && "photoUrl" in value;
    isInstance = isInstance && "provider" in value;
    isInstance = isInstance && "email" in value;
    isInstance = isInstance && "lastname" in value;
    isInstance = isInstance && "firstname" in value;
    isInstance = isInstance && "group" in value;
    isInstance = isInstance && "scopes" in value;

    return isInstance;
}

export function Status200ResponseFromJSON(json: any): Status200Response {
    return Status200ResponseFromJSONTyped(json, false);
}

export function Status200ResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): Status200Response {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'photoUrl': json['photo_url'],
        'provider': json['provider'],
        'email': json['email'],
        'lastname': json['lastname'],
        'firstname': json['firstname'],
        'group': json['group'],
        'scopes': json['scopes'],
    };
}

export function Status200ResponseToJSON(value?: Status200Response | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'photo_url': value.photoUrl,
        'provider': value.provider,
        'email': value.email,
        'lastname': value.lastname,
        'firstname': value.firstname,
        'group': value.group,
        'scopes': value.scopes,
    };
}

