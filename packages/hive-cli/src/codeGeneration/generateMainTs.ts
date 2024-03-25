import {generateCodeFile} from "../utils/utils";

export const generateMainTs = () => {
  generateCodeFile(
    {},
    {
      pathToTemplate: [__dirname, "..", "templates", "Main.ts.hbs"],
      pathToOutput: [process.cwd(), ".hive", "Main.ts"],
    },
  );
};
