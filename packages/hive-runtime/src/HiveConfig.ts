import Koa from 'koa';

export type HiveConfig = {
    port: number;
    middlewares: Koa.Middleware[];
}