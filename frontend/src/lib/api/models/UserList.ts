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
 * @interface UserList
 */
export interface UserList {
    /**
     * 
     * @type {string}
     * @memberof UserList
     */
    id: string;
    /**
     * 
     * @type {string}
     * @memberof UserList
     */
    email: string;
    /**
     * 
     * @type {string}
     * @memberof UserList
     */
    firstname: string;
    /**
     * 
     * @type {string}
     * @memberof UserList
     */
    lastname: string;
    /**
     * 
     * @type {string}
     * @memberof UserList
     */
    provider: string;
    /**
     * 
     * @type {string}
     * @memberof UserList
     */
    photoUrl?: string;
    /**
     * 
     * @type {boolean}
     * @memberof UserList
     */
    disabled: boolean;
    /**
     * 
     * @type {string}
     * @memberof UserList
     */
    disabledReason?: string;
}

/**
 * Check if a given object implements the UserList interface.
 */
export function instanceOfUserList(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "id" in value;
    isInstance = isInstance && "email" in value;
    isInstance = isInstance && "firstname" in value;
    isInstance = isInstance && "lastname" in value;
    isInstance = isInstance && "provider" in value;
    isInstance = isInstance && "disabled" in value;

    return isInstance;
}

export function UserListFromJSON(json: any): UserList {
    return UserListFromJSONTyped(json, false);
}

export function UserListFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserList {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'email': json['email'],
        'firstname': json['firstname'],
        'lastname': json['lastname'],
        'provider': json['provider'],
        'photoUrl': !exists(json, 'photo_url') ? undefined : json['photo_url'],
        'disabled': json['disabled'],
        'disabledReason': !exists(json, 'disabled_reason') ? undefined : json['disabled_reason'],
    };
}

export function UserListToJSON(value?: UserList | null): any {
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
        'photo_url': value.photoUrl,
        'disabled': value.disabled,
        'disabled_reason': value.disabledReason,
    };
}
