import {MethodOptions} from "../decorators";
import Koa from "koa";


export const execApi = async (ctx: Koa.ParameterizedContext, apiFn: () => Promise<any>, methodOpts: MethodOptions) => {
    try {
        ctx.body = await apiFn();
        ctx.status = methodOpts.successStatusCode || 200;
        Object.entries(methodOpts.successHeaders || {}).forEach(([key, value]) => {
            ctx.res.setHeader(key, value);
        });
    }
    catch (
        err
        ) {
        ctx.status = 400;
        const errorCode = ((err as any).code || "") as string;
        for (const [key, value] of Object.entries(methodOpts.errorStatusCodes || {})) {
            if(typeof value === "string" && value === errorCode) {
                ctx.status = parseInt(key);
                break;
            }
            else if(Array.isArray(value) && value.includes(errorCode)) {
                ctx.status = parseInt(key);
                break;
            }
        }

        Object.entries(methodOpts.errorHeaders || {}).forEach(([key, value]) => {
            ctx.res.setHeader(key, value);
        });
        throw err;
    }
}
