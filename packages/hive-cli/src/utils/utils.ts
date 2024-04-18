import * as fs from "fs";
import path from "path";
import * as Handlebars from "handlebars";

export function createDirectoryIfNotExists(directoryPath: string): void {
  if (!fs.existsSync(directoryPath)) {
    fs.mkdirSync(directoryPath, { recursive: true });
  }
}

export const writeToFile = (filePath: string, content: string) => {
  createDirectoryIfNotExists(path.dirname(filePath));
  fs.writeFileSync(filePath, content);
};

export const mapToKeyValue = (obj: { [key: string]: string }) => {
  return Object.entries(obj).map(([key, value]) => ({ key, value }));
};

export const prettyPrint = (obj: any, title?: string) => {
  console.log(JSON.stringify(obj, null, 2), title || "");
};

export const generateCodeFile = <T>(
  context: T,
  paths: { pathToTemplate: string[]; pathToOutput: string[] },
) => {
  const source = fs.readFileSync(path.join(...paths.pathToTemplate)).toString();
  const template = Handlebars.compile<T>(source);
  const result = template(context);
  writeToFile(path.join(...paths.pathToOutput), result);
};



