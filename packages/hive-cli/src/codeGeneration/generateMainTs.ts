import {hiveConfig} from "../utils/HiveConfig";
import {generateCodeFile} from "../utils/utils";

export const generateMainTs = () => {
  generateCodeFile(
    {port: hiveConfig.port},
    {
      pathToTemplate: [__dirname, "..", "templates", "Main.ts.hbs"],
      pathToOutput: [process.cwd(), ".hive", "Main.ts"],
    },
  );
};
