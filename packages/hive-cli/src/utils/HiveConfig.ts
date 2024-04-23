import {loadJsonFile} from "./loadJsonFile";

export type HiveConfig = {
  project: {
    sourceDir: string[];
    fileExtToWatch: string[];
  };
  port: number;
  typeorm?: {
    entities: string[];
  };
};

export const hiveConfig = loadJsonFile<HiveConfig>("hive.config.json");
