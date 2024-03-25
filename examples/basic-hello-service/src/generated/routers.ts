import Router from "@koa/router";
import HomeApi from "../api/home/api";
import {execApi} from "hive-runtime";

const homeApi = new HomeApi();

export const homeApiRouter = new Router()
    .get("/home/api-version", async (ctx) =>
        execApi(ctx, async () => homeApi.getAppVersion(), {}))
    .get("/home/params", async (ctx) =>
        execApi(ctx, async () => homeApi.getParams(), {}));
