
export type ErrorStatusCodes = { [key: number]: string | string[] };
export type Headers = { [name: string]: number | string | string[] }

export type MethodOptions = {
    successStatusCode?: number;
    successHeaders?: Headers;
    errorStatusCodes?: ErrorStatusCodes;
    errorHeaders?: Headers;
}

export const Get = (path?: string, opts?: MethodOptions): Function => {
    return () => {};
}
export const Post = (path?: string, opts?: MethodOptions): Function => {
    return () => {};
}
export const Put = (path?: string, opts?: MethodOptions): Function => {
    return () => {};
}
export const Patch = (path?: string, opts?: MethodOptions): Function => {
    return () => {};
}
export const Delete = (path?: string, opts?: MethodOptions): Function => {
    return () => {};
}
export const Options = (path?: string, opts?: MethodOptions): Function => {
    return () => {};
}