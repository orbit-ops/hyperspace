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
 * @interface DeviceList
 */
export interface DeviceList {
    /**
     * 
     * @type {number}
     * @memberof DeviceList
     */
    id: number;
    /**
     * 
     * @type {string}
     * @memberof DeviceList
     */
    name: string;
    /**
     * 
     * @type {string}
     * @memberof DeviceList
     */
    description?: string;
    /**
     * 
     * @type {string}
     * @memberof DeviceList
     */
    type: string;
    /**
     * 
     * @type {string}
     * @memberof DeviceList
     */
    publicKey: string;
    /**
     * 
     * @type {string}
     * @memberof DeviceList
     */
    endpoint: string;
    /**
     * 
     * @type {string}
     * @memberof DeviceList
     */
    allowedIps: string;
}

/**
 * Check if a given object implements the DeviceList interface.
 */
export function instanceOfDeviceList(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "id" in value;
    isInstance = isInstance && "name" in value;
    isInstance = isInstance && "type" in value;
    isInstance = isInstance && "publicKey" in value;
    isInstance = isInstance && "endpoint" in value;
    isInstance = isInstance && "allowedIps" in value;

    return isInstance;
}

export function DeviceListFromJSON(json: any): DeviceList {
    return DeviceListFromJSONTyped(json, false);
}

export function DeviceListFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeviceList {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'name': json['name'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'type': json['type'],
        'publicKey': json['public_key'],
        'endpoint': json['endpoint'],
        'allowedIps': json['allowed_ips'],
    };
}

export function DeviceListToJSON(value?: DeviceList | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'name': value.name,
        'description': value.description,
        'type': value.type,
        'public_key': value.publicKey,
        'endpoint': value.endpoint,
        'allowed_ips': value.allowedIps,
    };
}

