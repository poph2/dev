import {generateCodes} from "./generateCodes";

process.chdir("../../examples/basic-hello-service")

describe('generateRouter', () => {

    it('should generate router', () => {

        generateCodes()

        console.log("klasses.length");
        console.log(process.cwd())
        process.chdir("../../examples/basic-hello-service")
        console.log(process.cwd())

    });

});