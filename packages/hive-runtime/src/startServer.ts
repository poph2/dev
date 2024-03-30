import Koa from 'koa';
import {HiveConfig} from "./HiveConfig";
import { register } from 'ts-node';
import path from 'node:path';
import Router from "@koa/router";
import koaBody from "koa-body";
import cors from "@koa/cors";



export const startServer = async (routers: Router[], middlewares: Koa.Middleware[] = []) => {


    register({
        project: path.join(process.cwd(), "tsconfig.json"),
    })

    const hiveConfig = require(path.join(process.cwd(), 'hive.config.ts')) as HiveConfig;

    console.log(hiveConfig, "loaded at runtime")

    const app = new Koa();

    app
        .use(cors())
        .use(koaBody())


    middlewares.forEach(middleware => {
        app.use(middleware);
    })

    routers.forEach(router => {
        app.use(router.routes());
    })

    app.listen(hiveConfig.port);

    console.log(`Server is listening on http://localhost:${hiveConfig.port}`);
}