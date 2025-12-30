import {Body, Controller, Get, Param, Post} from '@nestjs/common';
import { UsersService } from "./users.service";

@Controller('users')
export class UsersController {
    constructor(private readonly usersService: UsersService) {}
    @Post()

    @Get()
    findAll() {
        return this.usersService.findAll()
    }

    @Get(':id')
    findById(@Param('id') id: string) {
        return this.usersService.findById(id);
    }

    @Get(':email')
    findByEmail(@Param('email') email: string) {
        return this.usersService.findByEmail(email);
    }

}
