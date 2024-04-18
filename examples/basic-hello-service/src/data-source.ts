import "reflect-metadata"
import {Photo} from "./models/Photo";
import {DataSource} from "typeorm";
import {User} from "./models/User";

const migrationPaths: string[] = process.env.TYPEORM_USE_CLI === 'true' ? ['src/migrations/*{.ts,.js}'] : [];


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
    migrations: migrationPaths,
})

// export const AppDataSource2 = new DataSource({
//     type: "postgres",
//     host: "localhost",
//     port: 5432,
//     username: "hive",
//     password: "hive",
//     database: "hive",
//     synchronize: true,
//     logging: false,
//     entities: [User],
//     migrations: [],
//     subscribers: [],
// })