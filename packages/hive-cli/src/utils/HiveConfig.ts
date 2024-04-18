import Koa from 'koa';


export type HiveConfig = {
    port: number;
    middlewares: Koa.Middleware[];

    project?: {
        sourceDir: string;
        fileExtToWatch: string;
    }

    typeorm: {
        entities: string[];
    }
}