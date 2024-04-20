import cors from "@koa/cors";
import Router from "@koa/router";
import Koa from "koa";
import koaBody from "koa-body";

export type HiveRuntimeConfig = {
  routers: Router[];
  middlewares: Koa.Middleware[];
  port: number;
};

export const startServer = async (config: HiveRuntimeConfig) => {
  const app = new Koa();

  app.use(cors()).use(koaBody());

  app.use(cors()).use(koaBody());

  config.middlewares.forEach((middleware) => {
    app.use(middleware);
  });

  config.routers.forEach((router) => {
    app.use(router.routes());
  });

  app.listen(config.port);

  console.log(`Server is listening on http://localhost:${config.port}`);
};
