import { Injectable } from '@nestjs/common';
import {PrismaService} from "../prisma/prisma.service";
import { UserRole } from '@prisma/client';

@Injectable()
export class UsersService {
    constructor(private prisma: PrismaService) {}

    create(data: { email: string; password: string; role: UserRole; }) {
        return this.prisma.user.create({ data });
    }

    findAll() {
        return this.prisma.user.findMany();
    }

    findById(id: string) {
        return this.prisma.user.findUnique({ where: { id } });
    }

    findByEmail(email: string) {
        return this.prisma.user.findUnique({ where: { email } })
    }

    deactivate(id: string) {
        return this.prisma.user.update({
            where: { id },
            data: { isActive: false },
        })
    }
}
