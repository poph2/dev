import {glob} from "glob";
import path from "path";
import ts from "typescript";
import _ from "lodash";
import {getDecoratorValues, getDecorators, readSourceFile, extractClasses, extractMethods, generateCodeFile} from "compiler-kit";


export enum ApiParameterType {
  Body = "Body",
}

export const ApiParameterCode: {
    [key in ApiParameterType]: string;
} = {
  Body: "(ctx.request as any).body",
}

export type ApiParameter = {
  type: ApiParameterType;
  code: string;
}

export type ApiMethodSummary = {
  httpMethod: string;
  path: string;
  methodName: string;
  apiParameters: ApiParameter[];
};

export type ApiClassSummary = {
  className: string;
  instanceName: string;
  apiFile: string;
  importPath: string;
  httpPath: string;
  methodSummaries: ApiMethodSummary[];
};

const getApiParameters = (method: ts.MethodDeclaration): ApiParameter[] => {
    const parameters = method.parameters.filter(node => ts.isParameter(node)) as ts.ParameterDeclaration[];
  const apiParameters: ApiParameter[] = [];

  for(const parameter of parameters) {
    const decorators = getDecorators(parameter);
    if (!decorators || decorators.length === 0) continue;

    if (decorators.length > 1) {
      throw new Error("Multiple decorators found on parameter: " + parameter.getText());
    }

    const {name: decoratorName} = getDecoratorValues(decorators[0]);

    if (!Object.keys(ApiParameterType).includes(decoratorName)) {
      throw new Error(`Invalid parameter decorator: ${decoratorName}`);
    }


    apiParameters.push({
      type: decoratorName as ApiParameterType,
      code: ApiParameterCode[decoratorName as ApiParameterType]
    });

  }
  return apiParameters;
}

const getApiMethodSummary = (method: ts.MethodDeclaration): ApiMethodSummary | undefined => {
  const decorators = ts.getDecorators(method);
  if (!decorators || decorators.length === 0) return undefined;

  for (const decorator of decorators) {
    const callExpression = decorator.expression as ts.CallExpression;
    const identifier = callExpression.expression as ts.Identifier;
    const decoratorName = identifier.getText();

    if (["Get", "Post"].includes(decoratorName)) {
      const args = callExpression.arguments.filter(node => ts.isStringLiteral(node)) as ts.StringLiteral[];
      return {
        httpMethod: decoratorName.toLowerCase(),
        path: args[0].text,
        methodName: ((method.name) as ts.Identifier).escapedText as string,
        apiParameters: getApiParameters(method),
      };
    }
  }

  return undefined;
};

const getApiClass = (filePath: string): ApiClassSummary => {

  const sourceFile = readSourceFile(filePath);
  const klasses = extractClasses(sourceFile);

  if (klasses.length === 0) {
    throw new Error(`No class found in ${filePath}`);
  }

  if (klasses.length > 1) {
    throw new Error(`Multiple classes found in ${filePath}`);
  }

  const klass = klasses[0];

  const className = klass.name?.getText() as string;

  return {
    className,
    instanceName:
      className.charAt(0).toLowerCase() + className.slice(1),
    importPath: `../${filePath.slice(0, -3)}`,
    httpPath: filePath.split(path.sep).slice(2, -1).join("/"),
    apiFile: filePath,
    methodSummaries: _(extractMethods(klass))
        .map(method => getApiMethodSummary(method))
        .compact().value(),
  };
};

export const generateRoutersTs = (): void => {
  // Fetch all api.ts files that match cwd/api/**/api.ts
  const filePaths = glob.sync(path.join("src", "api", "**", "api.ts"), {
    cwd: process.cwd(),
  });

  // Extract router summaries from each api.ts file
  const routerSummaries = filePaths.map(getApiClass).flat();

  // TODO: Validate the following:
  //  1. No duplicate class names
  //  2. No duplicate paths

  // prettyPrint(routerSummaries);

  generateCodeFile(
    { routerSummaries },
    {
      pathToTemplate: [__dirname, "..", "templates", "routers.ts.hbs"],
      pathToOutput: [process.cwd(), ".hive", "routers.ts"],
    },
  );
};
