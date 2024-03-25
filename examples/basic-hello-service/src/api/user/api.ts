import {Body, Post} from "hive-runtime";

export default class UserApi {

    // @Post("/app-version")
    // async getAppVersion() {
    //     return {version: "1.0.0"};
    // }
    //
    // @Get("/params")
    // async getParams() {
    //     return {name: "John Doe"};
    // }

    @Post("/")
    async createUser(@Body() createUserRequest: { name: string }) {
        return {name: "John Doe"};
    }
}
