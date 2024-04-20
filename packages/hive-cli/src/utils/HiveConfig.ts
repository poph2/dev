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
