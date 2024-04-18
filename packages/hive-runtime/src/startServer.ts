import path from "node:path";
import cors from "@koa/cors";
import Router from "@koa/router";
import { HiveServerConfig } from "hive-config";
import Koa from "koa";
import koaBody from "koa-body";
import { register } from "ts-node";

export const startServer = async (
  opts: HiveServerConfig<Router, Koa.Middleware>,
) => {
  register({
    project: path.join(process.cwd(), "tsconfig.json"),
  });

  // const hiveConfig = require(path.join(process.cwd(), 'hive.config.ts')) as HiveConfig;

  // console.log(hiveConfig, "loaded at runtime")

  const app = new Koa();

  app.use(cors()).use(koaBody());

  opts.middlewares.forEach((middleware) => {
    app.use(middleware);
  });

  opts.routers.forEach((router) => {
    app.use(router.routes());
  });

  app.listen(opts.port);

  console.log(`Server is listening on http://localhost:${opts.port}`);
};
