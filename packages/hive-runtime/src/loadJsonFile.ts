import fs from "fs";

export const loadJsonFile = <T>(filePath: string) => {
    return JSON.parse(fs.readFileSync(filePath, 'utf8')) as T;
}