import Koa from 'koa';
import {HiveConfig} from "./HiveConfig";
import { register } from 'ts-node';
import path from 'node:path';
import Router from "@koa/router";
import koaBody from "koa-body";

export const startServer = async (routers: Router[]) => {


    register({
        project: path.join(process.cwd(), "tsconfig.json"),
    })

    const hiveConfig = require(path.join(process.cwd(), 'hive.config.ts')) as HiveConfig;

    console.log(hiveConfig, "loaded at runtime")

    const app = new Koa().use(koaBody());

    for (const router of routers) {
        app.use(router.routes());
    }

    app.use(async (ctx: Koa.ParameterizedContext) => {
        ctx.body = 'Hello World';
    });

    app.listen(hiveConfig.port);

    console.log(`Server is listening on http://localhost:${hiveConfig.port}`);
}