import {Body, Controller, Post, Route, Tags} from "@tsoa/runtime";

@Tags("User")
@Route("user")
export class UserApi extends Controller {

    @Post("/")
    async createUser(@Body() createUserRequest: { name: string }) {
        return {name: "John Doe"};
    }

    @Post("/login")
    async login(@Body() loginRequest: { email: string, password: string }) {
        return {name: "John Doe"};
    }
}
