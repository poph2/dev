import "reflect-metadata"
import {Photo} from "./entities/Photo";
import {DataSource} from "typeorm";
import {User} from "./entities/User";


export const AppDataSource = new DataSource({
    type: "postgres",
    host: process.env.DATABASE_HOST,
    port: 5432,
    username: process.env.DATABASE_USER,
    password: process.env.DATABASE_PASSWORD,
    database: process.env.DATABASE_NAME,
    synchronize: true,
    logging: true,
    entities: [Photo, User],
    subscribers: [],
    migrations: ['src/migrations/*{.ts,.js}'],
})

