import ts from "typescript";
import fs from "fs";


export const readSourceFile = (filePath: string): ts.SourceFile => {
  const sourceCode = fs.readFileSync(filePath, "utf-8");

  return ts.createSourceFile(
    filePath,
    sourceCode,
    ts.ScriptTarget.Latest,
    /* setParentNodes */ true
  );
}