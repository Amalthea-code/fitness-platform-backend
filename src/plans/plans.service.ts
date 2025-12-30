import { Injectable } from '@nestjs/common';
import { UserRole } from "../generated/prisma/enums";
import { PrismaService } from "../prisma/prisma.service";

@Injectable()
export class PlansService {
    constructor(private prisma: PrismaService) {}

    create(data: { email: string; password: string; role: UserRole; }) {

    }
}
