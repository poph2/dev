import cors from "@koa/cors";
import Koa from "koa";
import koaBody from "koa-body";
import { HiveServerConfig } from "./HiveConfig";

export const startServer = async (opts: HiveServerConfig) => {
  const app = new Koa();

  app.use(cors()).use(koaBody());

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
