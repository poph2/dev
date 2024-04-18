import {glob} from "glob";
import path from "path";
import {HiveConfig} from "../utils/HiveConfig";

export const generateDataSources = () => {

    const hiveConfig = require(path.join(process.cwd(), 'hive.config.json')) as HiveConfig;

    console.log(hiveConfig);

    const filePaths = glob.sync(path.join("src", "**", "entities", "*.ts"), {
        cwd: process.cwd(),
    });

    console.log(filePaths);

}