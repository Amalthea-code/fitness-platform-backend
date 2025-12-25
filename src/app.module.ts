import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { AuthModule } from './auth/auth.module';
import { UsersModule } from './users/users.module';
import { TrainersModule } from './trainers/trainers.module';
import { ClientsModule } from './clients/clients.module';
import { PlansModule } from './plans/plans.module';
import { WorkoutsModule } from './workouts/workouts.module';
import { ExercisesModule } from './exercises/exercises.module';
import { PrismaService } from './prisma/prisma.service';

@Module({
  imports: [
      AuthModule,
      UsersModule,
      TrainersModule,
      ClientsModule,
      PlansModule,
      WorkoutsModule,
      ExercisesModule
  ],
  controllers: [AppController],
  providers: [AppService, PrismaService],
})
export class AppModule {}
