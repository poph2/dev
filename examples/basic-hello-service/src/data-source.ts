import "reflect-metadata"
import {Photo} from "./entities/Photo";
import {DataSource} from "typeorm";
import {User} from "./entities/User";

export const AppDataSource = new DataSource({
    type: "postgres",
    host: "localhost",
    port: 5432,
    username: "hive",
    password: "hive",
    database: "hive",
    synchronize: true,
    logging: true,
    entities: [Photo, User],
    subscribers: [],
    migrations: ['src/migrations/*{.ts,.js}'],
})

