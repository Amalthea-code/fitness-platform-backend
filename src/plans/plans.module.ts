import { Module } from '@nestjs/common';
import { PlansController } from './plans.controller';
import {PrismaService} from "../prisma/prisma.service";

@Module({
  providers: [PrismaService],
  controllers: [PlansController]
})
export class PlansModule {}
