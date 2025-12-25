import {Body, Controller, Get, Param, Post} from '@nestjs/common';
import {UsersService} from "./users.service";
import { UserRole } from "@prisma/client";

@Controller('users')
export class UsersController {
    constructor(private readonly usersService: UsersService) {}
    @Post()
    create(@Body() body: { email: string; password: string; role: string }) {
        if (!body || !body.email || !body.password || !body.role) {
            throw new Error('Missing fields in request body');
        }

        const role: UserRole = UserRole[body.role as keyof typeof UserRole];

        return this.usersService.create({
            email: body.email,
            password: body.password,
            role,
        });
    }

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
