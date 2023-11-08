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
 * @interface UserpassLoginRequest
 */
export interface UserpassLoginRequest {
    /**
     * 
     * @type {string}
     * @memberof UserpassLoginRequest
     */
    username: string;
    /**
     * 
     * @type {string}
     * @memberof UserpassLoginRequest
     */
    password: string;
}

/**
 * Check if a given object implements the UserpassLoginRequest interface.
 */
export function instanceOfUserpassLoginRequest(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "username" in value;
    isInstance = isInstance && "password" in value;

    return isInstance;
}

export function UserpassLoginRequestFromJSON(json: any): UserpassLoginRequest {
    return UserpassLoginRequestFromJSONTyped(json, false);
}

export function UserpassLoginRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserpassLoginRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'username': json['username'],
        'password': json['password'],
    };
}

export function UserpassLoginRequestToJSON(value?: UserpassLoginRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'username': value.username,
        'password': value.password,
    };
}

