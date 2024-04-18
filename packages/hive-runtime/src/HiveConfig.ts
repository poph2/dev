import Router from "@koa/router";
import Koa from "koa";

export type BaseHiveConfig = {
  port: number;

  typeorm: {
    entities: string[];
  };
};

export type HiveServerConfig = BaseHiveConfig & {
  routers: Router[];
  middlewares: Koa.Middleware[];
};

export type HiveConfig = BaseHiveConfig;
