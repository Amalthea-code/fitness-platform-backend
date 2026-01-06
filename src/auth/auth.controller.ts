import {BadRequestException, Body, Controller, Post, Res} from '@nestjs/common';
import {RegisterDto} from "./dto/register.dto";
import {UsersService} from "../users/users.service";
import {PrismaService} from "../prisma/prisma.service";
import {LoginDto} from "./dto/login.dto";

// POST   /auth/register
// POST   /auth/login
// POST   /auth/refresh
// POST   /auth/logout
// GET    /auth/me

@Controller('auth')
export class AuthController {
    constructor() {}
}
