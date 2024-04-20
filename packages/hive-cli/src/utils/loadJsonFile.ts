import fs from "fs";
import {HiveConfig} from "./HiveConfig";

export const loadJsonFile = <T>(filePath: string) => {
  return JSON.parse(fs.readFileSync(filePath, "utf8")) as T;
};

export const loadHiveConfig = () => loadJsonFile<HiveConfig>("hive.config.json");
