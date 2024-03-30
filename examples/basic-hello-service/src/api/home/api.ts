import {Controller, Get, Route, Tags} from "@tsoa/runtime";

@Tags("Home")
@Route("home")
export class HomeApi extends Controller {

    @Get("/app-version")
    async getAppVersion() {
        return {version: "1.0.0"};
    }

    @Get("/params")
    async getParams() {
        return {name: "John Doe"};
    }
}
