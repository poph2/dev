import path from "path";
import {glob} from "glob";

export const generateDataSource = () => {

    const filePaths = glob.sync(path.join("src", "**", "entities", "*.ts"), {
        cwd: process.cwd(),
    });

    console.log(filePaths);

}