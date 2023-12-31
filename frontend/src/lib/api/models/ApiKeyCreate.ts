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
 * @interface ApiKeyCreate
 */
export interface ApiKeyCreate {
    /**
     * 
     * @type {number}
     * @memberof ApiKeyCreate
     */
    id: number;
    /**
     * 
     * @type {string}
     * @memberof ApiKeyCreate
     */
    name: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof ApiKeyCreate
     */
    scopes: Array<ApiKeyCreateScopesEnum>;
    /**
     * 
     * @type {string}
     * @memberof ApiKeyCreate
     */
    userId: string;
}


/**
 * @export
 */
export const ApiKeyCreateScopesEnum = {
    User: 'user.*',
    Admin: 'admin.*'
} as const;
export type ApiKeyCreateScopesEnum = typeof ApiKeyCreateScopesEnum[keyof typeof ApiKeyCreateScopesEnum];


/**
 * Check if a given object implements the ApiKeyCreate interface.
 */
export function instanceOfApiKeyCreate(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "id" in value;
    isInstance = isInstance && "name" in value;
    isInstance = isInstance && "scopes" in value;
    isInstance = isInstance && "userId" in value;

    return isInstance;
}

export function ApiKeyCreateFromJSON(json: any): ApiKeyCreate {
    return ApiKeyCreateFromJSONTyped(json, false);
}

export function ApiKeyCreateFromJSONTyped(json: any, ignoreDiscriminator: boolean): ApiKeyCreate {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'name': json['name'],
        'scopes': json['scopes'],
        'userId': json['user_id'],
    };
}

export function ApiKeyCreateToJSON(value?: ApiKeyCreate | null): any {
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
        'user_id': value.userId,
    };
}

