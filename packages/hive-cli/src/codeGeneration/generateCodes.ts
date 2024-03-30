import {generateMainTs} from "./generateMainTs";
import {Config} from "@tsoa/runtime";
import {ExtendedRoutesConfig, ExtendedSpecConfig, generateRoutes, generateSpec} from "@tsoa/cli";

export const runTsoa = async () => {

  const config: Config = {
    entryFile: "./src/Main.ts",
    noImplicitAdditionalProperties: "throw-on-extras",
    controllerPathGlobs: ["./src/**/api.ts"],
    spec: {
      "outputDirectory": "./src/tsoa/generatedSpec",
      "host": "localhost:3000",
      "spec": {
        "servers": [
          {
            "url": "http://localhost:8080",
            "description": "Local development"
          },
          {
            "url": "http://sd-service-staging.us-east-1.elasticbeanstalk.com/",
            "description": "Staging Environment"
          },
          {
            "url": "https://api.poph2.com/",
            "description": "Production Environment"
          }
        ]
      },
      "basePath": "/",
      "yaml": false,
      "specVersion": 3,
      "contact": {
        "name": "Pop H2",
        "email": "niyipopp@yahoo.co.uk"
      },
      "name": "sd-service",
      "specMerging": "immediate",
      "version": "1.0.0",
      "schemes": ["http", "https"],
      "xEnumVarnames": true,
      "securityDefinitions": {
        "optional_firebase_user": {
          "type": "http",
          "scheme": "bearer"
        },
        "firebase_user": {
          "type": "http",
          "scheme": "bearer"
        },
        "service": {
          "type": "http",
          "scheme": "bearer"
        }
      }
    },
    routes: {
      "basePath": "/",
      "middleware": "koa",
      "routesDir": "./src/tsoa/generatedRoutes",
      // "authenticationModule": "./src/middlewares/Authentication.ts"
    },
    ignore: ["node_modules/**/*", "test/**/*"],
  }
  const controllerPathGlobs = ["./src/**/api.ts"];

  const specOptions: ExtendedSpecConfig = {
    noImplicitAdditionalProperties: "throw-on-extras",
    basePath: "/api",
    entryFile: "./api/server.ts",
    specVersion: 3,
    outputDirectory: "./.hive",
    controllerPathGlobs
  };

  const routeOptions: ExtendedRoutesConfig = {
    bodyCoercion: false,
    noImplicitAdditionalProperties: "throw-on-extras",
    basePath: "/api",
    entryFile: "./api/server.ts",
    routesDir: "./.hive",
    controllerPathGlobs,
    middleware: "koa"
  };

  await generateSpec(specOptions);

  await generateRoutes(routeOptions);


}

export const generateCodes = async () => {


  await runTsoa()

  generateMainTs();


};
