import {generateMainTs} from "./generateMainTs";
import {ExtendedRoutesConfig, ExtendedSpecConfig, generateRoutes, generateSpec} from "@tsoa/cli";
import {HiveConfig} from "hive-runtime";
import {loadJsonFile} from "hive-config";

export const runTsoa = async () => {

  const commonOpts = {
    controllerPathGlobs: ["./src/**/api.ts"],
    basePath: "/",
    entryFile: "./.hive/Main.ts",
    noImplicitAdditionalProperties: "throw-on-extras" as any,
    bodyCoercion: false,
  };

  const specOptions: ExtendedSpecConfig = {
    ...commonOpts,
    specVersion: 3,
    outputDirectory: "./.hive",
    contact: {
      "name": "Pop H2",
        "email": "niyipopp@yahoo.co.uk"
    },
    name: "sd-service",
    version: "1.0.0",
  };

  const routeOptions: ExtendedRoutesConfig = {
    ...commonOpts,
    routesDir: "./.hive",
    middleware: "koa"
  };

  await generateSpec(specOptions);

  await generateRoutes(routeOptions);


}

export const generateCodes = async () => {

  const hiveConfig = loadJsonFile<HiveConfig>('hive.config.json');

  // const {default: hiveConfig} = await import(path.join(process.cwd(), 'hive.config.ts')) as {default: HiveConfig};

  // const hiveConfig = require(path.join(process.cwd(), 'hive.config.ts')) as HiveConfig;
  console.log(hiveConfig, "loaded at runtime......")


  await runTsoa()

  generateMainTs();




};
