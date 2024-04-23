#!/usr/bin/env node

import {ChildProcess, spawn} from "child_process";
import * as chokidar from "chokidar";
import {debounce} from "lodash";
import {generateCodes} from "./codeGeneration/generateCodes";
import {hiveConfig} from "./utils/HiveConfig";


const fileToRun = "./.hive/Main.ts";

let childProcess: ChildProcess | null = null;

export const runTsFile = async () => {
  if (childProcess) {
    childProcess.kill();
  }

  await generateCodes();

  childProcess = spawn("ts-node", [fileToRun], { stdio: "inherit" });

  childProcess.on("exit", (code) => {
    console.log(`Child process exited with code ${code}`);
  });
};

export const runCli = async () => {
  console.log("running the cli.....");

  await runTsFile();

  const debouncedRunTsFile = debounce(runTsFile, 500);

  const watcher = chokidar.watch(hiveConfig.project.sourceDir, {
    ignored: /(^|[\/\\])\../, // ignore dotfiles
    persistent: true,
  });

  watcher.on("change", (path) => {
    if (hiveConfig.project.fileExtToWatch.includes(path.split(".").pop() || "")) {
      console.log("");
      console.log(`File ${path} has been changed. Restarting...`);
      debouncedRunTsFile();
    }
  });
};

if (require.main === module) {
  runCli().catch((e) => {
    console.error(e);
    process.exit(1);
  });
}
