import { Injectable } from '@nestjs/common';
import {PrismaService} from "../prisma/prisma.service";
import {CreateUserDto} from "./dto/CreateUserDto";

@Injectable()
export class UsersService {
    constructor(private prisma: PrismaService) {}
    async create(dto: CreateUserDto) {
        return this.prisma.$transaction(async (tx) => {
            const user = await tx.user.create({
                data: {
                    email: dto.email,
                    password: dto.password,
                    role: dto.role,
                },
            });

            if (dto.role === 'TRAINER') {
                await tx.trainerProfile.create({
                    data: {
                        userId: user.id,
                        specialization: dto.specialization ?? 'general',
                        experience: dto.experience ?? 0,
                    },
                });
            }

            if (dto.role === 'CLIENT') {
                await tx.clientProfile.create({
                    data: {
                        userId: user.id,
                        height: dto.height ?? 0,
                        weight: dto.weight ?? 0,
                    },
                });
            }
            return user;
        });
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
