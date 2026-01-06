import {BadRequestException, Injectable, UnauthorizedException} from '@nestjs/common';
import { PrismaService } from "../prisma/prisma.service";
import * as bcrypt from 'bcrypt'
import { RegisterDto } from "./dto/register.dto";

@Injectable()
export class AuthService {
    constructor(
        private prisma: PrismaService,
        // private jwt: JwtService,
    ) {}
}
