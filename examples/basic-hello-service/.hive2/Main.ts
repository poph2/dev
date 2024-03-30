import {startServer} from "hive-runtime";
import {RegisterRoutes} from "./routes";
import Router from "@koa/router";
import {koaSwagger} from "koa2-swagger-ui";

const router = new Router();
RegisterRoutes(router);

router.get(
    "/swagger",
    koaSwagger({routePrefix: false, swaggerOptions: {spec: require("./swagger.json")}}),
);

startServer([router]).then(() => {
});
