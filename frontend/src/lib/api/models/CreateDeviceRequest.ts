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
 * @interface CreateDeviceRequest
 */
export interface CreateDeviceRequest {
    /**
     * 
     * @type {string}
     * @memberof CreateDeviceRequest
     */
    name: string;
    /**
     * 
     * @type {string}
     * @memberof CreateDeviceRequest
     */
    description?: string;
    /**
     * 
     * @type {string}
     * @memberof CreateDeviceRequest
     */
    type: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof CreateDeviceRequest
     */
    dns: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof CreateDeviceRequest
     */
    publicKey: string;
    /**
     * 
     * @type {boolean}
     * @memberof CreateDeviceRequest
     */
    keepAlive: boolean;
    /**
     * 
     * @type {string}
     * @memberof CreateDeviceRequest
     */
    endpoint: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof CreateDeviceRequest
     */
    allowedIps: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof CreateDeviceRequest
     */
    user: string;
}

/**
 * Check if a given object implements the CreateDeviceRequest interface.
 */
export function instanceOfCreateDeviceRequest(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "name" in value;
    isInstance = isInstance && "type" in value;
    isInstance = isInstance && "dns" in value;
    isInstance = isInstance && "publicKey" in value;
    isInstance = isInstance && "keepAlive" in value;
    isInstance = isInstance && "endpoint" in value;
    isInstance = isInstance && "allowedIps" in value;
    isInstance = isInstance && "user" in value;

    return isInstance;
}

export function CreateDeviceRequestFromJSON(json: any): CreateDeviceRequest {
    return CreateDeviceRequestFromJSONTyped(json, false);
}

export function CreateDeviceRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateDeviceRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': json['name'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'type': json['type'],
        'dns': json['dns'],
        'publicKey': json['public_key'],
        'keepAlive': json['keep_alive'],
        'endpoint': json['endpoint'],
        'allowedIps': json['allowed_ips'],
        'user': json['user'],
    };
}

export function CreateDeviceRequestToJSON(value?: CreateDeviceRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
        'description': value.description,
        'type': value.type,
        'dns': value.dns,
        'public_key': value.publicKey,
        'keep_alive': value.keepAlive,
        'endpoint': value.endpoint,
        'allowed_ips': value.allowedIps,
        'user': value.user,
    };
}

