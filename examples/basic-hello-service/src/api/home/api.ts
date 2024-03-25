import {Get} from "hive-runtime";

export default class HomeApi {

    @Get("/app-version")
    async getAppVersion() {
        return {version: "1.0.0"};
    }

    @Get("/params")
    async getParams() {
        return {name: "John Doe"};
    }
}
