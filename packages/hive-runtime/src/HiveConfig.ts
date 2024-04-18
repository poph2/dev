export type BaseHiveConfig = {
    port: number;
};

export type HiveServerConfig<Router, Middleware> = BaseHiveConfig & {
    routers: Router[];
    middlewares: Middleware[];
};

export type HiveConfig = BaseHiveConfig;
