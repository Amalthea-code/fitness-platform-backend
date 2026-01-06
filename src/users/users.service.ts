import { Injectable } from '@nestjs/common';
import {PrismaService} from "../prisma/prisma.service";
import {ClientDto} from "./dto/client.dto";

@Injectable()
export class UsersService {
    constructor(private prisma: PrismaService) {}

    async createUser(dto: ClientDto) {
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

    async createTrainer(dto: TrainerDto) {

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
}
